package main

import (
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	ssh := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	ssh.Connect()
	ssh.UploadDirectory("./test", "/home/zhangdapeng")
}
