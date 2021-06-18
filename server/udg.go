// Unix Datagram Server
package server

import (
	"net"
	"os"
	"log"
	"fmt"
)

func unixDatagramHandler(conn *net.UnixConn,buffer []byte,wch *chan int64){
	defer worker_end(wch)
	/* DO ANYTHING */
	fmt.Println("Success,Unix Datagram")

	conn.Close()
	return
}

func unixDatagramServer(path string,worker int64,bufsize int64){
	wch := make(chan int64,worker)
	addr,err := net.ResolveUnixAddr("unixgram",path)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	conn,err := net.ListenUnixgram("unixgram",addr)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		wch<-0
		buf := make([]byte,bufsize)
		_,_,err :=  conn.ReadFromUnix(buf)
		if err!=nil{
			continue
		}
		go unixDatagramHandler(conn,buf,&wch)
	}
}
