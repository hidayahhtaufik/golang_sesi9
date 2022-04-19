package main

import "sesi9/routers"

func main() {
	var port = ":8881"

	server := routers.StartServer()
	server.Run(port)
}
