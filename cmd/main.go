package main

import (
	"log"
	"net"
)

func main() {
	lis, err := net.Lisen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
