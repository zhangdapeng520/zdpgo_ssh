package zdpgo_ssh

/*
@Time : 2022/6/26 11:18
@Author : 张大鹏
@File : config
@Software: Goland2021.3.1
@Description:
*/

type Config struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}
