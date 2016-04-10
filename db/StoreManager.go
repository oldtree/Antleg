package db

import (
	"errors"
)

type StoreConfig struct {
	DBPath string
	Batch  bool
}

func NewStoreConfig(dbpath string, batch bool) *StoreConfig {
	return &StoreConfig{
		DBPath: dbpath,
		Batch:  batch,
	}
}

type StoreFileObject interface {
	Get(key interface{}) (interface{}, bool, error)
	GetBatch(startkey interface{}, endkey interface{}) ([]interface{}, bool, error)
	Set(key interface{}, value interface{}) (bool, error)
	SetBatch(keylist []interface{}, valuelist []interface{}) (bool, error)
	Delete(key interface{}) (bool, error)
	OpenDB(path string) (bool, error)
	CloseDB() (bool, error)
}

type StoreManager struct {
	Store  StoreFileObject
	Config *StoreConfig
}

func (s *StoreManager) Init() (bool, error) {
	if s.Config == nil {
		return false, errors.New("db config not init")
	}
	return true, nil
}

func (s *StoreManager) OpenDB() (bool, error) {
	return s.Store.OpenDB(s.Config.DBPath)
}

func (s *StoreManager) CloseDB() (bool, error) {
	return s.CloseDB()
}

func (s *StoreManager) Set(key interface{}, value interface{}) (bool, error) {
	return s.Store.Set(key, value)
}

func (s *StoreManager) Get(key interface{}) (interface{}, bool, error) {
	return s.Store.Get(key)
}

func (s *StoreManager) SetBatch(keylist []interface{}, valuelist []interface{}) (bool, error) {
	return s.Store.SetBatch(keylist, valuelist)
}

func (s *StoreManager) GetBatch(startkey interface{}, endkey interface{}) ([]interface{}, bool, error) {
	return s.Store.GetBatch(startkey, endkey)
}
