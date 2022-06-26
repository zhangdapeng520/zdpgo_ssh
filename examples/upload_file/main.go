package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.NewWithConfig(&zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}, zdpgo_log.Tmp)
	output, err := s.Sudo("ls -lah")
	fmt.Printf("%v\n%v", output, err)

	// 上传文件
	s.UploadFile("README.md", "README111.md")

	output, err = s.Sudo("ls -lah")
	fmt.Printf("%v\n%v", output, err)
}
