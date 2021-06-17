package cli

import (
	"net"
	"log"
	"os"
	"fmt"
)

func tcpCli(addr string,port int32,data interface{}){
	addrport := net.JoinHostPort(addr,fmt.Sprint(port))
	address,err := net.ResolveTCPAddr("tcp",addrport)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	conn,err := net.DialTCP("tcp",nil,address)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()
	/* connect to server */
	_,err = conn.Write([]byte{1,2,3})
}

func udpCli(addr string,port int32,data interface{}){
	addrport := net.JoinHostPort(addr,fmt.Sprint(port))
	address,err := net.ResolveUDPAddr("udp",addrport)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	conn,err := net.DialUDP("udp",nil,address)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()
	/* connect to server */
}

func Cli(addr string,port int32,data interface{}){
	tcpCli(addr,port,data)
}
