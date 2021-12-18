package zdpgo_ssh

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

//执行sh命令
//@param command 命令
func (s *SSH) Run(command string) (string, error) {

	// 创建连接
	if s.SSHClient == nil {
		s.Connect()
	}

	// 创建session
	session, err := s.SSHClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// 执行命令
	buf, err := session.CombinedOutput(command)

	// 记录最后一次的执行命令结果
	s.LastResult = string(buf)

	// 返回命令执行结果
	return s.LastResult, err
}

//使用管理员身份执行sh命令
//@param command 命令
func (s *SSH) Sudo(command string) (string, error) {

	// 创建连接
	if s.SSHClient == nil {
		s.Connect()
	}

	// 创建session
	session, err := s.SSHClient.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	// 执行命令
	command = fmt.Sprintf("echo %s | sudo -S %s", s.Password, command)
	buf, err := session.CombinedOutput(command)

	// 记录最后一次的执行命令结果
	s.LastResult = string(buf)

	// 返回命令执行结果
	return s.LastResult, err
}

//执行带交互的命令
//@param command 命令
//@param stdout 标准输出
//@param stderr 标准错误输出
func (s *SSH) RunTerminal(command string, stdout, stderr io.Writer) error {
	// 创建连接
	if s.SSHClient == nil {
		s.Connect()
	}

	// 创建session
	session, err := s.SSHClient.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// 输入
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd, oldState)

	// 输出
	session.Stdout = stdout
	session.Stderr = stderr
	session.Stdin = os.Stdin

	// 获取命令行的宽高
	termWidth, termHeight, err := term.GetSize(fd)
	if err != nil {
		panic(err)
	}

	// 设置命令行模式
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// 获取命令行
	if err := session.RequestPty("xterm-256color", termHeight, termWidth, modes); err != nil {
		return err
	}

	// 执行命令
	session.Run(command)
	return nil
}
