package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/nyaosorg/go-readline-ny"
	"github.com/nyaosorg/go-readline-ny/coloring"
	"github.com/nyaosorg/go-readline-ny/simplehistory"
)

func prompt() (int, error) {
	user, err := user.Current()
	if err != nil {
		return 1, err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return 1, err
	}

	dir, err := os.Getwd()
	if err != nil {
		return 1, err
	}

	promptStr := fmt.Sprintf("%v@%v:%v\n> ", user.Name, hostname, dir)
	return fmt.Print(promptStr)
}

func loop() error {
	history := simplehistory.New()

	editor := readline.Editor{
		Prompt:         prompt,
		Writer:         colorable.NewColorableStdout(),
		History:        history,
		Coloring:       &coloring.VimBatch{},
		HistoryCycling: true,
	}

	fmt.Println("vent. Type Ctrl-D to quit.")
	for {
		text, err := editor.ReadLine(context.Background())
		if err != nil {
			return err
		}

		fields := strings.Fields(text)
		if len(fields) <= 0 {
			continue
		}
		ext := filepath.Ext(fields[0])

		cmd := &exec.Cmd{}
		if ext == ".go" {
			goRunFields := []string{"run", fields[0]}
			goRunFields = append(goRunFields, fields[1:]...)
			cmd = exec.Command("go", goRunFields...)
		} else {
			cmd = exec.Command(fields[0], fields[1:]...)
		}

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		cmd.Run()

		history.Add(text)
	}
}

func main() {
	if err := loop(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
