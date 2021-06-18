package server

import (
	"net"
	"os"
	"fmt"
	"log"
)

func unixHandler(conn *net.UnixConn,wch *chan int64){
	defer worker_end(wch)
	/* DO ANYTHING */
	fmt.Println("Success,Unix")
	fmt.Println(conn.RemoteAddr())

	conn.Close()
	return
}

func unixServer(path string,worker int64){
	wch := make(chan int64,worker)
	address,err := net.ResolveUnixAddr("unix",path)
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	listen,err := net.ListenUnix("unix",address)
	defer listen.Close()
	if err!=nil{
		log.Fatal(err)
		os.Exit(1)
	}
	for{
		conn,err := listen.AcceptUnix()
		wch<-0
		if err!=nil{
			continue
		}
		go unixHandler(conn,&wch)
	}

}
