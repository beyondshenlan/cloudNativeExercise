package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main()  {
	after:=time.After(5*time.Second)
	iter:=0

	conn,err:=net.Dial("tcp","localhost:8080")
	if err!=nil{
		log.Fatal(err.Error())

	}
	defer conn.Close()
	for{
		select {
		case <-after:
			log.Println("")
			return
		default:
			for i:=0;i<10;i++{
				content:="hello["+strconv.Itoa(iter)+"]"
				_,err:=conn.Write([]byte(content))
				if err!=nil{
					log.Fatal(err.Error())
				}
				iter++
			}
			time.Sleep(1*time.Second)
		}
	}

}



