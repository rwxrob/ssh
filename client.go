package ssh

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

// Client encapsulates the ssh.Client struct and associates a single
// user, host, and port number to target for an ssh server connection
// adding a Connect method which implicitly dials up a connection
// storing it as the internal ssh.Client when called.
// The combination in standard ssh target notation (see String()) is
// used as the key string in the Controller.ClientMap when dealing with
// multiple targets.
type Client struct {

	// Host contains the host name or IP address (Addr) along with any
	// authorization credentials and other data that would normally be
	// contained in authorized_hosts.
	Host *Host

	// Port is simply the port the ssh server is listening on on the
	// target server.
	Port int

	// User is the name of the user on the target server and contains the
	// PEM authentication data as well.
	User *User

	// SSH contains the internal ssh.Client if one has been created. Note
	// that this internal Client is *not* created until at least one call
	// to Connect has been made.
	SSH *ssh.Client
}

// DefaultPort for SSH connetions. See NewClient.
var DefaultPort = 22

// NewClient simple sets the Port to the DefaultPort and returns
// a pointer to a new Client struct. No other initialization is done.
func NewClient() *Client {
	cl := new(Client)
	cl.Port = DefaultPort
	return cl
}

// String fulfills the fmt.Stringer interface by providing the unique
// key identifier for Controller.ClientMap. See the String()
// implementations of Host and User as well.
func (c Client) String() string {
	return fmt.Sprintf(`%v@%v:%v`, c.User, c.Host, c.Port)
}

func (c *Client) Connect() error {
	// TODO
	return nil
}

func (c *Client) Run(cmd, stdin string) (stdout, stderr string, err error) {
	// TODO
	return
}
