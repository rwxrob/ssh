package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

var Allow = map[string]int8{`pwd`: 1, `ls`: 1, `id`: 1, `echo`: 1}

func main() {
	log.SetFlags(0)

	if len(os.Args) < 3 || os.Args[1] != `-c` {
		log.Println(`not an interactive shell (-c required)`)
		os.Exit(1)
	}

	f := strings.Fields(os.Args[2])

	if _, allowed := Allow[f[0]]; !allowed {
		log.Printf("command not allowed: %v\n", f[0])
		os.Exit(1)
	}

	cmd := exec.Command(f[0], f[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}

}
