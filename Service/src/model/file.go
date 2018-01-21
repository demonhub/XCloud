package model

type CloudFile struct {
	Id 			string
	// fix for filename error
	OldName 	string
	NewName 	string
	// 状态
	State		int
	// 占用空间
	TakeSpace	int
	// 特征码 hash
	HashCode	string
}