package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mattn/go-tty"
	"github.com/mattn/go-tty/ttyutil"
)

func loop() error {
	fmt.Println("Tiny shell. Type Ctrl-D to quit.")

	tty1, err := tty.Open()
	if err != nil {
		return err
	}
	defer tty1.Close()

	for {
		clean, err := tty1.Raw()
		if err != nil {
			return err
		}
		fmt.Println("Ready")
		text, err := ttyutil.ReadLine(tty1)
		clean()
		if err != nil {
			return err
		}

		fields := strings.Fields(text)
		if len(fields) <= 0 {
			continue
		}

		cmd := exec.Command(fields[0], fields[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Run()
	}
}

func main() {
	if err := loop(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
