package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	// 创建对象
	s := zdpgo_ssh.NewWithConfig(&zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}, zdpgo_log.Tmp)

	// 进行连接
	output, err := s.Run("free -h")

	// 查看命令结果
	fmt.Printf("%v\n%v", output, err)

	// 查看健康状态
	fmt.Println(s.Status())
}
