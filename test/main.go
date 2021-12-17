package main

import (
	"os"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	cli := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	cli.RunTerminal("ls -lah", os.Stdout, os.Stdin)
}
