package admin

import (
	"gin-demo/demo12-grom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentController struct{}

func (sc StudentController) Index(c *gin.Context) {
	// 1. 获取所有学生的信息
	// studentList := []models.Student{}
	// models.DB.Find(&studentList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": studentList,
	// })

	// 2. 获取所有的课程信息
	// lessonList := []models.Lesson{}
	// models.DB.Find(&lessonList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": lessonList,
	// })

	// 3. 查询学生信息的时候，展示学生选修的课程信息
	// studentList := []models.Student{}
	// models.DB.Preload("Lesson").Find(&studentList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": studentList,
	// })

	// 4. 查询张三选修的课程信息
	// studentList := []models.Student{}
	// models.DB.Preload("Lesson").Where("name = ?", "张三").Find(&studentList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": studentList,
	// })

	// 5. 查询某个课程以及选修该课程的学生信息
	lessonList := []models.Lesson{}
	models.DB.Preload("Student").Where("id = ?", 1).Find(&lessonList)
	c.JSON(http.StatusOK, gin.H{
		"data": lessonList,
	})
}
