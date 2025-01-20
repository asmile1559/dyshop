package registryx

import (
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

var RegistryAddr string

func init() {
	// TODO: config
	RegistryAddr = "127.0.0.1:8500"
}

type Service struct {
	Node  string
	Addr  string
	Agent *consulapi.AgentService
}

// 注册服务
func (srv Service) Register() error {
	client, err := consulapi.NewClient(&consulapi.Config{Address: RegistryAddr})
	if err != nil {
		return err
	}
	_, err = client.Catalog().Register(&consulapi.CatalogRegistration{
		Node:    srv.Node,
		Address: srv.Addr,
		Service: srv.Agent,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

// 注销服务
func (srv Service) DeRegister() error {
	client, err := consulapi.NewClient(&consulapi.Config{Address: RegistryAddr})
	if err != nil {
		logrus.Error(err)
		return err
	}
	_, err = client.Catalog().Deregister(&consulapi.CatalogDeregistration{
		Node:      srv.Node,
		Address:   srv.Addr,
		ServiceID: srv.Agent.ID,
	}, nil)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
