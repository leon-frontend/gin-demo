package admin

import (
	"gin-demo/demo12-grom/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// 查询 Users 表中的全部数据
func queryAll(userList *[]models.User) {
	models.DB.Find(&userList)
}

// 根据条件查询 Users 表中的某些数据
func queryByWhere(userList *[]models.User) {
	// 1. 查询 age 大于 20 的数据。其中，? 相当于占位符
	models.DB.Where("age > ?", 30).Find(&userList)
}

func (uc UserController) Index(c *gin.Context) {
	// 1. 定义一个 user 类型的切片，用于保存数据库中的数据
	userList := []models.User{}

	// 2.1 检索 users 表中的全部对象，并保存到 userList 中
	queryAll(&userList)
	// queryByWhere(&userList)

	// 3. 向前端返回 JSON 数据
	c.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}

// 向数据库中新增数据
func (uc UserController) Add(c *gin.Context) {
	// 1. 实例化结构体，注意不用添加 ID 数据
	user := models.User{
		Username: "CCNU",
		Age:      18,
		Email:    "ccnu@163.com",
		AddTime:  int(time.Now().Unix()),
	}

	// 2. 向数据库中新增数据
	models.DB.Create(&user)

	// 3. 向浏览器返回 JSON 数据
	c.String(http.StatusOK, "添加数据成功")
}

func (uc UserController) Update(c *gin.Context) {
	// 1. 修改单个数据中的所有字段或指定字段
	// 1.1 获取需要修改的数据
	// user := models.User{Id: 6} // 实例化结构体，并指明 Id 数据
	// models.DB.Find(&user)      // 获取 id = 6 的数据，并赋值到结构体中

	// // 1.2 修改结构体中的数据，即修改 id = 6 的数据
	// user.Username = "小红"
	// user.Email = "xiaohong@163.com"

	// // 1.3 更新数据库中 id = 6 的数据
	// models.DB.Save(&user)

	// 2. 使用 Model 函数和 Update 函数更新单列数据
	user := models.User{} // 实例化时不传递 Id 值，在 Where 函数中传递 id 值
	// 在查询时需要使用 Where 语句传递过滤条件，可以阻止全局更新
	models.DB.Model(&user).Where("id = ?", 6).Update("username", "小蓝") // 更新 id = 6 的数据的 username 列

	// 向浏览器返回信息
	c.String(http.StatusOK, "更新数据成功")
}

func (uc UserController) Delete(c *gin.Context) {
	user := models.User{Id: 6} // 实例化 id = 6 的 User 结构体
	models.DB.Delete(&user)    // 删除数据库中 id = 6 的数据
	c.String(http.StatusOK, "删除用户")
}
