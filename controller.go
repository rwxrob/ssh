package ssh

import (
	"encoding/json"
	"io"
	"math/rand"
	"net"

	"gopkg.in/yaml.v3"
)

// Controller is responsible for coordinating work requests destined for
// the ssh target servers as contained in its list of Clients.
type Controller struct {
	Clients []*Client
}

// NewController initializes a new Controller and returns a pointer to
// it. Any clients passed are appended to the internal Client list (but
// not yet connected). See Connect().
func NewController(clients ...*Client) *Controller {
	ctl := new(Controller)
	ctl.Clients = make([]*Client, len(clients))
	for n, client := range clients {
		ctl.Clients[n] = client
	}
	return ctl
}

// NewControllerFromYAML uses io.ReadAll to read all the bytes from the
// io.Reader passed and converts them to a new controller with multiple
// Clients. Note that YAML references are observed meaning that a single
// host entry can be used for multiple clients.
//
// Also see NewController.
func NewControllerFromYAML(r io.Reader) (*Controller, error) {
	ctl := new(Controller)
	byt, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(byt, ctl)
	return ctl, err
}

// NewControllerFromJSON uses io.ReadAll to read all the bytes from the
// io.Reader passed and converts them to a new controller with multiple
// Clients. Also see NewController.
func NewControllerFromJSON(r io.Reader) (*Controller, error) {
	ctl := new(Controller)
	byt, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byt, ctl)
	return ctl, err
}

// Connect synchronously calls Connect on all Clients ensuring that all
// Clients that can have successfully connected before returning. No
// attempt at error checking for successful connections is attempted but
// the Connected() and LastError() properties of each client can be can
// be examined for information. A reference to self is returned as
// convenience allowing this to be changed onto NewController.
func (c *Controller) Connect() *Controller {
	for _, client := range c.Clients {
		client.Connect()
	}
	return c
}

// JSON is a convenience method for marshaling as JSON string.
// Marshaling errors return a "null" string.
func (c Controller) JSON() string {
	byt, err := json.Marshal(c)
	if byt == nil || err != nil {
		return "null"
	}
	return string(byt)
}

// YAML is a convenience method for marshaling as YAML string.
// Marshaling errors return a "null" string.
func (c Controller) YAML() string {
	byt, err := yaml.Marshal(c)
	if byt == nil || err != nil {
		return "null"
	}
	return string(byt)
}

// RandomClient returns a random active client from the Clients list
// skipping any that are not Connected. Returns nil if no Connected
// clients are available.
func (c *Controller) RandomClient() *Client {
	var tried int
	count := len(c.Clients)
	if count == 0 {
		return nil
	}
	n := rand.Intn(count)
	for {
		client := c.Clients[n]
		if client.connected {
			return client
		}
		tried += 1
		if tried > count {
			return nil
		}
		n += 1
		if n >= count {
			n = 0
		}
	}
	return nil
}

// RunOnAny calls client.Run on a random client from the Clients list.
// If error returned is of type net.OpError the client.Connected is set
// to false and the next client in the ClientMap.Keys order is attempted
// while Connect is called on the Client with the net.OpError from
// a separate goroutine (which, if successful, restores the Connected
// status to true). If none of the clients are connected then an
// AllUnavailable error is returned.
func (c *Controller) RunOnAny(cmd, stdin string) (stdout, stderr string, err error) {
TOP:
	client := c.RandomClient()
	if client == nil {
		err = AllUnavailable{}
		return
	}
	stdout, stderr, err = client.Run(cmd, stdin)
	if _, is := err.(*net.OpError); is {
		client.connected = false
		go client.Connect()
		goto TOP
	}
	return
}
