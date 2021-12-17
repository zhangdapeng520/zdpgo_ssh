package zdpgo_ssh

import (
	"fmt"
	"os"
	"testing"
)

func MTestRun(t *testing.T) {
	cli := New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	output, err := cli.Run("free -h")
	fmt.Printf("%v\n%v", output, err)
}

func MTestRunTerminal(t *testing.T) {
	cli := New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	cli.RunTerminal("top", os.Stdout, os.Stdin)
}
