// key value storage

package utils

import (
	badger "github.com/dgraph-io/badger/v3"
)

var DefaultKvs Kvs

type Kvs interface {
	Get(key string) (string, error)
	Set(key, value string) error
	Del(key string) error
}

type SimpleKvs struct {
	db *badger.DB
}

func NewKVS() (Kvs, error) {
	db, err := badger.Open(badger.DefaultOptions(".cas/kvs"))
	return &SimpleKvs{db: db}, err
}

func (s *SimpleKvs) Get(key string) (string, error) {
	value := ""
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
		return nil
	})
	return value, err
}

func (s *SimpleKvs) Set(key, value string) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *SimpleKvs) Del(key string) error {
	err := s.db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func init() {
	DefaultKvs, _ = NewKVS()
}
