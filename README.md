# zdpgo_ssh

使用Golang执行ssh命令

## 版本历史

- v0.1.1 2022/06/26 升级：日志组件升级

## 快速入门

### 执行shell命令

```go
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
```

### 执行sudo命令

```go
package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.NewWithConfig(&zdpgo_ssh.Config{
		Host:     "192.168.33.10",
		Port:     22,
		Username: "zhangdapeng",
		Password: "zhangdapeng",
	}, zdpgo_log.Tmp)
	output, err := s.Sudo("ls -lah")
	fmt.Printf("%v\n%v", output, err)
}
```
