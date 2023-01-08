package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	var (
		user = "root"
		host = "192.168.1.81"
		port = 22
	)
	ssh, err := zdpgo_ssh.NewWithPublicKey(user, host, port)
	if err != nil {
		panic(err)
	}
	result, err := ssh.Run("pwd")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
