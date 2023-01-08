# zdpgo_ssh

使用Golang执行ssh命令

## 版本历史

- v1.0.3 2022/06/26 升级：日志组件升级
- v1.0.4 2022/06/26 新增：文件上传和下载
- v1.0.5 2022/06/26 优化：移除日志
- v1.0.6 2023/01/08 新增：通过公钥连接SSH，即就是免密登录

## 快速入门

### 执行shell命令

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	// 创建对象
	s := zdpgo_ssh.NewWithConfig(&zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	})

	// 进行连接
	output, err := s.Run("free -h")

	// 查看命令结果
	fmt.Printf("%v\n%v", output, err)

	// 查看健康状态
	fmt.Println(s.Status())
}
```

### 执行sudo命令

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.NewWithConfig(&zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	})
	output, err := s.Sudo("ls -lah")
	fmt.Printf("%v\n%v", output, err)
}
```

## 文件上传

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.NewWithConfig(&zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	})
	output, err := s.Sudo("ls -lah")
	fmt.Printf("%v\n%v", output, err)

	// 上传文件
	s.UploadFile("README.md", "README111.md")

	output, err = s.Sudo("ls -lah")
	fmt.Printf("%v\n%v", output, err)
}
```

## 文件下载

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.NewWithConfig(&zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	})
	output, err := s.Sudo("ls -lah")
	fmt.Printf("%v\n%v", output, err)

	// 下载文件
	s.DownloadFile("README111.md", "README111.md")

	output, err = s.Sudo("ls -lah")
	fmt.Printf("%v\n%v", output, err)
}
```

# 公共公钥连接SSH

## Windows配置SSH免密登录

这一步很重要，大家可以自行谷歌解决。如果嫌麻烦，也可以关注我的公众号“Python私教”，里面有详细的配置教程和zdpgo_ssh库的使用说明。

## 使用示例

本示例的地址位于本开源项目的：examples/new_with_public_key/main.go

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	var (
		user = "root"
		host = "192.168.1.81"
		port = 22
	)
	ssh, err := zdpgo_ssh.NewWithPublicKey(user, host, port)
	if err != nil {
		panic(err)
	}
	result, err := ssh.Run("pwd")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
```