package service_test

import (
	"testing"
	"github.com/hashicorp/consul/api"
	"github.com/cheebo/consul-utils/service"
	"github.com/cheebo/consul-utils/types"
	"github.com/stretchr/testify/assert"
)
/*
Consul Service Configuration for test:
  {
    "name": "microservice",
    "tags": ["micro"],
    "address": "127.0.0.1",
    "port": 35565
  },
 */

func TestGetServiceAddr(t *testing.T) {
	assert := assert.New(t)

	opt := types.QueryOptions{
		Datacenter: "dc1",
		Token: "",
	}

	config := &api.Config{Address: "localhost:8500", Scheme: "http", Token: ""}
	client, err := api.NewClient(config)

	// Service does not exist
	services, err := service.GetServiceAddr(client, "service", "tag", opt)

	assert.EqualError(err, service.ServiceError.Error(), "GetServiceAddr should return error")
	assert.True(len(services) == 0, "GetServiceAddr should return empty list")

	// Service exists, but tag does not
	services, err = service.GetServiceAddr(client, "microservice", "tag", opt)

	assert.EqualError(err, service.ServiceError.Error(), "GetServiceAddr should return error")
	assert.True(len(services) == 0, "GetServiceAddr should return empty list")

	// Service and tag exist
	services, err = service.GetServiceAddr(client, "microservice", "micro", opt)

	assert.NoError(err)
	assert.True(len(services) != 0, "GetServiceAddr should return list with one item")
	println(services[0].ServiceID)
	if len(services) == 1 {
		assert.Exactly(services[0], &service.ServiceAddr{
			Addr: "127.0.0.1",
			ServiceID: "microservice",
			ServiceAddr: "127.0.0.1",
			ServicePort: 35565,
		})
	}

}
