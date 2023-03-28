package main

import "github.com/Gabriel-Newton-dev/web-api-with-golang/server"

func main() {
	server := server.NewServer()

	server.Run()
}
