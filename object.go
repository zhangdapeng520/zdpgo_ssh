package zdpgo_ssh

import (
	"fmt"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// SSH连接对象
type SSH struct {
	Host       string       // 主机
	Port       int          // 端口
	Username   string       // 用户名
	Password   string       // 密码
	SFTPClient *sftp.Client // sftp客户端
	SSHClient  *ssh.Client  // ssh客户端
	LastResult string       //最近一次Run的结果
}

//创建SSH对象
//@param host 主机地址
//@param username 用户名
//@param password 密码
//@param port 端口号,默认22
func New(host, username, password string, port ...int) *SSH {
	s := new(SSH)
	s.Host = host
	s.Username = username
	s.Password = password
	if len(port) <= 0 {
		s.Port = 22
	} else {
		s.Port = port[0]
	}
	return s
}

// 建立连接
func (s *SSH) Connect() {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		err          error
	)

	// 获取权限方法
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(s.Password))
	clientConfig = &ssh.ClientConfig{
		User:            s.Username,
		Auth:            auth,
		Timeout:         30 * time.Second,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //ssh.FixedHostKey(hostKey),
	}

	// 连接到SSH
	addr = fmt.Sprintf("%s:%d", s.Host, s.Port)
	if s.SSHClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		panic(err)
	}

	// 创建sftp客户端
	if s.SFTPClient, err = sftp.NewClient(s.SSHClient); err != nil {
		panic(err)
	}
}
