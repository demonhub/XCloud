package model

import (
	"errors"
)

var (

	// 文件读取错误
	IOErr = errors.New("I/O Error")

	// 数据库打开错误
	DBErr = errors.New("DB Open Error")

	// 数据库写入错误
	OpErr = errors.New("Operate Error")

	// 数据库查询错误
	QueryErr = errors.New("Query Error")

)