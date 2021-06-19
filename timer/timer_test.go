package timer

import (
	"testing"
	"fmt"
)


func TestTimer(t *testing.T){
	t.Run("Timer",func(t *testing.T){
		HeartAddress = append(HeartAddress,"8.8.8.8","43.34.3.3453")
		fmt.Println(HeartAddress)
		Timer(1000,-1,TimerCli,HeartAddress)
	})
}
