package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	conn,err := net.Dial("tcp","127.0.0.1:8999")
	if err!= nil{

		fmt.Println(err)
	}
	for {
		_, err = conn.Write([]byte{'a', 'b'})
		if err != nil {
			fmt.Println(err)
		}
		buf := make([]byte, 512)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("return content is %s \n", buf)
		time.Sleep(1*time.Second)
	}
}
