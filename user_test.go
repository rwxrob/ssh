package ssh_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rwxrob/ssh"
	"gopkg.in/yaml.v3"
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

	fmt.Println(user.YAML())

	// Output:
	// name: user
	// key: |-
	//     -----BEGIN OPENSSH PRIVATE KEY-----
	//     b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//     QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
	//     LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
	//     AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
	//     4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
	//     -----END OPENSSH PRIVATE KEY-----
}

func ExampleUser_from_YAML() {
	byt, _ := os.ReadFile(`testdata/user.yaml`)
	user := new(ssh.User)
	err := yaml.Unmarshal(byt, user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.YAML())
	// Output:
	// name: user
	// key: |-
	//     -----BEGIN OPENSSH PRIVATE KEY-----
	//     b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
	//     QyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8gAAAJAZeyGhGXsh
	//     oQAAAAtzc2gtZWQyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8g
	//     AAAEDdV9IJ3LNTiK7D0MFz7IR1Cz/VdqqH6SgOtiDz8/5073SN2HaGCU8lyGytWCcmNNv1
	//     tUciO6BLnCUmKmcgmI7yAAAACXJ3eHJvYkB0dgECAwQ=
	//     -----END OPENSSH PRIVATE KEY-----
}

func ExampleUser_from_JSON() {
	byt, _ := os.ReadFile(`testdata/user.json`)
	user := new(ssh.User)
	err := json.Unmarshal(byt, user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.JSON())
	// Output:
	// {"Name":"user","Key":"-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8gAAAJAZeyGhGXsh\noQAAAAtzc2gtZWQyNTUxOQAAACB0jdh2hglPJchsrVgnJjTb9bVHIjugS5wlJipnIJiO8g\nAAAEDdV9IJ3LNTiK7D0MFz7IR1Cz/VdqqH6SgOtiDz8/5073SN2HaGCU8lyGytWCcmNNv1\ntUciO6BLnCUmKmcgmI7yAAAACXJ3eHJvYkB0dgECAwQ=\n-----END OPENSSH PRIVATE KEY-----"}
}
