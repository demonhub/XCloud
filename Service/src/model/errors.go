package model

import (
	"errors"
)

var (

	// User
	LoginErr = errors.New("用户名/密码错误")

	// FILE ERRORs
	// 文件读取错误
	IOErr = errors.New("I/O Error")


	// DB ERRORs
	// 数据库打开错误
	DBErr = errors.New("DB Open Error")

	// 数据库写入错误
	OpErr = errors.New("Operate Error")

	// 数据库查询错误
	QueryErr = errors.New("Query Error")

)