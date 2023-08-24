package ssh

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

// DefaultPort for SSH connetions. See NewClient.
var DefaultPort = 22

// DefaultTCPTimeout is the default number of seconds to wait to complete a TCP
// connection. See Client.Timeout also.
var DefaultTCPTimeout = 300 * time.Second

// Client encapsulates an internal ssh.Client and associates a single
// user, host, and port number to target for specific ssh server connection
// adding a Connect method which implicitly dials up a connection
// setting Connected() and internally caching the client as SSH().
// Client may be safely marshaled to/from YAML and  JSON directly.
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
	// connection. If unset DefaultTCPTimeout is used.
	Timeout time.Duration

	sshclient *ssh.Client
	connected bool
	lasterror error
}

// SSHClient returns a pointer to the internal ssh.Client used for all
// connections and sessions.
func (c Client) SSHClient() *ssh.Client { return c.sshclient }

// Connected returns the last connection state of the internal SSH
// client. This is set to true on Connect.
func (c Client) Connected() bool { return c.connected }

// LastError returns the last error (if any) from an attempt to Connect.
func (c Client) LastError() error { return c.lasterror }

// Addr returns addr:port suitable for use in TCP/IP connection strings.
// See Dest when user wanted. If the Host is nil returns an empty
// string.
func (c Client) Addr() string {
	if c.Host == nil {
		return ``
	}
	return fmt.Sprintf("%v:%v", c.Host.Addr, c.Port)
}

// Dest returns a user@host:port destination suitable for inclusion in a full URI
// with schema or the common short form. Empty User or User.Name omits it from
// the Dest string. Port, however, is always included.
func (c Client) Dest() string {
	str := c.Addr()
	if len(str) == 0 {
		return ``
	}
	if c.User != nil || len(c.User.Name) == 0 {
		return c.User.Name + `@` + str
	}
	return str
}

// JSON is a convenience method for marshaling as JSON string.
// Marshaling errors return a "null" string.
func (c Client) JSON() string {
	byt, err := json.Marshal(c)
	if byt == nil || err != nil {
		return "null"
	}
	return string(byt)
}

// YAML is a convenience method for marshaling as YAML string.
// Marshaling errors return a "null" string.
func (c Client) YAML() string {
	byt, err := yaml.Marshal(c)
	if byt == nil || err != nil {
		return "null"
	}
	return string(byt)
}

// NewClient returns a pointer to a new Client with defaults set
// (DefaultPort, DefaultTCPTimeout).
func NewClient() *Client {
	c := new(Client)
	c.Port = DefaultPort
	c.Timeout = DefaultTCPTimeout
	return c
}

// NewClientFromYAML uses io.ReadAll to read all the bytes from the
// io.Reader passed and converts them to a new Client. Note that YAML
// references are observed meaning that a single Client entry can have
// its Host or User set to a reference pointing elsewhere in the YAML
// file. The Port and TCPTimeout are set to defaults if unset. Also see
// NewClient.
func NewClientFromYAML(r io.Reader) (*Client, error) {
	c := new(Client)
	byt, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(byt, c)
	if c.Port == 0 {
		c.Port = DefaultPort
	}
	if c.Timeout == 0 {
		c.Timeout = DefaultTCPTimeout
	}
	return c, err
}

// NewClientFromJSON uses io.ReadAll to read all the bytes from the
// io.Reader passed and converts them to a new Client. The Port and
// Timeout are set to defaults if unset. Also see NewClient.
func NewClientFromJSON(r io.Reader) (*Client, error) {
	c := new(Client)
	byt, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byt, c)
	if c.Port == 0 {
		c.Port = DefaultPort
	}
	if c.Timeout == 0 {
		c.Timeout = DefaultTCPTimeout
	}
	return c, err
}

// Connect creates a new ssh.Client and assigns it to Host.SSH. If the
// Host has an Auth will attempt to authenticate the host and the
// c.User.Signer is always used. Connected is set to true if successful.
// Always reinitializes a new connection (see SSH()) even if already
// Connected.
func (c *Client) Connect() error {
	var err error
	c.sshclient, err = ssh.Dial(`tcp`, c.Addr(), &ssh.ClientConfig{
		User:            c.User.Name,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(c.User.Signer())},
		HostKeyCallback: c.Host.KeyCallback(),
		Timeout:         c.Timeout,
	})
	if err == nil {
		c.connected = true
	} else {
		c.connected = false
		c.lasterror = err
	}
	return err
}

// Run sends the command with optional standard input to the currently
// open client SSH target as a new ssh.Session. If the SSH connection
// has not yet been established (c.SSH is nil) Connect is called to
// establish a new client connection. Run returns an error if one is
// generated by the session.Run call or if a new session could not be
// created (including attempting a new session on a timed-out connection
// or one that has been closed for any reason.) It is the responsibility
// to respond to such errors according to controller policy and
// associated method calls.
func (c *Client) Run(cmd, stdin string) (stdout, stderr string, err error) {

	if c.sshclient == nil {
		err = c.Connect()
		if err != nil {
			return
		}
	}

	var sess *ssh.Session
	sess, err = c.sshclient.NewSession()
	if err != nil {
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
	sess.Close()

	return
}
