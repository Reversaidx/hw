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


	directory:= os.Args[1]
	cmd:=exec.Command(os.Args[2],os.Args[2:]...)
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}


	for _, file := range files {

		if file.Size()!=0{
			content, err := ioutil.ReadFile(directory+"/"+file.Name())
			if err != nil {
				log.Fatal(err)
			}
			firstLine:=strings.Split(string(content), "\n")[0]
			contentT:=strings.TrimRight(string(firstLine),"\t")

			contentT:=string(bytes.Replace(byte(contentT),'\x00',byte("\n"),1000))
			err=os.Setenv(file.Name(),contentT)
			if err != nil {
				fmt.Println(contentT)
				log.Fatal(err)
			}
		}else {
			err:=os.Unsetenv(file.Name())
			if err != nil {
				fmt.Println("kurwa")
				log.Fatal(err)
			}
		}
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

