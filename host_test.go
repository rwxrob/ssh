package ssh_test

import (
	"fmt"

	"github.com/rwxrob/ssh"
)

func ExampleHost_String() {

	host, err := ssh.NewHost(`host`, ``)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(host)

	// Output:
	// host
}

func ExampleHost_with_Authorized_Hosts_Key() {

	host, err := ssh.NewHost(`localhost`, `randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(host.Options[0])

	// Output:
	// randomoption
}
