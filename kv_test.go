package consul_utils_test

import (
	"testing"
	"github.com/hashicorp/consul/api"
	cu "github.com/cheebo/consul-utils"
	"github.com/stretchr/testify/assert"
)

/*
 Consul KV path to test:
 Path: /test/config
 Value: {"key": "value"}
 */

const (
	path = "/test/config"
	path2 = "test/config2"
	value = `{"key": "value"}`
	value2 = `{"k": "v"}`
)

func TestGetKV(t *testing.T) {
	assert := assert.New(t)

	opt := cu.QueryOptions{
		Datacenter: "dc1",
		Token: "",
	}

	config := &api.Config{Address: "localhost:8500", Scheme: "http", Token: ""}
	client, err := api.NewClient(config)

	val, err := cu.GetKV(client, "/wrong/path", opt)

	assert.NoError(err, "GetKV should not return error")
	assert.Empty(val, "GetKV should return empty string")


	val, err = cu.GetKV(client, path, opt)

	assert.NoError(err, "GetKV should not return error")
	assert.Exactly(value, val, "GetKV returns incorrect value")

}


func TestPutKV(t *testing.T) {
	assert := assert.New(t)

	opt := cu.QueryOptions{
		Datacenter: "dc1",
		Token: "",
	}

	config := &api.Config{Address: "localhost:8500", Scheme: "http", Token: ""}
	client, err := api.NewClient(config)

	// incorrect path
	ok, err := cu.PutKV(client, "/wrong/path/to/key", value2, opt)
	assert.False(ok, "PutKV should return false on incorrect path")
	assert.Error(err, "PutKV should return error on incorrect path")

	cu.Del(client, path2, opt)

	// correct data
	ok, err = cu.PutKV(client, path2, value2, opt)
	assert.True(ok, "PutKV should return true")
	assert.NoError(err, "PutKV should not return error")

	// correct data, duplicate update
	ok, err = cu.PutKV(client, path2, value2, opt)
	assert.False(ok, "PutKV should return false on existed data")
	assert.NoError(err, "PutKV should not return error")

	// delete key
	ok, err = cu.Del(client, path2, opt)
	assert.True(ok, "Del should return true")
	assert.NoError(err, "Del should not return error")
}
