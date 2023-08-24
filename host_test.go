package ssh_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/ssh"
)

func ExampleNewHost() {

	host, err := ssh.NewHost(`localhost`, `randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(host.Options()[0])
	fmt.Println(host.YAML())

	// Output:
	// randomoption
	// addr: localhost
	// auth: randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=

}

func ExampleNewHostFromYAML() {
	file, _ := os.Open(`testdata/host.yaml`)
	host, _ := ssh.NewHostFromYAML(file)
	fmt.Println(host.Options()[0])
	fmt.Println(host.YAML())
	// Output:
	// randomoption
	// addr: localhost
	// auth: |
	//     randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=
}

func ExampleNewHostFromJSON() {
	file, _ := os.Open(`testdata/host.json`)
	host, _ := ssh.NewHostFromJSON(file)
	fmt.Println(host.Options()[0])
	fmt.Println(host.YAML())
	// Output:
	// randomoption
	// addr: localhost
	// auth: |
	//     randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=

}
