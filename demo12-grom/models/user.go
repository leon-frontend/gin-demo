// 文件名为 user.go ，与数据库中想要操作的表名相同
package models

// 定义一个结构体，默认操作数据库中的 users 表
type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// 修改结构体默认操作的表名
func (u User) TableName() string {
	return "user" // User 结构体默认操作 users 表，但此时操作 user 表名
}
