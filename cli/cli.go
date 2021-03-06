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
	_,err = conn.Write([]byte{1,2,4})
}

func unixGen(path string,data interface{},network string){
	address,err := net.ResolveUnixAddr(network,path)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	conn,err := net.DialUnix(network,nil,address)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()
	/* connect to server */
	_,err = conn.Write([]byte{1,2,4})
}

func unixCli(path string,data interface{}){
	unixGen(path,data,"unix")
}

func unixDatagramCli(path string,data interface{}){
	unixGen(path,data,"unixgram")
}

func Cli(addr string,port int32,data interface{},network string){
	switch network{
	case "tcp":
		tcpCli(addr,port,data)
	case "udp":
		udpCli(addr,port,data)
	case "unix":
		unixCli(addr,data)
	case "unixgram":
		unixDatagramCli(addr,data)
	default:
		fmt.Printf("%s must be \"tcp\" or \"udp\"!",network)
		os.Exit(1)
	}
}
