package utils

import (
	"os"
	"os/exec"
	"strings"
)

func Run(file string, cmds []string) error {
	output := os.Stdout

	if file != "" {
		// open the file
		f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		output = f
	}

	for _, c := range cmds {
		if c[0] == '"' && c[len(c)-1] == '"' {
			c = c[1 : len(c)-1]
		}

		cs := strings.Split(c, " ")
		cmd := exec.Command(cs[0], cs[1:]...)

		// print to stdout
		cmd.Stdout = output
		cmd.Stderr = output

		err := cmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func RunTasks(output string, tasks map[string][]string) error {
	for _, cmds := range tasks {
		err := Run(output, cmds)
		if err != nil {
			return err
		}
	}

	return nil
}
