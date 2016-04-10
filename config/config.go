package config

import "Antleg/tools"
import "github.com/coreos/etcd/client"

type StoreDb struct {
	Dbpath      string
	Replication bool
}

func (s *StoreDb) LoadCondig() {

}

type EtcdService struct {
	EtcdAddress string
}

func (s *EtcdService) LoadCondig() {

}

func init() {

}
