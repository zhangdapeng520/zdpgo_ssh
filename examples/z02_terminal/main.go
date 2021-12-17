package main

import (
	"os"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	s.RunTerminal("top", os.Stdout, os.Stdin)
}
