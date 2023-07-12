package multihost

import (
	C "github.com/rwxrob/ssh/internal/config"
)

type Client struct {
	timeout int
}

type MultiHostClient interface {
	Timeout() int
}

// uncomment me to validate interface implementations
// var _ MultiHostClient = new(Client)

func (c Client) Timeout() int { return c.timeout }

func (c *Client) Set(cm C.Map) error {
	if v, has := cm[`timeout`].(int); has {
		c.timeout = v
	}
	return nil
}
