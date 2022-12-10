package remote_test

import (
	"fmt"
	"os"

	"github.com/rwxrob/remote"
)

func ExampleSSH() {
	ukey, _ := os.ReadFile(`testdata/blahpriv`)
	hkey, _ := os.ReadFile(`testdata/hostpubkey`)
	stdout, stderr, err := remote.SSH(`blah@localhost:22`, ukey, hkey, `cat`, `hello`)
	fmt.Printf("STDOUT\n%v\n", stdout)
	fmt.Printf("STDERR\n%v\n", stderr)
	fmt.Printf("ERROR\n%v\n", err)
	// Output:
	// ignored
}
