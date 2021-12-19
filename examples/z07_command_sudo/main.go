package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.New("192.168.18.101", "zhangdapeng", "zhangdapeng", 22)
	output, err := s.Sudo("./install_docker.sh")
	fmt.Printf("%v\n%v", output, err)
}
