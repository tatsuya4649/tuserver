package server

func Server(port int32) error{
	udpServer("localhost","10000",1,1000)
	return nil
}

