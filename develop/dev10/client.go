package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// https://www.developer.com/languages/intro-socket-programming-go/

func StdinReader(r io.Reader) chan string {
	result := make(chan string)

	go func() {
		defer close(result)
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			result <- scanner.Text()
		}
	}()
	return result
}

func main() {
	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "")

	d := net.Dialer{Timeout: timeout}

	conn, _ := d.Dial("tcp", "localhost:5050")
	defer conn.Close()

	read_stdin := StdinReader(os.Stdin)
	read_socket := StdinReader(conn)

	for {
		select {
		case x, ok := <-read_socket:
			if !ok {
				break
			}
			fmt.Println(x)
		case x := <-read_stdin:
			conn.Write([]byte(x)) // stdin записывается в сокет
		}
	}
}
