package cli

import (
	"testing"
	"github.com/tatsuya4649/tuserver/test"
)

func TestCli(t *testing.T){
	t.Run("Client Test",func (t *testing.T){
		t.Run("TCP Client",func(t *testing.T){
			t.Parallel()
			Cli(test.TestAddress,test.TestPort,100,"tcp")
		})
		t.Run("UDP Client",func(t *testing.T){
			t.Parallel()
			Cli(test.TestAddress,test.TestPort,100,"udp")
		})
		t.Run("UNIX Client",func(t *testing.T){
			t.Parallel()
			Cli(test.TestPath,test.DummyPort,100,"unix")
		})
	})
}
