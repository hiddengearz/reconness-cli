package db

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/MoreThanRecon/laelaps/internal"
	"github.com/MoreThanRecon/laelaps/internal/encrypt"
	"github.com/dgraph-io/badger/v3"

	l "github.com/MoreThanRecon/laelaps/internal/logger"
	"github.com/peterbourgon/diskv"
	"github.com/spf13/viper"
)

var (
	db2                *diskv.Diskv
	db                 *badger.DB
	badgerDiscardRatio = 0.5
	badgerGCInterval   = 10 * time.Minute
)

func InitDB() {
	var err error
	laelapsPath := filepath.Dir(viper.ConfigFileUsed())
	dbpath := laelapsPath + "/db/badger/"

	opts := badger.DefaultOptions(dbpath)
	opts.SyncWrites = true
	if internal.Debug == false {
		opts.Logger = nil
	}

	db, err = badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	go runGC()

}

func EncWrite(key string, data []byte, password string) (err error) {
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), data)
		return err
	})

	return err
}

func EncRead(key string, password string) (data []byte, err error) {
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

func Cleanup() {
again:
	err := db.RunValueLogGC(0.7)
	if err == nil {
		goto again
	}
}

func Close() {
	db.Close()
}

func InitDB2() {

	laelapsPath := filepath.Dir(viper.ConfigFileUsed())
	dbpath := laelapsPath + "/db/"

	_, err := os.Stat(dbpath)
	if os.IsNotExist(err) {
		l.Log.Debug("Database not found in", dbpath)
		l.Log.Info("Creating db at", dbpath)
		os.Mkdir(dbpath, 0700)
	}

	flatTransform := func(s string) []string { return []string{} }

	// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
	db2 = diskv.New(diskv.Options{
		BasePath:     dbpath,
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})

}

func EncWrite2(key string, data []byte, password string) (err error) {

	dataString, err := encrypt.EncryptData(data, password)
	if err != nil {
		l.Log.Debug(err)
		return
	}

	err = db2.Write(key, []byte(dataString))
	return err
}

func EncRead2(key string, password string) (data []byte, err error) {

	encString, err := db2.Read(key)
	if err != nil {
		l.Log.Debug(err)
		return
	}
	encData := string(encString)

	data, err = encrypt.DecryptData(encData, password)
	if err != nil {
		l.Log.Debug(err)
		return nil, err
	}
	return data, err
}
