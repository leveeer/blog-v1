package adminApi

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/logging"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	MenuService service.IMenuService = impl.NewMenuServiceImpl()
)

type MenuRestApi struct {
	base.Handler
}

func NewMenuRestApi() *MenuRestApi {
	return &MenuRestApi{}
}

func (c *MenuRestApi) GetUserMenus(ctx *gin.Context) {
	roleId, exists := ctx.Get("roleIds")
	if !exists {
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	userMenus, err := MenuService.GetUserMenus(int(roleId.(float64)))
	logging.Logger.Debug(userMenus)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetUserMenusFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		UserMenu:   userMenus,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}
