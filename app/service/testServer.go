package server

import (
	"binghambai.com/shortUrl/app/common"
	"binghambai.com/shortUrl/app/conn"
	"binghambai.com/shortUrl/app/dao"
	"binghambai.com/shortUrl/app/model"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	value := dao.Get("test_go")
	if value == nil {
		common.WrapperContext(c).Error(500, "The key is not exist")
		return
	}

	db := conn.Db
	var shortUrl model.ShortUrl
	db.First(&shortUrl, 838092237798912000)

	common.WrapperContext(c).Success(shortUrl)
}
