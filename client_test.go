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

func ExampleClient_Run_no_Host_Key() {

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
	client.Host, _ = ssh.NewHost(`localhost`, ``)

	stdout, stderr, err := client.Run(`echo hello`, ``)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout)
	fmt.Println(stderr)

	// Output:
	// hello
	//
}

func ExampleClient_Run_with_Host_Key() {

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
	var err error
	client.Host, err = ssh.NewHost(
		`localhost`,
		`|1|r3meMBTG9TZiPoVHg1n+o1N1xJk=|9I891Skl7BcqG/vaT6wXxt6bZUk= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDqzw7+sN4aVOQqgTA5tC9pN/+M0KOcib3lRAGQ+MSKk/4MbdJY2REavrwRetreaIZTkZx4ykTAJ3CeCK45IzsY=`,
	)
	if err != nil {
		fmt.Println(err)
	}

	stdout, stderr, err := client.Run(`echo hello`, ``)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout)
	fmt.Println(stderr)

	// Output:
	// hello
	//
}

func ExampleClient_Run_with_Stdin() {

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
	var err error
	client.Host, err = ssh.NewHost(
		`localhost`,
		`|1|r3meMBTG9TZiPoVHg1n+o1N1xJk=|9I891Skl7BcqG/vaT6wXxt6bZUk= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDqzw7+sN4aVOQqgTA5tC9pN/+M0KOcib3lRAGQ+MSKk/4MbdJY2REavrwRetreaIZTkZx4ykTAJ3CeCK45IzsY=`,
	)
	if err != nil {
		fmt.Println(err)
	}

	stdout, stderr, err := client.Run(`cat`, `i'm a cat`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout)
	fmt.Println(stderr)

	// Output:
	// i'm a cat
	//
}

func ExampleClient_Run_with_Error() {

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
	var err error
	client.Host, err = ssh.NewHost(
		`localhost`,
		`|1|r3meMBTG9TZiPoVHg1n+o1N1xJk=|9I891Skl7BcqG/vaT6wXxt6bZUk= ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBDqzw7+sN4aVOQqgTA5tC9pN/+M0KOcib3lRAGQ+MSKk/4MbdJY2REavrwRetreaIZTkZx4ykTAJ3CeCK45IzsY=`,
	)
	if err != nil {
		fmt.Println(err)
	}

	stdout, stderr, err := client.Run(`notathing`, ``)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout)
	fmt.Println(stderr)

	// Output:
	// bash: line 1: notathing: command not found
}
