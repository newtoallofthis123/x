package utils

import (
	"os"
	"os/exec"
	"strings"
)

// Run executes the provided commands on the specified file.
// It returns an error if any issues are encountered during execution.
func Run(file string, cmds []string) error {
	output := os.Stdout

	if file != "" {
		if file == "!" {
			file = os.DevNull
		}

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

		pipes := strings.Split(c, "|")
		if len(pipes) > 1 {
			execs := make([]*exec.Cmd, 0)

			for i := 0; i < len(pipes); i++ {
				pipes[i] = strings.TrimSpace(pipes[i])
				cs := strings.Split(pipes[i], " ")
				cmd := exec.Command(cs[0], cs[1:]...)
				execs = append(execs, cmd)
			}

			for i := 0; i < len(execs)-1; i++ {
				execs[i].Stdout, _ = execs[i+1].StdinPipe()
			}

			execs[len(execs)-1].Stdout = output

			for _, e := range execs {
				err := e.Start()
				if err != nil {
					return err
				}
			}
			continue
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

// RunTasks executes the given tasks and writes the output to the specified file.
// It returns an error if any issues are encountered during execution.
func RunTasks(output string, tasks map[string][]string) error {
	for _, cmds := range tasks {
		err := Run(output, cmds)
		if err != nil {
			return err
		}
	}

	return nil
}
