package dbutil

import (
	"github.com/dgraph-io/badger"
)

// DbRead is a key value search
func DbRead(db *badger.DB, k string) string {
	// Start a readable transaction
	txn := db.NewTransaction(false)
	//Implicityly called when Commit() is called or used to discard read
	// only transaction, either way safe to defer this function to the end
	defer txn.Discard()
	//k, _ = stdInRead("readMode")

	var v string

	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(k))
		if err != nil {
			panic(err)
		}

		err = item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.
			//fmt.Printf("The key is %s, the value is: %s\n", k, val)

			v = string(val)
			return nil
		})
		if err != nil {
			panic(err)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
	return v
}
