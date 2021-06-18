package timer

import (
	"fmt"
	"time"
)

func Timer(duration int64){
	timer := time.NewTicker(time.Duration(duration)*time.Millisecond)
	for {
		<-timer.C
		fmt.Println("tick")
	}
}
