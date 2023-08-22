package ssh

import (
	"fmt"
	"math/rand"
	"sync"
)

// ClientMap is a collection of one or more Clients each mapped to
// a specific target string key corresponding to common ssh command
// targets (user@host:22). A Controller uses an embedded ClientMap when given instructions that
// apply to one, all, or a filtered subset of Clients contained in the
// ClientMap but only one Client per target is allowed. Clients must be added
// with Add to keep ClientMap methods safe for concurrency although
// the internal Map is exported for convenience as well. Note that while
// it is possible to instantiate a ClientMap statically it is generally
// preferred to do so dynamically at runtime at which time each
// Client.Connect connection can be established and checked as well.
type ClientMap struct {
	sync.RWMutex

	Map map[string]*Client

	// Keys contains the cache of targets added. Be careful when modifying
	// this directly so that it does not get out of synch with Map.
	Keys []string
}

// NewClientMap instantiates and returns a new ClientMap creating the
// embedded Map suitable for use directly.
func NewClientMap() *ClientMap {
	m := new(ClientMap)
	m.Map = map[string]*Client{}
	m.Keys = []string{}
	return m
}

// Add adds an existing Client to the internal Map in a way that is safe
// for concurrency. The String version of the client (ex: user@host:22)
// is used as the key. No attempt to connect the client is made. Returns
// an error if a target (key) for client already exists.
func (m *ClientMap) Add(c *Client) error {
	m.RWMutex.Lock()
	defer m.RWMutex.Unlock()
	key := c.String()
	_, has := m.Map[key]
	if has {
		return fmt.Errorf(`%v already exists`, key)
	}
	m.Map[key] = c
	m.Keys = append(m.Keys, key)
	return nil
}

// Random returns a random client from the Map based on random selection
// from Keys.
func (m *ClientMap) Random() *Client {
	return m.Map[m.Keys[rand.Intn(len(m.Keys))]]
}
