package main

import (
	"fmt"
	"log"

	handler "github.com/go-malproxy/server/handler/detection"
)

// func main() {
// 	server.Run() //起動
// }

func main() {
	hostname, err := handler.GetHostnameFromIPAddress("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(hostname)
	}
}
