package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmdS []string, env Environment) (returnCode int) {
	cmd := exec.Command(cmdS[0], cmdS[1:]...) //nolint
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	envSlice := make([]string, len(env))
	i := 0
	for k, v := range env {
		envSlice[i] = k + "=" + v.Value
		i++
	}
	cmd.Env = envSlice
	// fmt.Println(cmd.String())
	if err := cmd.Run(); err != nil {
		fmt.Println(err)

		if exitError, ok := err.(*exec.ExitError); ok { //nolint
			return exitError.ExitCode()
		}
	}

	return 0
}
