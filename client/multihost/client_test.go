package multihost_test

import (
	"fmt"

	"github.com/rwxrob/ssh/client/multihost"
	C "github.com/rwxrob/ssh/internal/config"
)

func ExampleClient_Config() {
	c := new(multihost.Client)
	err := c.Set(C.Map{`timeout`: 2})
	fmt.Println(c.Timeout(), err)
	// Output:
	// 2 <nil>
}
