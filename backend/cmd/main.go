package main

import (
	"fmt"
	"time"

	server "MelkOnline/internal"
)

func main() {
	fmt.Println("Starting the server...")
	time.Sleep(20 * time.Second)
	es := server.NewEchoServer()
	defer es.Stop()
	es.Start("8080")
}
