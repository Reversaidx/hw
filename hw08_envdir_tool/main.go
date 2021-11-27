package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	directory := os.Args[1]
	cmd := exec.Command("/bin/bash", os.Args[3:]...)
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		content, err := ioutil.ReadFile(directory + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}
		if len(content) == 0 {
			os.Unsetenv(file.Name())
		}
		firstLine := strings.Split(string(content), "\n")[0]
		contentT := strings.TrimRight(string(firstLine), " \t")

		contentT = string(bytes.Replace([]byte(contentT), []byte("\x00"), []byte("\n"), -1))
		err = os.Setenv(file.Name(), contentT)
		if err != nil {
			fmt.Println(contentT)
			log.Fatal(err)
		}
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
