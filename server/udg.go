// Unix Datagram Server
package server

import (
	"net"
	"os"
)

func unixDatagramHandler(conn *net.UnixConn,wch *chan int64){
	defer worker_end(wch)
	/* DO ANYTHING */
	fmt.Println("Success,Unix Datagram")

	conn.Close()
	return
}

func unixDatagramServer(path string,worker int64){
	wch := make(chan int64,worker)
}
