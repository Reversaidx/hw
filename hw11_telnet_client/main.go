package main

import (
	"os"
	"time"
)

func main() {
	// Place your code here,
	//var timeout *time.Duration
	host := "127.0.0.1"
	port := "4242"
	timeout := time.Second * 10
	//*timeout = time.Second * 5
	//flag.DurationVar(timeout, "timeout", 0, "kurwa")
	NewTelnetClient(host+":"+port, timeout, os.Stdin, os.Stdout)
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
}
