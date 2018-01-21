package model

type User struct {
	Id 			string
	// 用户名
	Username 	string
	// 两次md5后的密码
	Password 	string
	// 权限
	Privilege 	int
	// 状态
	// 0 - 正常
	// 1 - 受限(仅能下载)
	State		int
	// 空间使用情况
	// 单位 1KB
	EmptySpace	int
	UseSpace	int
	// 分享码
	ShareCode	string
}