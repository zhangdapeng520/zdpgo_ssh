package zdpgo_ssh

import "testing"

// 上传文件
func MTestUploadFile(t *testing.T) {
	ssh := SSH{
		Host:     "192.168.18.11", // 主机
		Port:     22,              // 端口
		Username: "zhangdapeng",   // 用户名
		Password: "zhangdapeng",   // 密码
	}
	ssh.Connect()
	ssh.UploadFile("README.md", "/home/zhangdapeng")
}

// 上传文件夹
func TestUploadDirectory(t *testing.T) {
	ssh := SSH{
		Host:     "192.168.18.11", // 主机
		Port:     22,              // 端口
		Username: "zhangdapeng",   // 用户名
		Password: "zhangdapeng",   // 密码
	}
	ssh.Connect()
	ssh.UploadDirectory("./test", "/home/zhangdapeng")
}
