package ssh_test

import (
	"fmt"

	"github.com/rwxrob/ssh"
)

func ExampleClientMap_Random() {
	m := ssh.NewClientMap()

	c1 := ssh.NewClient()
	c1.Port = 1
	m.Add(c1)

	c2 := ssh.NewClient()
	c2.Port = 2
	m.Add(c2)

	c3 := ssh.NewClient()
	c3.Port = 3
	m.Add(c3)

	r := m.Random()
	fmt.Println(r) // nil, none Connected yet

	c1.Connected = true
	c2.Connected = true
	c3.Connected = true

	r = m.Random()
	fmt.Println(r)

	/// Output:
	// <nil>
	// <nil>@<nil>:2
}
