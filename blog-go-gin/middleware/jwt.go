package middleware

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	jwt "blog-go-gin/helper"
	"blog-go-gin/logging"
	"blog-go-gin/models/enum"
	"blog-go-gin/models/model"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

//如果是使用Go Module,gin-jwt模块应使用v2
//下载安装，开启Go Module "go env -w GO111MODULE=on",然后执行"go get github.com/appleboy/gin-jwt/v2"
//导入应写成 import "github.com/appleboy/gin-jwt/v2"
//如果不是使用Go Module
//下载安装gin-jwt，"go get github.com/appleboy/gin-jwt"
//导入import "github.com/appleboy/gin-jwt"

// JWT 注入IService
type JWT struct {
	UserAuthService service.IUserAuthService `inject:""`
	UserRoleService service.IUserRoleService `inject:""`
}

type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}

func NewJWT() *JWT {
	return &JWT{
		UserAuthService: impl.NewUserAuthServiceImpl(),
		UserRoleService: impl.NewUserRoleServiceImpl(),
	}
}

var (
	loginResponse *pb.LoginResponse
)

//GinJWTMiddlewareInit 初始化中间件
func (j *JWT) GinJWTMiddlewareInit() (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "blog-go",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 32,
		IdentityKey: common.IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			logging.Logger.Debug("执行PayloadFunc")
			if v, ok := data.(map[string]interface{}); ok {
				u, _ := v["user"].(*pb.UserAuth)
				r, _ := v["role"].(*pb.UserRole)
				return jwt.MapClaims{
					jwt.IdentityKey: u.UserInfoId,
					jwt.RoleIdKey:   r.RoleId,
					jwt.RoleKey:     enum.GetRoleKey(int(r.RoleId)).GetRoleZh(),
					jwt.UsernameKey: u.Username,
					jwt.RoleNameKey: enum.GetRoleKey(int(r.RoleId)).GetRoleCh(),
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			logging.Logger.Debug("执行IdentityHandler")
			claims := jwt.ExtractClaims(c)
			return map[string]interface{}{
				"IdentityKey": claims["identity"],
				"Username":    claims["username"],
				"RoleKey":     claims["rolekey"],
				"UserId":      claims["identity"],
				"RoleIds":     claims["roleid"],
				"RoleName":    claims["rolename"],
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			logging.Logger.Debug("执行Authenticator")
			//handles the login logic. On success LoginResponse is called, on failure Unauthorized is called
			body, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				logging.Logger.Error(err)
				return nil, err
			}
			request := &pb.RequestPkg{}
			if err = proto.Unmarshal(body, request); err != nil {
				logging.Logger.Error(err)
				return nil, err
			}
			logging.Logger.Debug(request)
			ok, err := j.UserAuthService.Login(request.User)
			if err != nil {
				return nil, errors.New(common.GetMsg(common.LoginFail))
			}
			if ok {
				//更新登录时间
				err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
					_ = model.UpdateUserLoginTime(tx, request.User.Username)
					return nil
				})
				if err != nil {
					return nil, err
				}
				loginResponse, err = j.UserAuthService.GetLoginResponse(request.User.Username)
				if err != nil {
					return nil, errors.New(common.GetMsg(common.GetUserInfoFail))
				}
				user, err := j.UserAuthService.GetUserAuthByUsername(request.User.Username)
				if err != nil {
					return nil, errors.New(common.GetMsg(common.GetUserInfoFail))
				}
				role := &pb.UserRole{
					RoleId:   user.RoleId,
					UserId:   user.UserInfoId,
					Username: user.Username,
				}
				m := map[string]interface{}{"user": user, "role": role}
				return m, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},

		LoginResponse: func(c *gin.Context, code int, token string, t time.Time) {
			loginResponse.Token = token
			c.Set("token", token)
			data := &pb.ResponsePkg{
				Code:          pb.ResultCode_SuccessOK,
				ServerTime:    time.Now().Unix(),
				Message:       "登录成功",
				LoginResponse: loginResponse,
			}
			c.ProtoBuf(code, data)
		},
		//receives identity and handles authorization logic
		Authorizator: func(data interface{}, c *gin.Context) bool {
			logging.Logger.Debug("执行Authorizator")
			if v, ok := data.(map[string]interface{}); ok {
				c.Set("role", v["RoleKey"])
				c.Set("roleIds", v["RoleIds"])
				c.Set("userId", v["UserId"])
				c.Set("username", v["Username"])
				return true
			}
			return false
		},
		//handles unauthorized logic
		Unauthorized: func(c *gin.Context, code int, message string) {
			data := &pb.ResponsePkg{
				Code:       pb.ResultCode(code),
				ServerTime: time.Now().Unix(),
				Message:    message,
			}
			c.ProtoBuf(http.StatusOK, data)
		},

		RefreshResponse: func(c *gin.Context, code int, message string, t time.Time) {
			logging.Logger.Debugf("code: %d,message: %s, time: %v", code, message, t)
			tokenString, expire, err := authMiddleware.RefreshToken(c)
			if err != nil {
				data := &pb.ResponsePkg{
					Code:       pb.ResultCode_SuccessOK,
					ServerTime: time.Now().Unix(),
					Message:    common.GetMsg(common.TokenHasExpired),
				}
				c.ProtoBuf(http.StatusOK, data)
				return
			}
			c.Set("token", tokenString)
			c.Set("expire", expire)
			data := &pb.ResponsePkg{
				Code:          pb.ResultCode_SuccessOK,
				ServerTime:    time.Now().Unix(),
				LoginResponse: &pb.LoginResponse{Token: tokenString},
			}
			c.ProtoBuf(http.StatusOK, data)
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		logging.Logger.Error("JWT Error:" + err.Error())
	}
	return
}

//NoRouteHandler 404 handler
//func NoRouteHandler(c *gin.Context) {
//	code := codes.PageNotFound
//	c.JSON(404, gin.H{"code": code, "message": codes.GetMsg(code)})
//}
