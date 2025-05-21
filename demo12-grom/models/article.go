package models

type Article struct {
	Id    int
	Title string
	// ArticleCateId int // foreign key, 前缀必须和 ArticleCate 相同或者使用`gorm:xxx`重写外键
	CateId int
	State  int
	// ArticleCate ArticleCate `gorm:"foreignKey:CateId"` // 与 ArticleCate 表进行关联，并且变量名用于 Preload("") 方法的参数
}

// 修改结构体默认操作的表名
func (Article) TableName() string {
	return "article" // User 结构体默认操作 users 表，但此时操作 user 表名
}
