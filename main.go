package main

import (
	"Badger-db/BadgerDbOperations"
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"log"
	"os"
)

func main() {
	// Open the Badger database located in the tmp/badger directory.
	// It will be created if it doesn't exist.

	//opts := badger.DefaultOptions
	//opts.Dir = "./data"
	//opts.Value = "./data"
	//kv, err := badger.NewKV(&opts)
	db, err := badger.Open(badger.DefaultOptions("tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Welcome to badger db implementation!")
	for {
		var choice int
		fmt.Println("Enter your choice")
		fmt.Println("1. Insert data")
		fmt.Println("2. Display")
		fmt.Println("3. Delete")
		fmt.Println("4. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var key, value string
			fmt.Println("Enter key")
			fmt.Scanln(&key)
			fmt.Println("Enter value")
			fmt.Scanln(&value)
			BadgerDbOperations.InsertData(db, key, value)

		case 2:
			BadgerDbOperations.DisplayData(db)
		case 3:
			var key string
			fmt.Println("Enter key")
			fmt.Scanln(&key)
			BadgerDbOperations.DeleteData(db, key)
		default:
			os.Exit(0)
		}

	}
}

/*
Files in badger db:
	1. Sorted string tables (SST's) Files- stores keys(Optionally values)
	2. Value log(Vlog) Files - stores values and serves as a write-ahead-log
	3. Key Registry File - stores info about encryption
	4. Manifest file - stores info about all SST's

Design:
	- Contains an in-Memory skipList used for reading and writing.
	- Contains on disk SST files which are used for reading only.
	- LSM tree can be visualised as Pyramid with each level containing more data than a previous level.
*/
