package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.New("192.168.18.101", "zhangdapeng", "zhangdapeng", 22)

	// debug模式sudo命令
	output, err := s.SudoDebug("ls -lah", true)
	fmt.Printf("%v\n%v", output, err)

	
	// debug模式run命令
	output, err = s.RunDebug("ls -lah", true)
	fmt.Printf("%v\n%v", output, err)
}
