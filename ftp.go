package zdpgo_ssh

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// 上传文件
// @param localPath：本地文件路径
// @param remotePath：远程文件路径
func (s *SSH) UploadFile(localPath, remotePath string) {
	// 创建文件流
	srcFile, err := os.Open(localPath)
	if err != nil {
		fmt.Println("os.Open 创建要上传的文件失败 : ", localPath)
		log.Fatal(err)
	}
	defer srcFile.Close()

	// 远程文件名
	var remoteFileName = path.Base(localPath)

	// 创建目标文件
	dstFile, err := s.SFTPClient.Create(path.Join(remotePath, remoteFileName))
	if err != nil {
		fmt.Println("sftpClient.Create 创建目标文件失败 : ", path.Join(remotePath, remoteFileName))
		log.Fatal(err)
	}
	defer dstFile.Close()

	// 读取本地文件
	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll 读取本地文件失败 : ", localPath)
		log.Fatal(err)
	}

	// 写入目标文件
	dstFile.Write(ff)
	fmt.Println(localPath + " 上传到SFTP服务器成功。")
}

// 下载文件
// @param remotePath：远程文件路径
// @param localPath：本地文件路径
func (s *SSH) DownloadFile(remoteFilePath, localDirPath string) {
	// 创建文件流
	srcFile, err := s.SFTPClient.Open(remoteFilePath)
	if err != nil {
		fmt.Println("os.Open 创建要下载的文件失败 : ", remoteFilePath)
		log.Fatal(err)
	}
	defer srcFile.Close()

	// 本地文件名
	var localFileName = path.Base(remoteFilePath)

	// 创建目标文件
	dstFile, err := os.Create(path.Join(localDirPath, localFileName))
	if err != nil {
		fmt.Println("sftpClient.Create 创建目标文件失败 : ", path.Join(localDirPath, localFileName))
		log.Fatal(err)
	}
	defer dstFile.Close()

	// 读取本地文件
	ff, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println("ReadAll 读取远程文件失败 : ", remoteFilePath)
		log.Fatal(err)
	}

	// 写入目标文件
	dstFile.Write(ff)
	fmt.Println(remoteFilePath + " 从SFTP服务器下载到本地成功。")
}

// 上传文件夹
// @param localPath：本地文件夹
// @param remotePath：远程服务器路径
func (s *SSH) UploadDirectory(localPath, remotePath string) {
	// 读取本地文件夹
	localFiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal("读取文件夹列表失败： ", err)
		return
	}

	// 创建上级目录
	remoteDir := path.Join(remotePath, path.Base(localPath))
	s.SFTPClient.Mkdir(remoteDir)

	// 遍历所有文件
	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())  // 本地文件路径
		remoteFilePath := path.Join(remoteDir, backupDir.Name()) // 远程文件路径

		if backupDir.IsDir() { // 如果是文件夹，递归上传
			s.UploadDirectory(localFilePath, remoteFilePath)
		} else { // 如果是文件，上传文件
			s.UploadFile(localFilePath, remoteDir)
		}
	}

	fmt.Println(localPath + " 上传文件夹到SFTP服务器成功。")
}

// 下载文件夹
// @param remoteDirPath：远程服务器文件夹
// @param localDirPath：本地文件夹
func (s *SSH) DownloadDirectory(remoteDirPath, localDirPath string) {
	// 读取本地文件夹
	remoteFiles, err := s.SFTPClient.ReadDir(remoteDirPath)

	if err != nil {
		log.Fatal("读取文件夹列表失败： ", err)
		return
	}

	// 创建上级目录
	localDir := path.Join(localDirPath, path.Base(remoteDirPath))
	os.MkdirAll(localDir, 0766)
	os.Chmod(localDir, 0766) // 改变文件夹的权限

	// 遍历所有文件
	for _, backupDir := range remoteFiles {
		localFilePath := path.Join(localDir, backupDir.Name())       // 本地文件路径
		remoteFilePath := path.Join(remoteDirPath, backupDir.Name()) // 远程文件路径

		if backupDir.IsDir() { // 如果是文件夹，递归下载
			s.DownloadDirectory(remoteFilePath, localFilePath)
		} else { // 如果是文件，上传文件
			s.DownloadFile(remoteFilePath, localDir)
		}
	}

	fmt.Println(remoteDirPath + " 从SFTP服务器下载文件夹成功。")
}
