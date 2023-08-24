package ssh

import (
	"encoding/json"
	"io"
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

// NewUser creates and initializes a new user (by calling Init()) with
// a given name using the passed PEM key string to parse the private key
// and derive the Signer() (which is essential for making connections).
func NewUser(name, key string) (*User, error) {
	u := new(User)
	u.Name = name
	u.Key = strings.TrimSpace(key)
	return u, u.Init()
}

// NewUserFromYAML reads the Name and Key from YAML instead of passing.
// See NewUser.
func NewUserFromYAML(r io.Reader) (*User, error) {
	u := new(User)
	byt, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(byt, u)
	if err != nil {
		return nil, err
	}
	return u, u.Init()
}

// NewUserFromJSON reads the Name and Key from JSON instead of passing.
// See NewUser.
func NewUserFromJSON(r io.Reader) (*User, error) {
	u := new(User)
	byt, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byt, u)
	if err != nil {
		return nil, err
	}
	return u, u.Init()
}

// Init sets the internal Signer() by parsing the Key field returning
// any error encountered.
func (u *User) Init() error {
	var err error
	u.signer, err = ssh.ParsePrivateKey([]byte(u.Key))
	return err
}
