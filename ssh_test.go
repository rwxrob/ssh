package ssh_test

import (
	"fmt"
	"log"
	"os"
	"testing"
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
func ExampleRunSafe() {
	ukey, _ := os.ReadFile(`testdata/blahpriv`)
	hkey, _ := os.ReadFile(`testdata/hostpubkey`)
	stdout, stderr, err := ssh.RunSafe(`blah@localhost:22`, ukey, hkey, `echo`, `hello world`)
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

	stdout, stderr, err := c.Run(`echo hello`, ``)
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

/*
# Benchmark comparative testing

## Scenarios

### One command per new connection

Same as using `ssh` from the command line with arguments passed to the target server shell.

### One command per session

Same as using `ssh` to login to an interactive session and restarting a shell on that connection for every command in a subshell.

### One command per cached connection

On first command initialize a connection and reuse that connection for subsequent sessions, restoring the connection if it drops for any reason.

### One command per connection pool

Initialize a connection pool per target server and randomize which connection gets each command restoring lost connections when then break or timeout.
*/

var ukey = `
-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
QyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQAAAJDswLYs7MC2
LAAAAAtzc2gtZWQyNTUxOQAAACB0/Xdc30JNJx+H1bvs9oZ7POIBi/YyZ4UBfQ/oEyffOQ
AAAEDWFaCmeeFjBMAzJvtf6z24ai1dHf2FSUmuHrONv/5K6XT9d1zfQk0nH4fVu+z2hns8
4gGL9jJnhQF9D+gTJ985AAAACXJ3eHJvYkB0dgECAwQ=
-----END OPENSSH PRIVATE KEY-----
`

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stdout, stderr, err := ssh.Run(`user@localhost:22`, []byte(ukey), nil, `echo hi there`, ``)
		if err != nil || stdout != "hi there\n" || stderr != `` {
			log.Print(stdout)
			return
		}
	}
}
