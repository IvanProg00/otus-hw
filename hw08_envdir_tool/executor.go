package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for key, e := range env {
		if e.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				panic(err)
			}
		} else {
			err := os.Setenv(key, e.Value)
			if err != nil {
				panic(err)
			}
		}
	}

	command := cmd[0]
	cmdRun := exec.Command(command, cmd[1:]...)
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr

	if err := cmdRun.Run(); err != nil {
		fmt.Printf("%+v\n", err)
		exitError := &exec.ExitError{}
		ok := errors.As(err, &exitError)
		if !ok {
			panic(err)
		}
		return exitError.ExitCode()
	}

	return
}
