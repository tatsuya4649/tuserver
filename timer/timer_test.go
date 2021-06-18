package timer

import (
	"testing"
)


func TestTimer(t *testing.T){
	t.Run("Timer",func(t *testing.T){
		Timer(1000)
	})
}
