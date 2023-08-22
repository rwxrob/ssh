package ssh

import (
	"strings"

	"golang.org/x/crypto/ssh"
)

type User struct {
	Name   string     `json:"name"`
	Key    string     `json:"key"` // original pem
	Signer ssh.Signer `json:"-"`
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
	u.Key = strings.TrimSpace(pem)
	u.Signer, err = ssh.ParsePrivateKey([]byte(u.Key))
	if err != nil {
		return u, err
	}
	return u, nil
}
