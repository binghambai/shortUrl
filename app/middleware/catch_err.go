package middleware

import (
	"binghambai.com/shortUrl/app/common"
	"binghambai.com/shortUrl/app/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func CatchError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); !utils.IsNil(r) {
				switch t := interface{}(r).(type) {
				case nil:
					break
				case *common.Response:
					log.Printf("panic: %v\n", t.Msg)
					common.WrapperContext(c).Error(500, t.Msg)
				default:
					log.Printf("panic: internal error; %v\n", t)
					common.WrapperContext(c).Error(500, "服务器内部错误")
				}
				c.Abort()
			}
		}()

		c.Next()
	}
}
