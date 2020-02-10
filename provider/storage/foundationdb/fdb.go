package foundationdb

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

// FDB ...
type FDB struct {
	conf   *config.Fdb
	logger *zap.Logger
	dbs    []*database
	rand   *rand.Rand
	stopC  chan struct{}
	//msgQueue chan []*message.Message
}

type database struct {
	db       fdb.Database
	msg      subspace.Subspace
	msgQueue chan []*message.Message
}

// New ...
func New(conf *config.Fdb, logger *zap.Logger) *FDB {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &FDB{
		conf:   conf,
		logger: logger,
		rand:   r,
		stopC:  make(chan struct{}, 1),
		//msgQueue: make(chan []*message.Message, 500),
	}
}

// Init ...
func (f *FDB) Init() error {
	f.dbs = make([]*database, f.conf.Threads)
	for i := 0; i < f.conf.Threads; i++ {
		if err := f.init(i); err != nil {
			f.logger.Error("init fdb failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func (f *FDB) init(i int) error {
	if err := fdb.APIVersion(610); err != nil {
		f.logger.Warn("api version failed", zap.Error(err))
		return err
	}
	db, err := fdb.OpenDefault()
	if err != nil {
		f.logger.Error("openDefault failed", zap.Error(err))
		return err
	}

	dir, err := directory.CreateOrOpen(db, []string{f.conf.DBSpace}, nil)
	if err != nil {
		f.logger.Error("CreateOrOpen failed", zap.Error(err))
		return err
	}

	msg := dir.Sub("msg-body")
	f.dbs[i] = &database{
		msg:      msg,
		db:       db,
		msgQueue: make(chan []*message.Message, 100),
	}
	go f.store(i)
	return nil
}

func (f *FDB) store(index int) {
	defer func() {
		close(f.dbs[index].msgQueue)
	}()
	for {
		select {
		case <-f.stopC:
			return
		case msgs, ok := <-f.dbs[index].msgQueue:
			if ok {
				_, err := f.dbs[index].db.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
					for _, msg := range msgs {
						key := f.dbs[index].msg.Pack(tuple.Tuple{[]byte(msg.Topic), []byte(msg.ID)})
						b, err := msg.Encode()
						if err != nil {
							f.logger.Warn("msg encode error", zap.Error(err))
							continue
						}
						tr.Set(key, b)
						log.Println("insert key", key)
					}
					return
				})
				if err != nil {
					f.logger.Warn("store message error", zap.Error(err))
				}
			}
			break
		}
	}
}

// Store ...
func (f *FDB) Store(msgs []*message.Message) error {
	index := f.rand.Intn(f.conf.Threads)
	if index >= f.conf.Threads && index < 0 {
		return fmt.Errorf("index err, index is %d, threads is %d", index, f.conf.Threads)
	}
	f.dbs[index].msgQueue <- msgs
	return nil
}

var (
	fdbStoreBegin = []byte("0")
	fdbStoreEnd   = []byte("ff")
)

// Get ...
func (f *FDB) Get(topic string, offset []byte, count int) ([]*message.Message, error) {
	var msgs []*message.Message
	index := f.rand.Intn(f.conf.Threads)
	if index >= f.conf.Threads && index < 0 {
		return nil, fmt.Errorf("index err, index is %d, threads is %d", index, f.conf.Threads)
	}

	_, err := f.dbs[index].db.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		//pr, _ := fdb.PrefixRange([]byte(topic))

		pr, _ := fdb.PrefixRange(f.dbs[index].msg.Pack(tuple.Tuple{[]byte(topic), offset}))
		//pr.Begin = f.dbs[index].msg.Pack(tuple.Tuple{[]byte(topic), fdbStoreBegin})
		//if bytes.Compare(offset, message.MSG_NEWEST_OFFSET) == 0 {
		//	pr.End = f.dbs[index].msg.Pack(tuple.Tuple{[]byte(topic), fdbStoreEnd})
		//} else {
		//	pr.End = f.dbs[index].msg.Pack(tuple.Tuple{[]byte(topic), offset})
		//}

		log.Println("start key", pr.Begin)
		log.Println("end key", pr.End)
		log.Println("topic", topic)
		log.Println("offset", string(offset))
		log.Println("count", count)
		//@performance
		//get one by one in advance
		ir := tr.GetRange(pr, fdb.RangeOptions{Limit: count, Reverse: true}).Iterator()
		for ir.Advance() {
			b := ir.MustGet().Value
			m := &message.Message{}
			if err := m.Decode(b); err != nil {
				f.logger.Warn("msg decode error", zap.Error(err))
				continue
			}
			msgs = append(msgs, m)
			log.Println("拉取新消息为:", m.Topic, m.ID, string(m.Payload))
		}
		return
	})
	return msgs, err
}

// Close ...
func (f *FDB) Close() error {
	close(f.stopC)
	return nil
}
