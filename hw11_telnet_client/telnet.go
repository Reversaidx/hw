package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

//type TelnetClientS struct{}
//
//func (TelnetClientS) Connect() error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (TelnetClientS) Close() error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (TelnetClientS) Send() error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (TelnetClientS) Receive() error {
//	//TODO implement me
//	panic("implement me")
//}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	dialer := &net.Dialer{}
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	var client TelnetClient
	conn, err := dialer.Dial("tcp", address)
	stdin := stdinScan(cancel)
	if err != nil {
		log.Fatal(err)
	}
	os.Stderr.WriteString(fmt.Sprintf("Connected to %v\n", address))
	wg.Add(1)
	go func() {
		Receive(ctx, conn, wg)
		cancel()
	}()

	wg.Add(1)
	go func() {
		Send(ctx, conn, wg, stdin)
	}()
	wg.Wait()
	// Place your code here.
	return client
}

func stdinScan(cancel context.CancelFunc) chan string {
	out := make(chan string)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for !scanner.Scan() {
			out <- scanner.Text()
		}
		fmt.Println("Closing")
		//TODO add EOF check
		if scanner.Err() != nil {
			fmt.Println("EOF??")
			cancel()
			close(out)
		}
	}()
	return out
}

func Receive(ctx context.Context, conn net.Conn, wg *sync.WaitGroup) error {
	defer wg.Done()
	scanner := bufio.NewScanner(conn)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				break OUTER
			}
			text := scanner.Text()
			log.Printf(text)
		}
	}
	os.Stderr.WriteString("Closing read from connection")
	return nil
}

func Send(ctx context.Context, conn net.Conn, wg *sync.WaitGroup, stdin chan string) error {
	defer wg.Done()
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		case str := <-stdin:
			//if !scanner.Scan() {
			//	break OUTER
			//}
			//str := scanner.Text()
			//log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}

	os.Stderr.WriteString("Closing write to connection")
	return nil
}
