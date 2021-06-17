package server

import (
	"net"
	"strconv"
)

type TUNetInfo struct{
	Id	   int64
	RemoteHost string
	RemotePort int64
	LocalHost  string
	LocalPort  int64
}

/* TCP/UDP interface for get Address */
type TU interface{
	RemoteAddr() net.Addr
	LocalAddr()  net.Addr
}

/* At the end of connection, release worker with channel */
func worker_end(wch *chan int64){
	<-(*wch)
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
