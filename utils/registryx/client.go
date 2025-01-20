package registryx

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
)

func NewConsulResolverBuilder(address string) ConsulResolverBuilder {
	return ConsulResolverBuilder{consulAddress: address}
}

type ConsulResolverBuilder struct {
	consulAddress string
}

func (c ConsulResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	consulResolver, err := newConsulResolver(c.consulAddress, target, cc)
	if err != nil {
		return nil, err
	}
	consulResolver.resolve()
	return consulResolver, nil
}

func (c ConsulResolverBuilder) Scheme() string {
	return "consul"
}

func newConsulResolver(address string, target resolver.Target, cc resolver.ClientConn) (ConsulResolver, error) {
	var reso ConsulResolver
	client, err := consulapi.NewClient(&consulapi.Config{Address: address})
	if err != nil {
		return reso, err
	}
	return ConsulResolver{
		target: target,
		cc:     cc,
		client: client,
	}, nil
}

type ConsulResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
	client *consulapi.Client
}

func (c ConsulResolver) resolve() {
	service := c.target.URL.Opaque
	services, _, err := c.client.Catalog().Service(service, "", nil)
	if err != nil {
		c.cc.ReportError(err)
		return
	}
	var adds []resolver.Address
	for _, catalogService := range services {
		adds = append(adds, resolver.Address{Addr: fmt.Sprintf("%s:%d", catalogService.Address, catalogService.ServicePort)})
	}

	c.cc.UpdateState(resolver.State{
		Addresses: adds,
		// 轮询策略
		ServiceConfig: c.cc.ParseServiceConfig(
			`{"loadBalancingPolicy":"round_robin"}`),
	})
}

func (c ConsulResolver) ResolveNow(options resolver.ResolveNowOptions) {
	c.resolve()
}

func (c ConsulResolver) Close() {

}
