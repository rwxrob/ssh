package ssh

import (
	"encoding/json"
	"strings"

	"golang.org/x/crypto/ssh"
	"gopkg.in/yaml.v3"
)

type User struct {
	Name string
	Key  string

	signer ssh.Signer
}

// Signer is the crypto/ssh.Signer created from the Key when NewUser is
// called.
func (u User) Signer() ssh.Signer { return u.signer }

// JSON is a convenience method for marshaling as JSON string.
// Marshaling errors return a "null" string.
func (u User) JSON() string {
	byt, err := json.Marshal(u)
	if byt == nil || err != nil {
		return "null"
	}
	return string(byt)
}

// YAML is a convenience method for marshaling as YAML string.
// Marshaling errors return a "null" string.
func (u User) YAML() string {
	byt, err := yaml.Marshal(u)
	if byt == nil || err != nil {
		return "null"
	}
	return string(byt)
}

// NewUser creates and initializes a new user with a given name using
// the passed PEM key string to parse the private key and
// derive the Signer (which is essential for making connections).
func NewUser(name, pem string) (*User, error) {
	var err error
	u := new(User)
	u.Name = name
	u.Key = strings.TrimSpace(pem)
	u.signer, err = ssh.ParsePrivateKey([]byte(u.Key))
	if err != nil {
		return u, err
	}
	return u, nil
}
