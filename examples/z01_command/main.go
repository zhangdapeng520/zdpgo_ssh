package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	output, err := s.Run("free -h")
	fmt.Printf("%v\n%v", output, err)
}
