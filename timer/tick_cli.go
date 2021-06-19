package timer

import (
	"fmt"
)

var HeartAddress []string

func TimerCli(data ...interface{}){
	addresses := data[0]
	fmt.Println(addresses)
	for i:=0;i<len(data);i++{
		go func(){
			fmt.Println("Client ",i)
		}()
	}
}
