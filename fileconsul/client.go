package fileconsul

import (
	"time"

	consulapi "github.com/armon/consul-api"
)

type Client struct {
	ConsulClient *consulapi.Client
}

type ClientConfig struct {
	ConsulAddr     string
	ConsulDC       string
	Timeout        time.Duration
}

func NewClient(config *ClientConfig) (*Client, error) {
	kvConfig := consulapi.DefaultConfig()
	kvConfig.Address = config.ConsulAddr
	kvConfig.Datacenter = config.ConsulDC

	consulClient, err := consulapi.NewClient(kvConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		ConsulClient: consulClient,
	}, nil
}

func (c *Client) GetKVByKeyprefix(prefix string) (consulapi.KVPairs, error) {
	pairs, _, err := c.ConsulClient.KV().List(prefix, nil)
	if err != nil {
		return nil, err
	}

	return pairs, nil
}

func (c *Client) PutKVByKeyprefix(prefix string, value string) error {
	p := &consulapi.KVPair{Key: prefix, Value: []byte(value)}
	_, err := c.ConsulClient.KV().Put(p, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteKVByKeyprefix(prefix string) error {
	_, err := c.ConsulClient.KV().DeleteTree(prefix, nil)
	if err != nil {
		return err
	}

	return nil
}
