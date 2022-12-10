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
func SSH(target string, ukey, hkey []byte, cmd string) ([]byte, error) {

	t := strings.Split(target, "@")
	if len(t) != 2 {
		return nil, fmt.Errorf(`invalid target: %q`, target)
	}
	user := t[0]
	addr := t[1]

	signer, err := ssh.ParsePrivateKey(ukey)
	if err != nil {
		return nil, err
	}

	pk, _, _, _, err := ssh.ParseAuthorizedKey(hkey)
	if err != nil {
		return nil, err
	}

	pubkey, err := ssh.ParsePublicKey(pk.Marshal())
	if err != nil {
		return nil, err
	}

	conn, err := ssh.Dial(`tcp`, addr, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.FixedHostKey(pubkey),
	})
	if err != nil {
		return nil, err
	}

	sess, err := conn.NewSession()
	if err != nil {
		return nil, err
	}

	return sess.Output(cmd)
}
