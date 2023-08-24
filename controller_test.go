package ssh_test

import (
	"fmt"
	"strings"

	"github.com/rwxrob/ssh"
)

func ExampleController_YAML() {

	c1 := ssh.NewClient()
	c1.User, _ = ssh.NewUser(
		`user`, `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
-----END OPENSSH PRIVATE KEY-----
`)
	c1.Host, _ = ssh.NewHost(
		`localhost`, `localhost ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=`,
	)

	ctl := ssh.NewController(c1)
	fmt.Println(ctl.YAML())

	// Output:
	// clients:
	//     - host:
	//         addr: localhost
	//         auth: localhost ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=
	//       port: 22
	//       user:
	//         name: user
	//         key: |-
	//             -----BEGIN OPENSSH PRIVATE KEY-----
	//             b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//             QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
	//             LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
	//             AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
	//             4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
	//             -----END OPENSSH PRIVATE KEY-----
	//       timeout: 5m0s

}

func ExampleController_JSON() {

	c1 := ssh.NewClient()
	c1.User, _ = ssh.NewUser(
		`user`, `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
-----END OPENSSH PRIVATE KEY-----
`)
	c1.Host, _ = ssh.NewHost(
		`localhost`, `localhost ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=`,
	)

	ctl := ssh.NewController(c1)
	fmt.Println(ctl.JSON())

	// Output:
	// {"Clients":[{"Host":{"Addr":"localhost","Auth":"localhost ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8="},"Port":22,"User":{"Name":"user","Key":"-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2\nLAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ\nAAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8\n4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=\n-----END OPENSSH PRIVATE KEY-----"},"Timeout":300000000000}]}

}

func ExampleController_RunOnAny() {

	c1 := ssh.NewClient()
	c1.User, _ = ssh.NewUser(
		`user`, `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
-----END OPENSSH PRIVATE KEY-----
`)
	c1.Host, _ = ssh.NewHost(
		`localhost`, ``,
	)

	ctl := ssh.NewController(c1).Connect()
	fmt.Println(c1.Connected())
	stdout, stderr, err := ctl.RunOnAny(`echo hello`, ``)
	fmt.Printf("stdout: %q stderr: %q err: %q\n", stdout, stderr, err)

	// Output:
	// true
	// stdout: "hello\n" stderr: "" err: %!q(<nil>)

}

func ExampleController_RunOnAny_no_Clients() {

	ctl := ssh.NewController()
	_, _, err := ctl.RunOnAny(`echo hello`, ``)
	fmt.Println(err)

	// Output:
	// all SSH client targets are unavailable

}

func ExampleController_RandomClient_none() {

	// idle, unconnected
	c1 := ssh.NewClient()
	c2 := ssh.NewClient()
	c3 := ssh.NewClient()

	ctl := ssh.NewController(c1, c2, c3)
	client := ctl.RandomClient()
	fmt.Println(client)

	// Output:
	// <nil>
}

func ExampleController_RandomClient_single_Good() {

	// good one
	c1 := ssh.NewClient()
	c1.User, _ = ssh.NewUser(
		`user`, `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
-----END OPENSSH PRIVATE KEY-----
`)
	c1.Host, _ = ssh.NewHost(
		`localhost`, ``,
	)
	c1.Connect()

	// idle
	c2 := ssh.NewClient()
	c3 := ssh.NewClient()

	ctl := ssh.NewController(c1, c2, c3)
	client := ctl.RandomClient()
	fmt.Println(client.User.Name)

	// Output:
	// user
}

func ExampleNewControllerFromYAML() {

	// note that YAML references are allowed and expanded

	source := `
users:
  user: &user
    name: user
    key: |
      -----BEGIN OPENSSH PRIVATE KEY-----
      b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
      QyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8gAAAJAZeyGhGXsh
      oQAAAAtzc2gtZWQyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8g
      AAAEDdV9IJ3LNTiK7D0MFz7IR1Cz/VdqqH6SgOtiDz8/5073SN2HaGCU8lyGytWCcmNNv1
      tUciO6BLnCUmKmcgmI7yAAAACXJ3eHJvYkB0dgECAwQ=
      -----END OPENSSH PRIVATE KEY-----

hosts:
  localhost: &localhost
    addr: localhost
    auth: |
      randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=

clients:
  - host: *localhost
    user: *user
    port: 2221
    timeout: 5m
  - host: *localhost
    user: *user
    port: 2222
    timeout: 5m
  - host: *localhost
    user: *user
    port: 2223
    timeout: 5m
`

	ctl, _ := ssh.NewControllerFromYAML(strings.NewReader(source))
	fmt.Println(ctl.YAML())

	// Output:
	// clients:
	//     - host:
	//         addr: localhost
	//         auth: |
	//             randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=
	//       port: 2221
	//       user:
	//         name: user
	//         key: |
	//             -----BEGIN OPENSSH PRIVATE KEY-----
	//             b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//             QyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8gAAAJAZeyGhGXsh
	//             oQAAAAtzc2gtZWQyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8g
	//             AAAEDdV9IJ3LNTiK7D0MFz7IR1Cz/VdqqH6SgOtiDz8/5073SN2HaGCU8lyGytWCcmNNv1
	//             tUciO6BLnCUmKmcgmI7yAAAACXJ3eHJvYkB0dgECAwQ=
	//             -----END OPENSSH PRIVATE KEY-----
	//       timeout: 5m0s
	//     - host:
	//         addr: localhost
	//         auth: |
	//             randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=
	//       port: 2222
	//       user:
	//         name: user
	//         key: |
	//             -----BEGIN OPENSSH PRIVATE KEY-----
	//             b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//             QyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8gAAAJAZeyGhGXsh
	//             oQAAAAtzc2gtZWQyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8g
	//             AAAEDdV9IJ3LNTiK7D0MFz7IR1Cz/VdqqH6SgOtiDz8/5073SN2HaGCU8lyGytWCcmNNv1
	//             tUciO6BLnCUmKmcgmI7yAAAACXJ3eHJvYkB0dgECAwQ=
	//             -----END OPENSSH PRIVATE KEY-----
	//       timeout: 5m0s
	//     - host:
	//         addr: localhost
	//         auth: |
	//             randomoption ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBI/WBBaaNFajVHCL0+rQqWP3zhpyXo357iPUvl0GGHWrY6t42WTNJ+bk8shRq7eq8KwefZeL4YvsnekcZb8Uq+8=
	//       port: 2223
	//       user:
	//         name: user
	//         key: |
	//             -----BEGIN OPENSSH PRIVATE KEY-----
	//             b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//             QyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8gAAAJAZeyGhGXsh
	//             oQAAAAtzc2gtZWQyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8g
	//             AAAEDdV9IJ3LNTiK7D0MFz7IR1Cz/VdqqH6SgOtiDz8/5073SN2HaGCU8lyGytWCcmNNv1
	//             tUciO6BLnCUmKmcgmI7yAAAACXJ3eHJvYkB0dgECAwQ=
	//             -----END OPENSSH PRIVATE KEY-----
	//       timeout: 5m0s

}
