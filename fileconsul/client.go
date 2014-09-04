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

func (c *Client) GetFileStatus(prefix string, baseDir string) error {
	pairs, meta, err := c.ConsulClient.KV().List(prefix, nil)
	if err != nil {
		return err
	}

	println("got file status of " + baseDir)

	println(meta.LastIndex, meta.LastContact, meta.KnownLeader, meta.RequestTime)

	for _, pair := range pairs {
		println(pair.Key + "/" + string(pair.Value))
		println(pair.CreateIndex, pair.ModifyIndex, pair.LockIndex, pair.Flags, pair.Session)
	}

	return nil
}
