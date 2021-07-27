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
	RoleService service.IRoleService = impl.NewRoleServiceImpl()
)

type RoleRestApi struct {
	base.Handler
}

func NewRoleRestApi() *RoleRestApi {
	return &RoleRestApi{}
}

func (a *RoleRestApi) GetAdminUsersRole(ctx *gin.Context) {
	roles, err := RoleService.GetAdminUsersRole()
	if err != nil {
		a.ProtoBufFail(ctx, http.StatusOK, common.GetRolesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		AdminRoles: roles,
	}
	a.ProtoBuf(ctx, http.StatusOK, data)
}

func (a *RoleRestApi) GetRoles(ctx *gin.Context) {
	type Condition struct {
		Current  int64  `json:"current" form:"current"`
		Size     int32  `json:"size" form:"size"`
		Keywords string `json:"keywords" form:"keywords"`
	}
	var condition Condition
	err := ctx.ShouldBind(&condition)
	if err != nil {
		logging.Logger.Error(err)
		a.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(condition)

	adminRoles, err := RoleService.GetRoles(&pb.CsCondition{
		Current:  condition.Current,
		Size:     condition.Size,
		Keywords: condition.Keywords,
	})
	if err != nil {
		logging.Logger.Error(err)
		a.ProtoBufFail(ctx, http.StatusOK, common.GetRolesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Roles:      adminRoles,
	}
	a.ProtoBuf(ctx, http.StatusOK, data)
}
