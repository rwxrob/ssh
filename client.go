package ssh

import (
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

// DefaultTCPTimeout is the default number of seconds to wait to complete a TCP
// connection. See Client.Timeout also.
var DefaultTCPTimeout = 300 * time.Second

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

	// User contains the name of the user on the target server and
	// contains the PEM authentication data as well.
	User *User

	// Timeout is the default number of seconds to wait to complete a TCP
	// connection. This is set to DefaultTCPTimeout by NewClient.
	Timeout time.Duration

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
	c := new(Client)
	c.Port = DefaultPort
	c.Timeout = DefaultTCPTimeout
	return c
}

// Addr returns addr:port.
func (c Client) Addr() string {
	return fmt.Sprintf("%v:%v", c.Host.Addr, c.Port)
}

// String fulfills the fmt.Stringer interface by providing the unique
// key identifier for Controller.ClientMap. See the String()
// implementations of Host and User as well.
func (c Client) String() string {
	return fmt.Sprintf(`%v@%v:%v`, c.User, c.Host, c.Port)
}

// Connect creates a new ssh.Client and assigns it to Host.SSH. If the
// Host has an Auth will attempt to authenticate the host and the
// c.User.Signer is always used.
func (c *Client) Connect() error {
	var err error
	c.SSH, err = ssh.Dial(`tcp`, c.Addr(), &ssh.ClientConfig{
		User:            c.User.Name,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(c.User.Signer)},
		HostKeyCallback: c.Host.KeyCallback(),
		Timeout:         c.Timeout,
	})
	return err
}

func (c *Client) Run(cmd, stdin string) (stdout, stderr string, err error) {

	if c.SSH == nil {
		err = c.Connect()
		if err != nil {
			return
		}
	}

	var sess *ssh.Session
	sess, err = c.SSH.NewSession()
	if err != nil {
		// TODO attempt (once) to reconnect
		return
	}

	if stdin != "" {
		sess.Stdin = strings.NewReader(stdin)
	}

	_out := new(strings.Builder)
	_err := new(strings.Builder)
	sess.Stdout = _out
	sess.Stderr = _err

	sess.Run(cmd)
	stdout = _out.String()
	stderr = _err.String()

	return
}
