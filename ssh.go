package zdpgo_ssh

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ssh/sftp"
	"io"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSH SSH连接对象
type SSH struct {
	Client *ssh.Client // ssh客户端
	Config *Config
}

//New 创建SSH对象
//@param host 主机地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func New() *SSH {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *SSH {
	s := &SSH{}

	// 配置
	s.Config = config

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
		return "", err
	}
	defer session.Close()

	// 执行命令
	buf, err := session.CombinedOutput(command)
	if err != nil {
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
		return "", err
	}
	defer session.Close()

	// 执行命令
	command = fmt.Sprintf("echo %s | sudo -S %s", s.Config.Password, command)
	buf, err := session.CombinedOutput(command)
	if err != nil {
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

// GetSshConfig 获取SSH客户端配置
func (s *SSH) GetSshConfig() *ssh.ClientConfig {
	sshConfig := &ssh.ClientConfig{
		User: s.Config.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.Config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		ClientVersion:   "",
		Timeout:         10 * time.Second,
	}
	return sshConfig
}

// UploadFile 上传文件
func (s *SSH) UploadFile(localFileName, remoteFileName string) FileResult {
	result := FileResult{
		LocalFileName:  localFileName,
		RemoteFileName: remoteFileName,
	}

	// 创建客户端
	sshConfig := s.GetSshConfig()

	//建立与SSH服务器的连接
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port), sshConfig)
	if err != nil {
		return result
	}
	defer sshClient.Close()

	// 获取SFTP客户端
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return result
	}
	defer sftpClient.Close()

	//获取当前目录
	cwd, err := sftpClient.Getwd()
	if err != nil {
		return result
	}

	//上传文件(将本地文件通过sftp传到远程服务器)
	remoteFile, err := sftpClient.Create(sftp.Join(cwd, remoteFileName))
	if err != nil {
		return result
	}
	defer remoteFile.Close()

	//打开本地文件
	localFile, err := os.Open(localFileName)
	if err != nil {
		return result
	}
	defer localFile.Close()

	// 本地文件流拷贝到上传文件流
	n, err := io.Copy(remoteFile, localFile)
	if err != nil {
		return result
	}

	// 获取本地文件大小
	localFileInfo, err := os.Stat(localFileName)
	if err != nil {
		return result
	}

	// 上传结果
	result.Status = true
	result.LocalFileSize = uint64(localFileInfo.Size())
	result.RemoteFileSize = uint64(n)

	return result
}

// DownloadFile 下载文件
func (s *SSH) DownloadFile(remoteFileName, localFileName string) FileResult {
	result := FileResult{
		LocalFileName:  localFileName,
		RemoteFileName: remoteFileName,
	}

	// 创建客户端
	sshConfig := s.GetSshConfig()

	//建立与SSH服务器的连接
	sshClient, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port), sshConfig)
	if err != nil {
		return result
	}
	defer sshClient.Close()

	// 获取SFTP客户端
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		return result
	}
	defer sftpClient.Close()

	//下载文件,将远程服务器的/bin/bash文件下载到本地
	remoteFile, err := sftpClient.Open(remoteFileName)
	if err != nil {
		return result
	}
	defer remoteFile.Close()

	localFile, err := os.Create(localFileName)
	if err != nil {
		return result
	}
	defer localFile.Close()

	// 将远程文件复制到本地文件流
	n, err := io.Copy(localFile, remoteFile)
	if err != nil {
		return result
	}

	//获取远程文件大小
	remoteFileInfo, err := sftpClient.Stat(remoteFileName)
	if err != nil {
		return result
	}

	// 下载结果
	result.Status = true
	result.LocalFileSize = uint64(n)
	result.RemoteFileSize = uint64(remoteFileInfo.Size())

	return result
}

// FormatFileSize 字节的单位转换 保留两位小数
func (ssh *SSH) FormatFileSize(s int64) (size string) {
	if s < 1024 {
		return fmt.Sprintf("%.2fB", float64(s)/float64(1))
	} else if s < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(s)/float64(1024))
	} else if s < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(s)/float64(1024*1024))
	} else if s < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(s)/float64(1024*1024*1024))
	} else if s < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(s)/float64(1024*1024*1024*1024))
	} else { //if s < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(s)/float64(1024*1024*1024*1024*1024))
	}
}
