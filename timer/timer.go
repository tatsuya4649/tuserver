package timer

import (
	"fmt"
	"time"
)

func TestTock(data ...interface{}){
	fmt.Println("Test Tock")
}

/*
 * duration => millisecond
 * count => how many times "tick"
 * tock,tock_arg => any function at tick
 */
func Timer(duration int64,count int64,tock func(data ...interface{}),tock_arg ...interface{}){
	orig_count := count
	timer := time.NewTicker(time.Duration(duration)*time.Millisecond)
	for {
		<-timer.C
		tock(tock_arg)
		count--
		if count==0{
			break
		}
	}
	fmt.Printf("end count %d\n",orig_count)
}
