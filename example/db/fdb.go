package main

import (
	"fmt"
	"log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
)

func main() {
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

	// Database reads and writes happen inside transactions
	ret, e := db.Transact(func(tr fdb.Transaction) (interface{}, error) {
		tr.Set(fdb.Key("hello"), []byte("world"))
		return tr.Get(fdb.Key("foo")).MustGet(), nil
		// db.Transact automatically commits (and if necessary,
		// retries) the transaction
	})
	if e != nil {
		log.Fatalf("Unable to perform FDB transaction (%v)", e)
	}

	fmt.Printf("hello is now world, foo was: %s\n", string(ret.([]byte)), e)
}
