package main

import (
	"log"
	"os"
)

func main() {
	directory := os.Args[1]
	env, err := ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	code := RunCmd(os.Args[3:], env)
	if code != 0 {
		os.Exit(code)
	}
}
