package helper

import (
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/enum"
	"github.com/gin-gonic/gin"
)

//IAuthorizator 授权规则接口
type IAuthorizator interface {
	HandleAuthorizator(data interface{}, c *gin.Context) bool
}

//AdminAuthorizator 管理员授权规则
type AdminAuthorizator struct {
}

//HandleAuthorizator 处理管理员授权规则
func (*AdminAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*pb.UserRole); ok && int(v.RoleId) == enum.Admin.GetRoleId() {
		logging.Logger.Debug(v.RoleId)
		return true
	}
	return false
}

//TestAuthorizator 测试用户授权规则
type TestAuthorizator struct {
}

//HandleAuthorizator 处理测试用户授权规则
func (*TestAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*pb.UserRole); ok && int(v.RoleId) == enum.Test.GetRoleId() {
		logging.Logger.Debug(v.RoleId)
		return true
	}
	return false
}

//UserAuthorizator 普通用户授权规则
type UserAuthorizator struct {
}

//HandleAuthorizator 处理普通用户授权规则
func (*UserAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	logging.Logger.Debug(data)
	if v, ok := data.(*pb.UserRole); ok && int(v.RoleId) == enum.User.GetRoleId() {
		logging.Logger.Debug(v.RoleId)
		return true
	}
	return false
}

//AllUserAuthorizator 普通用户授权规则
type AllUserAuthorizator struct {
}

//HandleAuthorizator 处理普通用户授权规则
func (*AllUserAuthorizator) HandleAuthorizator(data interface{}, c *gin.Context) bool {
	logging.Logger.Debug("放行")
	return true
}
