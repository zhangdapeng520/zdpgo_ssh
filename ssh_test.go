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
