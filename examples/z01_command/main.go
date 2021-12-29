package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	ssh := zdpgo_ssh.New("192.168.18.101", "zhangdapeng", "zhangdapeng", 22)
	ssh.Connect()
	output, err := ssh.Run("free -h")
	fmt.Printf("%v\n%v", output, err)
}
