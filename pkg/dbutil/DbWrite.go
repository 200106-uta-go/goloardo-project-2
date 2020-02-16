package dbutil

import (
	"log"

	"github.com/dgraph-io/badger"
)

// ByteConversion takes a vey-value pair in the form of two strings and returns a pair of []bytes
func ByteConversion(key, value string) ([]byte, []byte) {
	return []byte(key), []byte(value)
}

// DbWrite is an update to the json database
func DbWrite(db *badger.DB, k string, v string) (string, string) {

	// Start a writable transaction
	txn := db.NewTransaction(true)
	//Implicityly called when Commit() is called or used to discard read
	// only transaction, either way safe to defer this function to the end
	defer txn.Discard()

	err := db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(ByteConversion(k, v))
		err := txn.SetEntry(e)
		return err
	})
	if err != nil {
		log.Panic(err)
	}
	return k, v
}
