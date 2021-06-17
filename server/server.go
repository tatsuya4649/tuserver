package server

import (
	"fmt"
	"net"
	"os"
	"log"
	"strconv"
)

type TUNetInfo struct{
	Id	   int64
	RemoteHost string
	RemotePort int64
	LocalHost  string
	LocalPort  int64
}

/* At the end of connection, release worker with channel */
func worker_end(wch *chan int64){
	<-(*wch)
}

func tcpHandler(conn *net.TCPConn,info *TUNetInfo,wch *chan int64){
	defer worker_end(wch)
	/* DO ANYTHING */


	conn.Close()
	return
}

/* TCP/UDP interface for get Address */
type TU interface{
	RemoteAddr() net.Addr
	LocalAddr()  net.Addr
}

//func TUNetParse(conn TU) (id,remotePort,localPort int64,remoteHost,localHost string,err error){
func TUNetParse(conn TU) (info TUNetInfo,err error){
	remote := conn.RemoteAddr()
	var remote_address string
	var remote_port string
	var remop int64
	if remote!=nil{
		remote_address,remote_port,err = net.SplitHostPort(remote.String())
		if err!=nil{
			return TUNetInfo{},err
		}
		remop,err = strconv.ParseInt(remote_port,10,64)
		if err!=nil{
			return TUNetInfo{},err
		}
	}
	local := conn.LocalAddr()
	local_address,local_port,err := net.SplitHostPort(local.String())
	if err!=nil{
		return TUNetInfo{},err
	}
	locap,err := strconv.ParseInt(local_port,10,64)
	if err!=nil{
		return TUNetInfo{},err
	}
	info = TUNetInfo{
		RemoteHost : remote_address,
		RemotePort : remop,
		LocalHost : local_address,
		LocalPort : locap,
	}
	return info,nil
}

/*  tcpServer received server address,port number,
 *  worker count (work threads)
 */
func tcpServer(address string,port string,worker int64){
	wch := make(chan int64,worker)
	addport := net.JoinHostPort(address,port)
	addr,err := net.ResolveTCPAddr("tcp",addport)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	listen,err := net.ListenTCP("tcp",addr)
	defer listen.Close()
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	for{
		conn,err := listen.AcceptTCP()
		wch<-0
		if err!=nil{
			continue
		}
		info,err := TUNetParse(conn)
		if err != nil{
			conn.Close()
			continue
		}
		go tcpHandler(conn,&info,&wch)
	}
}

func udpHandler(conn *net.UDPConn,info *TUNetInfo,wch *chan int64){
	defer worker_end(wch)
	/* DO ANYTHING */
	fmt.Println(*info)

	conn.Close()
	return
}

func udpServer(address,port string,worker,bufsize int64){
	wch := make(chan int64,worker)
	addport := net.JoinHostPort(address,port)
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

func Server(port int32) error{
	udpServer("localhost","10000",1,1000)
	return nil
}

