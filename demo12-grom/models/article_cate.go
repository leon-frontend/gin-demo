package models

type ArticleCate struct {
	Id    int // Primary key
	Title string
	State int
	// has many 一对多关系，分类下面拥有多篇文章。并且配置 Article 表的外键
	Article []Article `gorm:"foreignKey:CateId"`
}

// 修改结构体默认操作的表名
func (ArticleCate) TableName() string {
	return "article_cate"
}
