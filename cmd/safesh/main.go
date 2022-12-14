package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) < 3 || os.Args[1] != `-c` {
		log.Println(`not an interactive shell (-c required)`)
		os.Exit(1)
	}

	f := strings.Fields(os.Args[2])

	cmd := exec.Command(f[0], f[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}

}
