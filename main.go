package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()
	listen, err := net.Listen("tcp", ":8888")

	if err != nil {
		log.Fatalf("error == ", err.Error())
	}
	defer listen.Close()
	log.Printf("start server on :8888")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("unableto accept connection: %s", err.Error())
			continue
		}
		go s.newClient(conn)
	}
}
