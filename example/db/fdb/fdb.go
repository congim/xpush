package main

import (
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"github.com/congim/xpush/pkg/message"
)

var (
	fdbStoreBegin = []byte("0")
	fdbStoreEnd   = []byte("ff")
	//MSG_NEWEST_OFFSET = []byte("0")
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	v, _ := fdb.GetAPIVersion()
	log.Println("fdb api version is", v)

	// Different API versions may expose different runtime behaviors.
	if err := fdb.APIVersion(610); err != nil {
		v, _ := fdb.GetAPIVersion()
		log.Println("fdb api version isn't 610, version is", v)
		return
	}

	// Open the default database from the system cluster
	db := fdb.MustOpenDefault()

	dir, err := directory.CreateOrOpen(db, []string{"xpush"}, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var msg subspace.Subspace
	msg = dir.Sub("msg-body")
	var msgs []*message.Message
	_, _ = db.Transact(func(tr fdb.Transaction) (ret interface{}, err error) {
		pr, _ := fdb.PrefixRange([]byte("test"))
		pr.Begin = msg.Pack(tuple.Tuple{[]byte("test"), fdbStoreBegin})
		//if bytes.Compare([]byte("0"), message.MSG_NEWEST_OFFSET) == 0 {
		//	pr.End = msg.Pack(tuple.Tuple{[]byte("test"), fdbStoreEnd})
		//} else {
		pr.End = msg.Pack(tuple.Tuple{[]byte("test"), "1573712723843489000"})
		//}

		//@performance
		//get one by one in advance
		ir := tr.GetRange(pr, fdb.RangeOptions{Limit: 50, Reverse: true}).Iterator()
		for ir.Advance() {
			b := ir.MustGet().Value
			m := &message.Message{}
			if err := m.Decode(b); err != nil {
				log.Println("msg decode", err)
				continue
			}
			msgs = append(msgs, m)
			log.Println("拉取新消息为:", m.Topic, m.ID, string(m.Payload))
		}
		return
	})
	log.Println("运行结束")
}
