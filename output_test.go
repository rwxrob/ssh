package remote_test

import (
	"fmt"

	"github.com/rwxrob/remote"
)

func ExampleTest() {
	uid32 := remote.Base32()
	fmt.Println(len(uid32))
	// Output:
	// 32
}

/*
func ExampleOutput_String() {
	out := remote.Output{
		`stdout`:  "some standard output\n\non multiple lines",
		`stderr`:  "some standard err on single line",
		`exitval`: "-1",
	}
	fmt.Print(out)
	// Output:
	// ignored
}

func ExampleOutput_UnmarshalText() {
	out := remote.Output{`dummy`: `just checking`}
	buf := `NNQLO9MP27BRECLC6CED8QC2RGHQPHRL stdout
some standard output

on multiple lines
NNQLO9MP27BRECLC6CED8QC2RGHQPHRL stderr
some standard err on single line
NNQLO9MP27BRECLC6CED8QC2RGHQPHRL exitval
-1
NNQLO9MP27BRECLC6CED8QC2RGHQPHRL end
`
	err := out.UnmarshalText([]byte(buf))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(out)
	// Output:
	// ignored
}
*/
