package consul_utils

import (
	"errors"

	"github.com/cheebo/consul-utils/types"
	"github.com/hashicorp/consul/api"
)

type ServiceAddr struct {
	Addr        string
	ServiceID   string
	ServiceAddr string
	ServicePort int
}

var (
	ServiceError = errors.New("Service not found")
)

func GetServiceAddr(client *api.Client, service, tag string, opt types.QueryOptions) (addrs []*ServiceAddr, err error) {
	services, _, err := client.Catalog().Service(service, tag, &api.QueryOptions{
		Datacenter:        opt.Datacenter,
		Token:             opt.Token,
		RequireConsistent: true,
	})
	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return addrs, ServiceError
	}

	for _, s := range services {
		addrs = append(addrs, &ServiceAddr{
			Addr:        s.Address,
			ServiceID:   s.ServiceID,
			ServiceAddr: s.ServiceAddress,
			ServicePort: s.ServicePort,
		})
	}

	return addrs, nil
}
