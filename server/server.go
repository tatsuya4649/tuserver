package server

import (
	"fmt"
	"os"
)

func Server(addr string,port int32,network string){
	switch network{
	case "tcp":
		tcpServer(addr,port,1)
	case "udp":
		udpServer(addr,port,1,1000)
	case "unix":
		unixServer(addr,1)
	default:
		fmt.Printf("%s must be \"tcp\" or \"udp\"!",network)
		os.Exit(1)
	}
}

