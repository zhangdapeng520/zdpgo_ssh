package zdpgo_ssh

import (
	"strings"
)

// 获取服务器的健康状态
func (ssh *SSH) Status() bool {
	ssh.Connect()
	output, err := ssh.Run("echo 'ok'")
	flag := err == nil && strings.TrimSpace(output) == "ok"
	return flag
}
