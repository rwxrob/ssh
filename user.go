package ssh

import (
	"strings"

	"golang.org/x/crypto/ssh"
)

type User struct {
	Name   string
	Key    []byte // original pem
	Signer ssh.Signer
}

// String implements fmt.Stringer as Name.
func (u User) String() string { return u.Name }

// NewUser creates and initializes a new user with a given name using
// the passed PEM key string to parse the private key and
// derive the Signer (which is essential for making connections).
func NewUser(name, pem string) (*User, error) {
	var err error
	u := new(User)
	u.Name = name
	u.Key = []byte(strings.TrimSpace(pem))
	u.Signer, err = ssh.ParsePrivateKey(u.Key)
	if err != nil {
		return u, err
	}
	return u, nil
}
