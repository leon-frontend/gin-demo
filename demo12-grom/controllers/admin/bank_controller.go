package admin

import (
	"gin-demo/demo12-grom/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BankController struct{}

func (bc BankController) success(c *gin.Context) {
	c.String(http.StatusOK, "转账成功")
}

func (bc BankController) error(c *gin.Context) {
	c.String(http.StatusOK, "转账失败")
}

func (bc BankController) Index(c *gin.Context) {

	// 1. 开启事务
	tx := models.DB.Begin()

	// 2. 处理异常，执行回滚事务
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			bc.error(c)
			return
		}
	}()

	// 3. 实现张三给李四转账（使用事务）
	u1 := models.Bank{}
	tx.Where("username = ?", "张三").Find(&u1) // 查询张三的数据

	// 更新张三的数据
	u1.Balance -= 100
	if err := tx.Save(&u1).Error; err != nil {
		tx.Rollback()
		bc.error(c)
		return
	}

	// panic("模拟异常")

	// 更新李四的数据
	u2 := models.Bank{}
	tx.Where("username = ?", "李四").Find(&u2)
	u2.Balance += 100
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		bc.error(c)
		return
	}

	tx.Commit()

	bc.success(c)
}
