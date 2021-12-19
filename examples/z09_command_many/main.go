package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.New("192.168.18.101", "zhangdapeng", "zhangdapeng", 22)

	// 批量执行命令
	command1 := "mkdir test"
	command2 := "mkdir test/a"
	command3 := "ls -lah test"
	output, err := s.RunMany(command1, command2, command3)
	fmt.Printf("%v\n%v", output, err)
}
