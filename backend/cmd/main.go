package internal

import (
	_ "MelkOnline/docs"
	"MelkOnline/internal"
	"fmt"
	"time"
)

//	@title			MelkOnline API
//	@version		1.0
//	@description	SE Project MelkOnline backend API

// @host		45.147.97.39:8080
// @BasePath	/
func main() {
	fmt.Println("Starting the server...")
	time.Sleep(20 * time.Second)
	es := internal.NewEchoServer()
	defer es.Stop()
	es.Start("8080")
}
