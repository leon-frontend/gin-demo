package admin

import (
	"gin-demo/demo12-grom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (ac ArticleController) Index(c *gin.Context) {
	// 1. 查询 article 表中的所有数据
	// articleList := []models.Article{}
	// models.DB.Find(&articleList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": articleList,
	// })

	// 2. 查询文章时，获取文章对应的分类
	// articleList := []models.Article{}
	// // 获取所有文章数据，并预加载每篇文章关联的分类信息，即与分类表进行关联。
	// models.DB.Preload("ArticleCate").Find(&articleList)
	// c.JSON(http.StatusOK, gin.H{
	// 	"data": articleList,
	// })

	// 3. 查询文章分类，获取每个分类下的所有文章
	articleCateList := []models.ArticleCate{}
	models.DB.Preload("Article").Find(&articleCateList)
	c.JSON(http.StatusOK, gin.H{
		"data": articleCateList,
	})
}
