package main

import "fmt"

type server struct {
	name string
}

type server2 struct {
	name string
	server
}

func (s server) getServer() server {
	return server{name: "server111"}
}

func main() {
	// Create an instance of server2
	//srv2 := server2{
	//	name:   "server2",
	//	server: server{name: "embedded server"},
	//}

	// Call getServer method on the embedded server field
	ok := server2{}.getServer()
	fmt.Println(ok)
}
