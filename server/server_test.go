package server

import (
	"testing"
	"github.com/tatsuya4649/tuserver/test"
)

func TestServer(t *testing.T){
	Server(test.TestAddress,test.TestPort)
}
