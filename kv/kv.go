package kv

import (
	"strings"
	"github.com/hashicorp/consul/api"
	"github.com/cheebo/consul-utils/types"
)

func GetKV(client *api.Client, key string, opt types.QueryOptions) (string, error) {
	q := &api.QueryOptions{
		Datacenter: opt.Datacenter,
		Token: opt.Token,
		RequireConsistent: true,
	}
	kvpair, _, err := client.KV().Get(key, q)
	if err != nil {
		return "", err
	}
	if kvpair == nil {
		return "", nil
	}
	return strings.TrimSpace(string(kvpair.Value)), nil
}

func PutKV(client *api.Client, key, value string, opt types.QueryOptions) (bool, error) {
	p := &api.KVPair{Key: key, Value: []byte(value)}
	ok, _, err := client.KV().CAS(p, &api.WriteOptions{
		Datacenter: opt.Datacenter,
		Token: opt.Token,
	})
	if err != nil {
		return false, err
	}
	return ok, nil
}

func Del(client *api.Client, key string, opt types.QueryOptions) (bool, error) {
	_, err := client.KV().Delete(key, &api.WriteOptions{
		Datacenter: opt.Datacenter,
		Token: opt.Token,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func DelTree(client *api.Client, prefix string, opt types.QueryOptions) (bool, error) {
	_, err := client.KV().DeleteTree(prefix, &api.WriteOptions{
		Datacenter: opt.Datacenter,
		Token: opt.Token,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}