package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

// 与112服务器连接，并查看其状态
func main() {
	ssh := zdpgo_ssh.New("192.168.31.112", "dell", "zhongkeqihang", 22)
	// ssh := zdpgo_ssh.New("192.168.18.101", "zhangdapeng", "zhangdapeng", 22)
	ssh.Connect()
	output, err := ssh.Run("echo 'ok'")
	fmt.Printf("%v\n%v", output, err)

	// 查看健康状态
	fmt.Println(ssh.Status())
}
