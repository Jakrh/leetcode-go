package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "test", "./...")

	out, err := cmd.Output()
	if err != nil {
		msg := "failed to run command:\n%s"
		if e, ok := err.(*exec.ExitError); ok {
			fmt.Fprintf(os.Stderr, msg, e.Stderr)
		} else {
			fmt.Fprintf(os.Stderr, msg+"\n", err)
		}
	}

	fmt.Println(string(out))
}
