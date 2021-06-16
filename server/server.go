package server

import (
	"fmt"
	"net"
	"io"
)

func tcpServer(){
}

func Server(port int32) error{
	var portString string = fmt.Sprintf(":%d",port)
	ln,err := net.Listen("tcp",portString)
	defer ln.Close()
	if err != nil{
		return err
	}
	for {
		con, err := ln.Accept()
		fmt.Println(con)
		if err != nil{
			return err
		}
		go handler(con,1000000)
	}
	return nil
}

func handler(con net.Conn,buflen int64){
	var b []byte = make([]byte,buflen)
	for {
		n,err:=con.Read(b)
		if err!=nil{
			/* close */
			if err == io.EOF{
				return
			}
			return
		}
		n,err=con.Write(b[:n])
		if err!=nil{
			fmt.Println(err)
			return
		}
	}
}

