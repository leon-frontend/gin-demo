package admin

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (ac AdminController) UploadSingleFile(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/file.html", gin.H{
		"title": "首页",
	})
}

// 添加单个文件
func (ac AdminController) AddSingleFile(c *gin.Context) {
	// 获取表单中的用户名
	username := c.PostForm("username")

	// 根据 name 属性值获取上传文件
	avatarFile, _ := c.FormFile("avatar")

	// 获取文件后缀名，判断文件类型是否正确 .jpg .png .gif .jpeg
	fileExt := path.Ext(avatarFile.Filename) // 获取文件后缀名

	// 定义一个 Map 用于保存哪些文件后缀名符合要求
	allowExt := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	// 判断当前上传的文件名后缀是否合法
	if _, ok := allowExt[fileExt]; !ok {
		c.String(http.StatusOK, "上传文件格式不正确")
		return
	}

	// 以当前时间为文件名创建目录
	dir := path.Join("./static/upload", time.Now().Format("20060102"))
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		c.String(http.StatusOK, "创建目录失败")
		return
	}

	// 以当前时间戳给文件命名，并加上文件后缀名
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + fileExt

	// 将文件保存至 static/uploads/20060102 目录中
	dst := path.Join(dir, fileName)
	if err := c.SaveUploadedFile(avatarFile, dst); err != nil {
		fmt.Println("保存文件失败", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"message":  "上传成功",
		"username": username,
		"data":     dst,
	})
}

func (ac AdminController) UploadMultipleFiles(c *gin.Context) {
	c.HTML(http.StatusOK, "templates/multiple_files.html", gin.H{
		"title": "实现多个同名文件上传",
	})
}

// 添加多个同名文件
func (ac AdminController) AddMultipleSameNameFiles(c *gin.Context) {
	username := c.PostForm("username")
	form, _ := c.MultipartForm()         // 获取表单
	avatarFiles := form.File["avatar[]"] // 获取多个上传文件

	// 遍历 avatarFiles 获取多个上传文件
	for _, file := range avatarFiles {
		// 将文件保存到文件指定路径
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"message":  "上传成功",
		"username": username,
	})
}
