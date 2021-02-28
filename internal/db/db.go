package db

import (
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/hiddengearz/reconness-cli/internal"
	"github.com/hiddengearz/reconness-cli/internal/crypto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	db                 *badger.DB
	badgerDiscardRatio = 0.5
	badgerGCInterval   = 10 * time.Minute
)

//InitDB will initialize the Badger Database.
func InitDB() {
	var err error
	rcliPath := filepath.Dir(viper.ConfigFileUsed())
	dbpath := rcliPath + "/db/"

	opts := badger.DefaultOptions(dbpath)
	opts.SyncWrites = true
	if internal.Debug == false {
		opts.Logger = nil
	}

	if internal.EncKey == "" {
		log.Fatal("Encryption key can't be empty")
	}

	if internal.Encryption == true { // All data written to the DB will be encrypted with a 32bit hash of the user's key
		key := crypto.HashTo32Bytes(internal.EncKey)
		internal.EncKey = "" //will this ensure that the plaintext key is no longer in memory?
		opts = opts.WithEncryptionKey(key)
	}

	db, err = badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	go runGC()

}

//Write will write the specified data at the key value address
func Write(key string, data []byte) (err error) {
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), data)
		return err
	})

	return err
}

//Read will read the specified key value address and return the stored data
func Read(key string) (data []byte, err error) {
	var item *badger.Item

	err = db.View(func(txn *badger.Txn) error {
		item, err = txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return item.ValueCopy(nil)

}

//ReadAllWithPrefix will read all key value adresses with the specified prefix and return the stored data in a 2D slice
func ReadAllWithPrefix(prefix string) (data [][]byte, err error) {

	err = db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(prefix)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				data = append(data, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

//runGC runs the garbage collector for the database
func runGC() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
	again:
		err := db.RunValueLogGC(0.7)
		if err == nil {
			goto again
		}
	}
}

//Cleanup is similar to runGC, cleansup the DB
func Cleanup() {
again:
	err := db.RunValueLogGC(0.7)
	if err == nil {
		goto again
	}
}

//Close closes the DB
func Close() {
	db.Close()
}
