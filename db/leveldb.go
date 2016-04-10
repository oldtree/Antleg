package db

import (
	"errors"

	"github.com/syndtr/goleveldb/leveldb"
)
import "github.com/syndtr/goleveldb/leveldb/util"

type LeveldbStore struct {
	DB *leveldb.DB
}

func (ls *LeveldbStore) Get(key interface{}) (interface{}, bool, error) {
	data, err := ls.DB.Get([]byte(key.([]byte)), nil)
	if err == nil {
		return nil, false, err
	}
	return data, false, nil
}
func (ls *LeveldbStore) Set(key interface{}, value interface{}) (bool, error) {
	err := ls.DB.Put([]byte(key.([]byte)), []byte(value.([]byte)), nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (ls *LeveldbStore) Delete(key interface{}) (bool, error) {
	err := ls.DB.Delete([]byte(key.([]byte)), nil)
	if err != nil {
		return false, nil
	}
	return true, nil
}

func (ls *LeveldbStore) SetBatch(keylist []interface{}, valuelist []interface{}) (bool, error) {
	if len(keylist) != len(valuelist) {
		return false, errors.New("key and value not match")
	}
	if len(keylist) == 0 && len(valuelist) == 0 {
		return true, nil
	}
	batchwrite := new(leveldb.Batch)
	for index, keyvalue := range keylist {
		batchwrite.Put([]byte(keyvalue.([]byte)), []byte(valuelist[index].([]byte)))
	}

	err := ls.DB.Write(batchwrite, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ls *LeveldbStore) GetBatch(startkey interface{}, endkey interface{}) ([]interface{}, bool, error) {
	iter := ls.DB.NewIterator(&util.Range{Start: []byte(startkey.([]byte)), Limit: []byte(endkey.([]byte))}, nil)
	var datas []interface{}
	for iter.Next() {
		datas = append(datas, iter.Value())
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		return nil, false, err
	}
	return datas, true, nil
}

func (ls *LeveldbStore) OpenDB(path string) (bool, error) {
	var err error
	ls.DB, err = leveldb.OpenFile(path, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (ls *LeveldbStore) CloseDB() (bool, error) {
	err := ls.DB.Close()
	if err != nil {
		return false, err
	}
	return true, nil
}
