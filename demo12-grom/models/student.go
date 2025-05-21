package models

type Student struct {
	Id       int
	Name     string
	Number   string
	Password string
	ClassId  int
	// 多对多，一个学生拥有多个课程信息；一个课程有多个学生信息。`lesson_student` 是连接表
	Lesson []Lesson `gorm:"many2many:lesson_student;"`
}

func (Student) TableName() string {
	return "student"
}
