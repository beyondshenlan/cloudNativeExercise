package main

import (
	"bytes"
	"io"
	"log"
	"net"
	config "socket-practice/e1/delimiter"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handler(conn)
	}
}

func handler(c net.Conn) {
	defer c.Close()
	buf := make([]byte, config.BufferSize)
	result := bytes.NewBuffer(nil)
	for {
		n, err := c.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("ending: " + err.Error())
				return
			} else {
				log.Println("Read error: " + err.Error())
				break
			}
		}
		result.Write(buf[0:n])

		// pointer for read buffer
		var start int
		var end int
		for k, v := range result.Bytes() {
			// when byte equals to defined delimeter, then set to end pointer
			if v == config.Delimeter {
				end = k
				log.Printf("recevie: %v", string(result.Bytes()[start:end]))
				// move start pointer
				start = end + 1
			}
		}
		result.Reset()
	}
}
