package ssh

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

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

	conn, err := ssh.Dial(`tcp`, addr, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.FixedHostKey(hostpub),
	})
	if err != nil {
		return
	}

	sess, err := conn.NewSession()
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
