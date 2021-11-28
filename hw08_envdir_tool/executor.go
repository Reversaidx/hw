package main

import (
	"errors"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmdS []string, env Environment) (returnCode int) {
	var ExitError exec.ExitError
	cmd := exec.Command("/bin/bash", cmdS...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	envSlice := make([]string, len(env))
	i := 0
	for k, v := range env {
		envSlice[i] = k + "=" + v.Value
		i++
	}
	cmd.Env = envSlice
	if err := cmd.Run(); err != nil {
		if errors.As(err, ExitError) { //nolint
			return err.(*exec.ExitError).ExitCode() //nolint
		}
	}
	return 0
}
