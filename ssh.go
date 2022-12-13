package ssh

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

// TCPTimeout is the default number of seconds to wait to complete a TCP
// connection.
var TCPTimeout = 300 * time.Second

// Run wraps the ssh.Session.Run command with sensible, stand-alone
// defaults. This function has no dependencies on any underlying ssh
// host installation making it idea for light-weight, remote ssh calls.
//
// Run combines several steps. First, a client secure shell connection
// is Dialed to the target (user@host:PORT) using the private PEM user
// key (ukey) and public host key in authorized_keys format (hkey,
// usually ecdsa-sha2-nistp256). Run then attempts to create a Session
// calling Run on it to execute the passed cmd feeding it any standard
// input (in) provided.  The standard output, standard error are then
// buffered and returned as strings. The exit value is captured in err
// for any exit code other than 0. See the ssh.Session.Run method for
// more information.
//
// Note that there are no limitations on the size of input and output
// meaning Run should only be used when calling remote commands that can
// be trusted not to produce too much output.
func Run(target string, ukey, hkey []byte, cmd, in string) (stdout, stderr string, err error) {

	t := strings.Split(target, "@")
	if len(t) != 2 {
		err = fmt.Errorf(`invalid target: %q`, target)
		return
	}
	user := t[0]
	addr := t[1]

	signer, err := ssh.ParsePrivateKey(ukey)
	if err != nil {
		return
	}

	hostkey, _, _, _, err := ssh.ParseAuthorizedKey(hkey)
	if err != nil {
		return
	}

	hostpub, err := ssh.ParsePublicKey(hostkey.Marshal())
	if err != nil {
		return
	}

	client, err := ssh.Dial(`tcp`, addr, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.FixedHostKey(hostpub),
		Timeout:         TCPTimeout,
	})
	if err != nil {
		return
	}

	sess, err := client.NewSession()
	if err != nil {
		return
	}

	if in != "" {
		sess.Stdin = strings.NewReader(in)
	}

	_out := new(strings.Builder)
	_err := new(strings.Builder)
	sess.Stdout = _out
	sess.Stderr = _err

	err = sess.Run(cmd)
	stdout = _out.String()
	stderr = _err.String()

	return

}

// ------------------------------- User -------------------------------

// we use an interface for flexibility and to allow the work to create
// a signer to only be needed once upon creation

type User struct {
	Name   string
	Key    []byte // original pemkey
	Signer ssh.Signer
}

func NewUser(name string, pemkey []byte) (*User, error) {
	var err error
	u := new(User)
	u.Name = name
	u.Key = pemkey
	u.Signer, err = ssh.ParsePrivateKey(pemkey)
	if err != nil {
		return u, err
	}
	return u, nil
}

// ------------------------------- Host -------------------------------

// we use an interface for flexibility and to allow the work to create
// a signer to only be needed once upon creation

type Host struct {
	Addr    string        // name or IP
	Auth    []byte        // authorized_hosts format
	Netkey  ssh.PublicKey // RFC 4235, section 6.6
	Pubkey  ssh.PublicKey // suitable for ssh.FixedHostkey
	Comment string        // authorized_hosts comment
	Options []string      // authorized_hosts options
}

func NewHost(addr string, authkey []byte) (*Host, error) {
	var err error
	host := new(Host)
	host.Addr = addr
	host.Auth = authkey

	host.Netkey, host.Comment, host.Options, _, err = ssh.ParseAuthorizedKey(authkey)
	if err != nil {
		return host, err
	}

	// required since host.net (also ssh.PublicKey) is in RFC format
	// (which fails for ssh.FixedHostKey)

	host.Pubkey, err = ssh.ParsePublicKey(host.Netkey.Marshal())
	if err != nil {
		return host, err
	}

	return host, nil
}

// -------------------------- MultiHostClient -------------------------

type MultiHostClient struct {
	User     *User
	Hosts    []*Host
	Timeout  time.Duration
	Attempts int

	last int
}

func (c MultiHostClient) assert() {
	switch {
	case c.User == nil:
		panic(`undefined User`)
	case c.User.Name == "":
		panic(`undefined User.Name`)
	case c.User.Signer == nil:
		panic(`undefined User.Signer`)
	case c.Hosts == nil:
		panic(`undefined Hosts`)
	case c.Timeout == 0:
		panic(`Timeout cannot be 0`)
	case c.Attempts == 0:
		panic(`Attempts cannot be 0`)
	}
}

// Run attempts to dial a random host from Hosts and waits the Timeout
// duration for a TCP connection before moving to the next host in Hosts
// and attempting to Dial it repeating the cycle once until number of
// Attempts is reached. The total amount of time to wait is therefore
// equal to Timeout * Attempts. The first host to respond to Dial is used.
// Note that the err returned by a command does not cause additional
// attempts, only failed Dail attempts. Panics if any User, Hosts,
// Timeout, or Attempts is undefined.
func (c *MultiHostClient) Run(cmd, in string) (stdout, stderr string, err error) {
	c.assert()

	c.last = rand.Intn(len(c.Hosts))
	host := c.Hosts[c.last]
	var client *ssh.Client

	// keep trying until attempt exhausted
	for attempts := 0; attempts < c.Attempts; attempts++ {

		client, err = ssh.Dial(`tcp`, host.Addr, &ssh.ClientConfig{
			User:            c.User.Name,
			Auth:            []ssh.AuthMethod{ssh.PublicKeys(c.User.Signer)},
			HostKeyCallback: ssh.FixedHostKey(host.Pubkey),
			Timeout:         c.Timeout,
		})

		// error during dial
		if err != nil {
			if c.last == len(c.Hosts)-1 {
				c.last = 0
			} else {
				c.last++
			}
		}

		host = c.Hosts[c.last]

	}

	if client == nil {
		return
	}

	// successful dialup
	var sess *ssh.Session
	sess, err = client.NewSession()
	if err != nil {
		return
	}

	if in != "" {
		sess.Stdin = strings.NewReader(in)
	}

	_out := new(strings.Builder)
	_err := new(strings.Builder)
	sess.Stdout = _out
	sess.Stderr = _err

	err = sess.Run(cmd)
	stdout = _out.String()
	stderr = _err.String()

	return
}
