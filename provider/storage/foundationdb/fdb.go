package foundationdb

import (
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/congim/xpush/config"
	"github.com/congim/xpush/pkg/message"
	"go.uber.org/zap"
)

type FDB struct {
	conf   *config.Storage
	logger *zap.Logger
	dbs    []*database
}

type database struct {
	db      fdb.Database
	msg     subspace.Subspace
	counter subspace.Subspace
	topic   subspace.Subspace
	session subspace.Subspace
}

func New(conf *config.Storage, logger *zap.Logger) *FDB {
	return &FDB{
		conf:   conf,
		logger: logger,
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
	counter := dir.Sub("msg-count")
	//topic := dir.Sub("topic")
	//session := dir.Sub("session")

	f.dbs[i] = &database{
		msg:     msg,
		counter: counter,
		db:      db,
		//topic:   topic,
		//session: session,
	}

	return nil
}

func (f *FDB) Store(msg *message.Message) error {
	return nil
}
