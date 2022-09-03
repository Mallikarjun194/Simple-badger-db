package BadgerDbOperations

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
)

func InsertData(db *badger.DB, key string, value string) {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
	if err != nil {
		fmt.Println("Error while inserting the data", err)
	} else {
		fmt.Println("Record Inserted!")
	}
}

func DisplayData(db *badger.DB) {
	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error while displaying data from db", err)
	} else {
		fmt.Println("Successfully displayed all the data!")
	}

}

func DeleteData(db *badger.DB, key string) {
	err := db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			if string(k) == key {
				err := txn.Delete(k)
				if err != nil {
					fmt.Println("Error while deleting the key", err)
				} else {
					fmt.Println("Key deleted", k)
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

}
