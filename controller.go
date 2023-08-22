package ssh

import (
	"encoding/json"
	"fmt"
)

// Controller is responsible for coordinating work requests destined for
// the ssh target servers as contained in the embedded ClientMap. Note
// that the methods of ClientMap are available to Controller directly
// through embedded delegation.
type Controller struct {
	ClientMap
}

// NewController initializes a new Controller and returns a pointer to
// it. Any clients passed to it are added to the internal ClientMap in
// order. Note that clients with conflicting target (String)
// representations clobber previously added ones. Clients can also be
// added after the fact with Add (delegated to ClientMap.Add).
func NewController(clients ...*Client) *Controller {
	ctl := new(Controller)
	ctl.ClientMap = *NewClientMap()
	for _, client := range clients {
		ctl.ClientMap.Add(client)
	}
	return ctl
}

// String fulfills the fmt.Stringer interface by first printing the
// clientCount and then looping through the embedded clients in the
// ClientMap printing their target identifier strings and a status
// summary for each as a convenience. See Client.String as well.
func (c Controller) String() string {
	var buf string
	buf += fmt.Sprintf("client-count: %v\n", len(c.ClientMap.Map))
	for _, client := range c.ClientMap.Map {
		buf += fmt.Sprintln(client)
	}
	return buf
}

// Print simply prints String() with an additional line return for
// convenience.
func (c *Controller) Print() { fmt.Println(c.String()) }

// JSON does the same String but outputs a JSON string better suited for
// parsing by applications wishing to display the state of the
// Controller. This provides some basic observability into the
// Controller's behavior as well. JSON never has any actual line returns
// allowing the full JSON to be saved to a single line of text (for
// logging, etc.) Note that in the rare case the Controller itself (the
// receiver) is nil that this will output null (not <nil>, no quotes,
// which is valid JSON).
func (c *Controller) JSON() string {
	if c == nil {
		return "null"
	}
	tmp := struct {
		clientCount int `json:"client-count"`
	}{len(c.ClientMap.Map)}
	buf, err := json.Marshal(tmp)
	if err != nil {
		return "null"
	}
	return string(buf)
}

// RunOnAny calls client.Run on a random client in the ClientMap.
func (c *Controller) RunOnAny(cmd, stdin string) (string, string, error) {
	return c.Random().Run(cmd, stdin)
}
