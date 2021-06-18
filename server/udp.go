package server

import (
	"fmt"
	"net"
	"os"
	"log"
	"strconv"
)

func udpHandler(conn *net.UDPConn,info *TUNetInfo,wch *chan int64){
	defer worker_end(wch)
	/* DO ANYTHING */
	fmt.Println(*info)

	conn.Close()
	return
}

func udpServer(address string,port int32,worker,bufsize int64){
	wch := make(chan int64,worker)
	addport := net.JoinHostPort(address,fmt.Sprint(port))
	addr,err := net.ResolveUDPAddr("udp",addport)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	conn,err := net.ListenUDP("udp",addr)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	info,err := TUNetParse(conn)
	for {
		wch<-0
		buf := make([]byte,bufsize)
		_,remote,err := conn.ReadFromUDP(buf)
		if err != nil{
			continue
		}
		remote_address,remote_port,err := net.SplitHostPort(remote.String())
		if err!=nil{
			continue
		}
		remop,err := strconv.ParseInt(remote_port,10,64)
		if err!=nil{
			continue
		}
		info.RemoteHost = remote_address
		info.RemotePort = remop
		go udpHandler(conn,&info,&wch)
	}
}
