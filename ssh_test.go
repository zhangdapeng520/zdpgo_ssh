package zdpgo_ssh

import (
	"fmt"
	"testing"
)

func TestSSH_Run(t *testing.T) {
	s := NewWithConfig(&Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	})
	output, err := s.Run("free -h")
	fmt.Printf("%v\n%v", output, err)
}

// 测试通过公钥连接
func TestSSH_NewWithPublicKey(t *testing.T) {
	var (
		user = "root"
		host = "192.168.1.81"
		port = 22
	)
	ssh, err := NewWithPublicKey(user, host, port)
	if err != nil {
		t.Error(err)
	}
	result, err := ssh.Run("pwd")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
