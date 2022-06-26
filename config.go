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

type FileResult struct {
	Status         bool   `yaml:"status" json:"status"`
	LocalFileName  string `yaml:"local_file_name" json:"local_file_name"`
	LocalFileSize  uint64 `yaml:"local_file_size" json:"local_file_size"`
	RemoteFileName string `yaml:"remote_file_name" json:"remote_file_name"`
	RemoteFileSize uint64 `yaml:"remote_file_size" json:"remote_file_size"`
}
