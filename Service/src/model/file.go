package model

import (
	"crypto/md5"
	"os"
	"io"
)

type CloudFile struct {
	Id 			string
	// 文件类型
	Type 		string
	// 进行名字替换防止出现重名问题
	SrcName 	string
	NewName 	string
	// 存储目录
	Path  		string
	// 状态
	// 0 - 正常
	// 1 - 存在问题
	State		int
	// 占用空间
	// 单位1KB
	TakeSpace	int
	// 特征码
	// hash值
	HashCode	string
}

func (file *CloudFile)getFullPath() string {
	return file.Path + file.NewName
}

func (file *CloudFile)setHashCode() {
	hash := md5.New()

	if fileIn, err := os.Open(file.getFullPath());err != nil {
		IOErr.Error()
		return
	} else {
		io.Copy(hash, fileIn)
	}

	file.HashCode = (string)(hash.Sum(nil))
}