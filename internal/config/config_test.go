package C_test

import (
	"fmt"

	C "github.com/rwxrob/ssh/internal/config"
)

func ExampleMap() {

	c := C.Map{
		`foo`:  `bar`,
		`one`:  1,
		`true`: true,
		`null`: nil,
		`struct`: struct {
			Str string
			Int int
		}{`str`, 42},
	}
	fmt.Println(c)
	c = nil
	fmt.Println(c)

	// Output:
	// {"foo":"bar","null":null,"one":1,"struct":{"Str":"str","Int":42},"true":true}
	// null

}
