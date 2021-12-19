# zdpgo_ssh
使用Golang执行ssh命令

功能列表：
- 在go代码里连接linux系统，执行shell命令
- FTP文件上传和下载
- FTP文件夹上传和下载
- 执行sudo命令
- 批量执行ssh命令

zdpgo_ssh框架使用教程（官方博客）：https://blog.csdn.net/qq_37703224/category_11546273.html

## 一、快速入门

### 1.1 执行shell命令
```go
package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	output, err := s.Run("free -h")
	fmt.Printf("%v\n%v", output, err)
}
```

### 1.2 交互式的shell
```go
package main

import (
	"os"

	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	s := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	s.RunTerminal("top", os.Stdout, os.Stdin)
}
```

### 1.3 上传文件
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	ssh := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	ssh.Connect()
	ssh.UploadFile("README.md", "/home/zhangdapeng")
}
```

### 1.4 上传文件夹
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	ssh := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	ssh.Connect()
	ssh.UploadDirectory("./test", "/home/zhangdapeng")
}
```

### 1.5 下载文件
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	ssh := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	ssh.Connect()
	ssh.DownloadFile("/home/zhangdapeng/README.md", "./examples/z05_download_file/")
}
```

### 1.6 下载文件夹
```go
package main

import (
	"github.com/zhangdapeng520/zdpgo_ssh"
)

func main() {
	ssh := zdpgo_ssh.New("192.168.18.11", "zhangdapeng", "zhangdapeng", 22)
	ssh.Connect()
	ssh.DownloadDirectory("/home/zhangdapeng/test", "./examples/z06_download_dir/")
}
```

