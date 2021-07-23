package usersApi

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
	UserRoleService service.IUserRoleService = impl.NewUserRoleServiceImpl()
)

type UserRoleRestApi struct {
	base.Handler
}

func (a *UserRoleRestApi) UpdateUserRole(ctx *gin.Context) {
	requestPkg, err := a.ReadRequestBody(ctx)
	if err != nil {
		logging.Logger.Error(err)
		a.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	err = UserRoleService.UpdateUserRole(requestPkg.UserRole)
	if err != nil {
		logging.Logger.Error(err)
		a.ProtoBufFail(ctx, http.StatusOK, common.UpdateUserRoleFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    "更新成功",
	}
	a.ProtoBuf(ctx, http.StatusOK, data)
}

func NewUserRoleRestApi() *UserRoleRestApi {
	return &UserRoleRestApi{}
}
