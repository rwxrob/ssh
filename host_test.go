package ssh_test

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/rwxrob/ssh"
)

func ExampleHost_with_Authorized_Hosts_Key() {

	host, err := ssh.NewHost(`localhost`, `randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(host.Options()[0])

	// Output:
	// randomoption
}

func ExampleHost_from_YAML() {
	host := new(ssh.Host)
	byt, _ := os.ReadFile(`testdata/host.yaml`)
	err := yaml.Unmarshal(byt, host)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(host.YAML())
	// Output:
	// addr: localhost
	// auth: |
	//     randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=

}
