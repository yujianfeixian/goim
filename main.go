package main

import "goim/server"

func main() {
	server.NewServer("0.0.0.0", 8080).Start()
}
