package cli

import (
	"testing"
	"github.com/tatsuya4649/tuserver/test"
)

func TestCli(t *testing.T){
	Cli(test.TestAddress,test.TestPort,100)
}
