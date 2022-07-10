package main

import (
	"log"
	"net"
	config "socket-practice/e1/fixLength"
	"strconv"
	"time"
)

func main() {
	after := time.After(5 * time.Second)
	iter := 0
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	// sending message
	for {
		select {
		case <-after:
			log.Println("time out")
			return
		default:
			for i := 0; i < 10; i++ {
				content := "hello[" + strconv.Itoa(iter) + "]"
				_, err = conn.Write(patch(content))
				if err != nil {
					log.Fatal(err.Error())
				}
				iter++
			}
			time.Sleep(1 * time.Second)
		}
	}
}

// patching empty byte into origin message
func patch(message string) []byte {
	res := make([]byte, config.BufferSize)
	copy(res, []byte(message))
	return res
}
