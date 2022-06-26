package zdpgo_ssh

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSH SSH连接对象
type SSH struct {
	Client *ssh.Client // ssh客户端
	Log    *zdpgo_log.Log
	Config *Config
}

//New 创建SSH对象
//@param host 主机地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func New(log *zdpgo_log.Log) *SSH {
	return NewWithConfig(&Config{}, log)
}

func NewWithConfig(config *Config, log *zdpgo_log.Log) *SSH {
	s := &SSH{}

	// 配置
	s.Config = config

	// 日志
	s.Log = log

	// 返回
	return s
}

// Connect 建立连接
func (s *SSH) Connect() error {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		err          error
	)

	// 获取权限方法
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(s.Config.Password))
	clientConfig = &ssh.ClientConfig{
		User:            s.Config.Username,
		Auth:            auth,
		Timeout:         3600 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接到SSH
	addr = fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port)
	if s.Client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		s.Log.Error("与服务器建立SSH连接失败", "error", err)
		return err
	}

	// 返回
	return nil
}

// Run 执行sh命令
// @param command 命令
func (s *SSH) Run(command string) (string, error) {

	// 创建连接
	if s.Client == nil {
		err := s.Connect()
		if err != nil {
			return "", err
		}
	}

	// 创建session
	session, err := s.Client.NewSession()
	if err != nil {
		s.Log.Error("创建session失败", "error", err)
		return "", err
	}
	defer session.Close()

	// 执行命令
	buf, err := session.CombinedOutput(command)
	if err != nil {
		s.Log.Error("执行命令失败", "error", err)
		return "", err
	}

	// 返回命令执行结果
	return string(buf), nil
}

// Sudo 使用管理员身份执行sh命令
// @param command 命令
func (s *SSH) Sudo(command string) (string, error) {

	// 创建连接
	if s.Client == nil {
		err := s.Connect()
		if err != nil {
			return "", err
		}
	}

	// 创建session
	session, err := s.Client.NewSession()
	if err != nil {
		s.Log.Error("创建session失败", "error", err)
		return "", err
	}
	defer session.Close()

	// 执行命令
	command = fmt.Sprintf("echo %s | sudo -S %s", s.Config.Password, command)
	buf, err := session.CombinedOutput(command)
	if err != nil {
		s.Log.Error("执行命令失败", "error", err)
		return "", err
	}

	// 返回命令执行结果
	return string(buf), nil
}

// Status 获取服务器的健康状态
func (ssh *SSH) Status() bool {
	// 连接
	err := ssh.Connect()
	if err != nil {
		return false
	}

	// 执行简单的命令
	output, err := ssh.Run("echo 'ok'")
	flag := err == nil && strings.TrimSpace(output) == "ok"
	return flag
}
