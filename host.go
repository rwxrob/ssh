package ssh

import (
	"strings"

	"golang.org/x/crypto/ssh"
)

// A Host is one part of a Client.Target (which also includes the User).
// Prefer
type Host struct {

	// Name or IP address (only)
	Addr string `json:"addr"`

	// Complete line taken in the authorized_hosts format (included Addr)
	Auth string `json:"auth"`

	// RFC 4234 (section 6.6). Note that although this is an ssh.PublicKey
	// that the exact format is very different from other ssh.PublicKey
	// values (such as Pubkey below).
	Netkey ssh.PublicKey `json:"-"`

	// Suitable for use with ssh.FixedHostkey. Note this is different than
	// other ssh.PublicKey formats (see comment for Netkey).
	Pubkey ssh.PublicKey `json:"-"`

	// Comment from the line that would appear in the authorized_hosts
	// file along with other data for this host (see Auth, Options).
	Comment string `json:"-"`

	// Options allowed in the authorized_hosts file for this given host.
	// (See Auth and Comment)
	Options []string `json:"-"`
}

// KeyCallback returns ssh.FixedHostKey(h.Pubkey) if Auth is not nil.
// Otherwise, returns ssh.InsecureIgnoreHostKey().
func (h Host) KeyCallback() ssh.HostKeyCallback {
	if h.Auth != "" {
		return ssh.FixedHostKey(h.Pubkey)
	}
	return ssh.InsecureIgnoreHostKey()
}

// String implements the fmt.Stringer interface with the Addr as
// a string.
func (h Host) String() string { return h.Addr }

// The NewHost function creates, initializes, and returns a new Host
// suitable for use in SSH connections. The first argument is the host
// name or IP address or IP address to use as the target (also see
// ClientMap). The optional auth is assumed to match the known_hosts
// format (which can be taken directly from most ~/.ssh/known_hosts
// files). If the auth is not empty it triggers the assignment of the
// remaining Host fields. Otherwise, they remain blank.
func NewHost(addr, auth string) (*Host, error) {
	var err error

	host := new(Host)
	host.Addr = addr
	if len(auth) == 0 {
		return host, nil
	}
	host.Auth = strings.TrimSpace(auth)

	host.Netkey, host.Comment, host.Options, _, err = ssh.ParseAuthorizedKey([]byte(host.Auth))
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
