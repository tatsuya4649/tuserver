package server

import (
	"testing"
	"github.com/tatsuya4649/tuserver/test"
)

func TestServer(t *testing.T){
	t.Run("Server Test", func (t *testing.T){
		t.Run("TCP server", func(t *testing.T){
			t.Parallel()
			Server(test.TestAddress,test.TestPort,"tcp")
		})
		t.Run("UDP server", func(t *testing.T){
			t.Parallel()
			Server(test.TestAddress,test.TestPort,"udp")
		})
		t.Run("Unix server",func(t *testing.T){
			t.Parallel()
			Server(test.TestUnixPath,test.TestPort,"unix")
		})
		t.Run("Unix Datagram server",func(t *testing.T){
			t.Parallel()
			Server(test.TestUnixDatagramPath,test.TestPort,"unixgram")
		})
	})
}
