package main

import (
	_ "MelkOnline/docs"
	server "MelkOnline/internal"
	"fmt"
)

//	@title			MelkOnline API
//	@version		1.0
//	@description	SE Project MelkOnline backend API

// @host		localhost:8080
// @BasePath	/
func main() {
	fmt.Println("Starting the server...")
	//	time.Sleep(20 * time.Second)
	es := server.NewEchoServer()
	defer es.Stop()
	es.Start("8080")
}
