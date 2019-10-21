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

type FDB struct {
	conf   *config.Storage
	logger *zap.Logger
	dbs    []*database
	rand   *rand.Rand
}

type database struct {
	db  fdb.Database
	msg subspace.Subspace
	//counter subspace.Subspace
	//topic   subspace.Subspace
	//session subspace.Subspace
}

func New(conf *config.Storage, logger *zap.Logger) *FDB {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &FDB{
		conf:   conf,
		logger: logger,
		rand:   r,
	}
}

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
	//counter := dir.Sub("msg-count")
	//topic := dir.Sub("topic")
	//session := dir.Sub("session")

	f.dbs[i] = &database{
		msg: msg,
		//counter: counter,
		db: db,
		//topic:   topic,
		//session: session,
	}

	return nil
}

func (f *FDB) Store(msgs []*message.Message) error {
	index := f.rand.Intn(f.conf.Threads)
	if index >= f.conf.Threads && index < 0 {
		return fmt.Errorf("index err, index is %d, threads is %d", index, f.conf.Threads)
	}

	_, err := f.dbs[index].db.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		for _, msg := range msgs {
			key := f.dbs[index].msg.Pack(tuple.Tuple{msg.Topic, msg.ID})
			b, _ := msg.Encode()
			log.Println("存储", string(msg.Payload))
			tr.Set(key, b)
		}
		return
	})

	//for index, msg := range msgs {
	//	f.Get(msg.Topic, index, msg.ID)
	//}

	if err != nil {
		f.logger.Info("store messsage error", zap.Error(err))
	}

	return nil
}

func (f *FDB) Get(topic string, count int, offset string) ([]*message.Message, error) {
	var msgs []*message.Message
	index := f.rand.Intn(f.conf.Threads)
	if index >= f.conf.Threads && index < 0 {
		return nil, fmt.Errorf("index err, index is %d, threads is %d", index, f.conf.Threads)
	}

	_, err := f.dbs[index].db.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		pr, _ := fdb.PrefixRange([]byte(topic))
		pr.Begin = f.dbs[index].msg.Pack(tuple.Tuple{[]byte(topic), []byte("0")})

		pr.End = f.dbs[index].msg.Pack(tuple.Tuple{[]byte(topic), offset})

		//@performance
		//get one by one in advance
		ir := tr.GetRange(pr, fdb.RangeOptions{Limit: count, Reverse: true}).Iterator()
		for ir.Advance() {
			b := ir.MustGet().Value
			m := &message.Message{}
			m.Decode(b)
			msgs = append(msgs, m)
			log.Println(string(m.Payload))
		}
		return
	})
	return msgs, err
}
