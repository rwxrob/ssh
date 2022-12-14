package ssh_test

import (
	"fmt"
	"os"
	"time"

	"github.com/rwxrob/ssh"
)

/*
func ExampleRun() {
	ukey, _ := os.ReadFile(`testdata/blahpriv`)
	hkey, _ := os.ReadFile(`testdata/hostpubkey`)
	stdout, stderr, err := ssh.Run(`blah@localhost:22`, ukey, hkey, `cat`, `hello`)
	fmt.Printf("STDOUT\n%v\n", stdout)
	fmt.Printf("STDERR\n%v\n", stderr)
	fmt.Printf("ERROR\n%v\n", err)
	// Output:
	// ignored
}
*/

/*
func ExampleMultiHostClient_Run() {

	c := new(ssh.MultiHostClient)
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	ukey, _ := os.ReadFile(`testdata/blahpriv`)
	c.User, _ = ssh.NewUser(`blah`, ukey)
	hkey, _ := os.ReadFile(`testdata/hostpubkey`)
	host, _ := ssh.NewHost(`localhost:22`, hkey)
	dummy, _ := ssh.NewHost(`localhost:23434`, hkey) // will fail
	c.Hosts = []*ssh.Host{dummy, host, dummy, dummy}
	c.Timeout = 10 * time.Second
	c.Sleep = 5 * time.Second
	c.Attempts = 2

	stdout, stderr, err := c.Run(`cat`, `hello`)
	fmt.Printf("STDOUT\n%v\n", stdout)
	fmt.Printf("STDERR\n%v\n", stderr)
	fmt.Printf("ERROR\n%v\n", err)

	// Output:
	// ignored
}
*/

func ExampleMultiHostClient_run_assert_user() {
	c := new(ssh.MultiHostClient)
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	c.Run(`ls -l ~`, "")

	// Output:
	// undefined User
}

func ExampleMultiHostClient_run_assert_hosts() {
	c := new(ssh.MultiHostClient)
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	key, _ := os.ReadFile(`testdata/blahpriv`)
	c.User, _ = ssh.NewUser(`blah`, key)
	c.Run(`ls -l ~`, "")
	// Output:
	// undefined Hosts
}

func ExampleMultiHostClient_run_assert_timeout() {
	c := new(ssh.MultiHostClient)
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	ukey, _ := os.ReadFile(`testdata/blahpriv`)
	c.User, _ = ssh.NewUser(`blah`, ukey)
	hkey, _ := os.ReadFile(`testdata/hostpubkey`)
	host, _ := ssh.NewHost(`localhost`, hkey)
	c.Hosts = []*ssh.Host{host}

	c.Run(`ls -l ~`, "")
	// Output:
	// Timeout cannot be 0
}

func ExampleMultiHostClient_run_assert_attempts() {

	c := new(ssh.MultiHostClient)
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	ukey, _ := os.ReadFile(`testdata/blahpriv`)
	c.User, _ = ssh.NewUser(`blah`, ukey)
	hkey, _ := os.ReadFile(`testdata/hostpubkey`)
	host, _ := ssh.NewHost(`localhost`, hkey)
	c.Hosts = []*ssh.Host{host}
	c.Timeout = 10 * time.Second

	c.Run(`ls -l ~`, "")
	// Output:
	// Attempts cannot be 0
}
