package remote

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

// SSH sets up a secure shell connection to the target (user@host:PORT)
// and uses public key method from private key in pem ukey
// argument. The hkey parameter must be in "authorized_keys" format for
// hosts allowed (usually ecdsa-sha2-nistp256). It then attempts to
// create a Session calling Output to return only the standard output of
// that command.
func SSH(target string, ukey, hkey []byte, cmd, in string) (stdout, stderr string, err error) {

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
