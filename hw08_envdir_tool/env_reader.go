package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var errEnv = errors.New("something went wrong")

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(directory string) (Environment, error) {
	envMap, err := parseEnv(os.Environ())
	if err != nil {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(directory + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		if len(content) == 0 {
			envMap[file.Name()] = EnvValue{
				Value:      envMap[file.Name()].Value,
				NeedRemove: true,
			}
		}
		firstLine := strings.Split(string(content), "\n")[0]
		contentT := strings.TrimRight(firstLine, " \t")

		contentT = string(bytes.ReplaceAll([]byte(contentT), []byte("\x00"), []byte("\n")))
		envMap[file.Name()] = EnvValue{
			Value:      contentT,
			NeedRemove: false,
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	// Place your code here
	return envMap, nil
}

func parseEnv(env []string) (Environment, error) {
	result := make(Environment)
	for _, v := range env {
		s := strings.Split(v, "=")
		if len(s) != 2 {
			return nil, errEnv
		}
		result[s[0]] = EnvValue{
			Value:      s[1],
			NeedRemove: false,
		}
	}
	return result, nil
}
