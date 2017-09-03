package consul_utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/cheebo/consul-utils/types"
	"github.com/cheebo/consul-utils/kv"
	"github.com/hashicorp/consul/api"
)

func GetRsaPublicKey(client *api.Client, key string, opt types.QueryOptions) *rsa.PublicKey {
	val, err := kv.GetKV(client, key, opt)
	if err != nil || len(val) == 0 {
		return nil
	}
	block, _ := pem.Decode([]byte(val))
	if block == nil {
		return nil
	}
	if block.Type != "PUBLIC KEY" && block.Type != "RSA PUBLIC KEY" {
		return nil
	}

	pubkeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	rsaPublicKey, ok := pubkeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil
	}
	return rsaPublicKey
}

func GetRsaPrivateKey(client *api.Client, key string, opt types.QueryOptions) *rsa.PrivateKey {
	val, err := kv.GetKV(client, key, opt)
	if err != nil || len(val) == 0 {
		return nil
	}
	block, _ := pem.Decode([]byte(val))
	if block == nil {
		return nil
	}
	if block.Type != "RSA PRIVATE KEY" {
		return nil
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil
	}
	return privateKey
}
