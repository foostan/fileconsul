package fileconsul

import (
	"time"

	consulapi "github.com/armon/consul-api"
)

type Client struct {
	ConsulClient *consulapi.Client
}

type ClientConfig struct {
	ConsulAddr string
	ConsulDC   string
	Timeout    time.Duration
}

type ConsulAgentInfo struct {
	Name   string
	Port   float64
	Addr   string
	Status float64
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

func (c *Client) GetKV(key string) (*consulapi.KVPair, error) {
	pair, _, err := c.ConsulClient.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}

	return pair, nil
}

func (c *Client) ListKV(prefix string) (consulapi.KVPairs, error) {
	pairs, _, err := c.ConsulClient.KV().List(prefix, nil)
	if err != nil {
		return nil, err
	}

	return pairs, nil
}

func (c *Client) PutKV(prefix string, value []byte) error {
	p := &consulapi.KVPair{Key: prefix, Value: value}
	_, err := c.ConsulClient.KV().Put(p, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteKV(prefix string) error {
	_, err := c.ConsulClient.KV().DeleteTree(prefix, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ConsulAgentInfo() (*ConsulAgentInfo, error) {
	info, err := c.ConsulClient.Agent().Self()
	if err != nil {
		return nil, err
	}

	return &ConsulAgentInfo{
		Name:   info["Member"]["Name"].(string),
		Port:   info["Member"]["Port"].(float64),
		Addr:   info["Member"]["Addr"].(string),
		Status: info["Member"]["Status"].(float64),
	}, nil
}
