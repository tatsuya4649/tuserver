package server

import (
	"net"
	"os"
	"log"
)

func tcpHandler(conn *net.TCPConn,info *TUNetInfo,wch *chan int64){
	defer worker_end(wch)
	/* DO ANYTHING */


	conn.Close()
	return
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
