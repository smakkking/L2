package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	server, _ := net.Listen("tcp", "localhost:5050")

	for {
		conn, err := server.Accept()

		if err != nil {
			fmt.Println("error while connectiong, try again")
			conn.Close()
			continue
		}

		fmt.Println("successfully connected!")

		buffReader := bufio.NewReader(conn)

		go func(conn net.Conn) {
			defer conn.Close()

			for {
				rbyte, err := buffReader.ReadByte()

				if err != nil {
					fmt.Println("Can not read", err)
					break
				}

				fmt.Println(string(rbyte))
				conn.Write([]byte("recieved"))
			}
		}(conn)
	}
}
