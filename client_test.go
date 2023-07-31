package ssh_test

import (
	"fmt"

	"github.com/rwxrob/ssh"
)

func ExampleClient_String() {

	host := &ssh.Host{Addr: `host`}
	user := &ssh.User{Name: `user`}
	client := &ssh.Client{Host: host, User: user, Port: 22}

	fmt.Println(client)
	fmt.Println(client.String())

	user.Name = `other`
	fmt.Println(client)

	client.Host = nil
	fmt.Println(client)

	client.User = nil
	fmt.Println(client)

	client.Port = 0
	fmt.Println(client)

	// Output:
	// user@host:22
	// user@host:22
	// other@host:22
	// other@<nil>:22
	// <nil>@<nil>:22
	// <nil>@<nil>:0
}

func ExampleClient_Run() {

	client := ssh.NewClient()
	client.User, _ = ssh.NewUser(
		`user`, `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
-----END OPENSSH PRIVATE KEY-----
`)
	client.Host, _ = ssh.NewHost(
		`localhost`, `localhost ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=`,
	)

	// LEFTOFF
	client.Run(`echo hello`)

	// Output:
	// hello
}