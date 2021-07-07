package middleware

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/helper"
	"fmt"
	"net/http"
	"time"

	jwt "blog-go-gin/helper"
	"github.com/gin-gonic/gin"
)

// AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get(jwt.JwtPayloadKey)
		v := data.(jwt.MapClaims)
		e := helper.Casbin()
		//检查权限
		ok, err := e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		common.HasError(err, "", 500)
		fmt.Printf("[INFO] %s %s %s \r\n", c.Request.Method, c.Request.URL.Path, v["rolekey"])
		if !ok {
			data := &pb.ResponsePkg{
				Code:       pb.ResultCode_Forbidden,
				ServerTime: time.Now().Unix(),
				Message:    common.GetMsg(common.AdminPrivilegeNeeded),
			}
			c.ProtoBuf(http.StatusOK, data)
			c.Abort()
		}
		c.Next()
	}
}
