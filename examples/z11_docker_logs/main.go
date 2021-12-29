package main

import (
	"fmt"
	"time"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

// 查看docker日志
func main() {
	ssh := zdpgo_ssh.New("192.168.31.112", "dell", "zhongkeqihang", 22)
	ssh.Connect()
	for {
		result, err := ssh.Sudo("docker logs --tail=5 hz_lagrange_acquisition")
		fmt.Println(result, err)
		time.Sleep(time.Second)
	}
}
