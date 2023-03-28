package server

import (
	"binghambai.com/shortUrl/middleware"
	"binghambai.com/shortUrl/vo"
	"github.com/gin-gonic/gin"
	"log"
)

func Test(c *gin.Context) {
	value, err := middleware.RedisGet("test_go")
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, vo.TestResp{
		Res: "success , Hello World, get vale:" + value,
	})
}
