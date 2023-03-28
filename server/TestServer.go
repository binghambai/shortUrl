package server

import (
	"binghambai.com/shortUrl/vo"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, vo.TestResp{
		Res: "success , Hello World",
	})
}
