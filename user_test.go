package ssh_test

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/rwxrob/ssh"
)

func ExampleUser() {

	user, err := ssh.NewUser(
		`user`, `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
-----END OPENSSH PRIVATE KEY-----
`)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
	//	fmt.Println(user.Signer)

	// Output:
	// user
}

func ExampleUser_from_YAML() {
	byt, _ := os.ReadFile(`testdata/user.yaml`)
	user := new(ssh.User)
	err := yaml.Unmarshal(byt, user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	// Output:
	// user
}
