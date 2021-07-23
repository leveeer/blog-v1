package impl

import (
	"blog-go-gin/common"
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/enum"
	"blog-go-gin/models/model"
	"blog-go-gin/models/page"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jordan-wright/email"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"net/smtp"
	"strconv"
	"sync"
	"time"
)

type UserAuthServiceImpl struct {
	wg sync.WaitGroup
}

func (u *UserAuthServiceImpl) GetAdminUsers(c *pb.CsCondition) (*pb.ScAdminUsers, error) {
	users, err := model.GetUsersByConditionWithPage(c.GetKeywords(), &page.IPage{Current: int(c.GetCurrent()), Size: int(c.GetSize())}, "%"+c.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	var userSlice []*pb.ScUsers
	var isDisable int32
	for _, user := range users {
		logging.Logger.Debug(user)
		if user.IsDisable {
			isDisable = 1
		} else {
			isDisable = 0
		}
		roles, err := model.GetUserRoles("user_id = ?", user.UserInfoID)
		if err != nil {
			return nil, err
		}
		var roleSlice []*pb.ScUserRole
		for _, role := range roles {
			roleSlice = append(roleSlice, &pb.ScUserRole{
				Id:       int64(role.RoleID),
				RoleName: enum.GetRoleKey(role.RoleID).GetRoleCh(),
			})
		}
		userSlice = append(userSlice, &pb.ScUsers{
			Id:            int64(user.ID),
			UserInfoId:    int64(user.UserInfoID),
			IpAddr:        user.IPAddr,
			IpSource:      user.IPSource,
			Nickname:      user.Nickname,
			Avatar:        user.Avatar,
			LoginType:     int32(user.LoginType),
			CreateTime:    user.CreateTime,
			LastLoginTime: user.LastLoginTime,
			IsDisable:     isDisable,
			UserRoleList:  roleSlice,
		})
	}
	usersCount, err := model.GetUsersCountByCondition(c.GetKeywords(), "%"+c.GetKeywords()+"%")
	if err != nil {
		return nil, err
	}
	return &pb.ScAdminUsers{
		UserList: userSlice,
		Count:    usersCount,
	}, nil
}

func (u *UserAuthServiceImpl) GetLoginResponse(username string) (*pb.LoginResponse, error) {
	userAuth, err := model.GetLoginResponse(username)
	if err != nil {
		return nil, err
	}
	//获取用户点赞的文章id集合
	var articleLikeSet []int32
	likeArticleIds, err := common.GetRedisUtil().HashGet(common.ArticleUserLike, strconv.Itoa(userAuth.UserInfoID))
	if err != nil && err != redis.Nil {
		logging.Logger.Debug(err)
		return nil, err
	} else {
		articleLikeSet = []int32{}
	}
	if likeArticleIds != "" {
		err = json.Unmarshal([]byte(likeArticleIds), &articleLikeSet)
		if err != nil {
			logging.Logger.Debug(err)
			return nil, err
		}
	}
	//获取用户点赞的评论d集合
	var commentLikeSet []int32
	likeCommentIds, err := common.GetRedisUtil().HashGet(common.CommentUserLike, strconv.Itoa(userAuth.UserInfoID))
	if err != nil && err != redis.Nil {
		logging.Logger.Debug(err)
		return nil, err
	} else {
		commentLikeSet = []int32{}
	}
	if likeCommentIds != "" {
		err = json.Unmarshal([]byte(likeCommentIds), &commentLikeSet)
		if err != nil {
			logging.Logger.Debug(err)
			return nil, err
		}
	}

	return &pb.LoginResponse{
		UserId:         int32(userAuth.UserInfoID),
		Email:          userAuth.Username,
		NickName:       userAuth.Nickname,
		Avatar:         userAuth.Avatar,
		Intro:          userAuth.Intro,
		Website:        userAuth.WebSite,
		IsDisable:      userAuth.IsDisable,
		LoginType:      int32(userAuth.LoginType),
		ArticleLikeSet: articleLikeSet,
		CommentLikeSet: commentLikeSet,
	}, nil
}

func (u *UserAuthServiceImpl) GetUserAuthByUsername(username string) (*pb.UserAuth, error) {
	userAuth, err := model.GetUserAuthByUsername(username)
	if err != nil {
		return nil, err
	}
	return &pb.UserAuth{
		Id:            int32(userAuth.ID),
		UserInfoId:    int32(userAuth.UserInfoID),
		Username:      userAuth.Username,
		LoginType:     int32(userAuth.LoginType),
		CreateTime:    userAuth.CreateTime,
		IpAddr:        userAuth.IPAddr,
		IpSource:      userAuth.IPSource,
		LastLoginTime: userAuth.LastLoginTime,
		RoleId:        int32(userAuth.RoleId),
		NickName:      userAuth.Nickname,
		Avatar:        userAuth.Avatar,
		WebSite:       userAuth.WebSite,
		Intro:         userAuth.Intro,
		IsDisable:     userAuth.IsDisable,
	}, nil

}

func NewUserAuthServiceImpl() *UserAuthServiceImpl {
	return &UserAuthServiceImpl{}
}

func (u *UserAuthServiceImpl) Register(user *pb.User) error {
	//TODO 检测账号是否存在

	err := dao.SqlTransaction(dao.Db.Begin(), func(tx *gorm.DB) error {
		//新增用户信息
		userId, err := model.AddUserInfo(tx, &model.UserInfo{
			Email:      user.GetUsername(),
			Nickname:   fmt.Sprintf("用户%s", common.GetRandomString()),
			Avatar:     common.DefaultAvatar,
			CreateTime: time.Now().Unix(),
		})
		if err != nil {
			return err
		}
		//新增用户账号
		hashPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) //加密处理
		if err != nil {
			return err
		}
		err = model.AddUserAuth(tx, &model.UserAuth{
			UserInfoID: userId,
			Username:   user.GetUsername(),
			Password:   string(hashPwd),
			CreateTime: time.Now().Unix(),
			LoginType:  enum.EMAIL.GetLoginType(),
		})
		if err != nil {
			return err
		}
		//绑定用户角色
		err = model.AddUserRole(tx, &model.UserRole{
			UserID: userId,
			RoleID: enum.User.GetRoleId(),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserAuthServiceImpl) Login(user *pb.User) (bool, error) {
	//从数据库获取用户密码
	userAuth, err := model.GetUserAuthByUsername(user.Username)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userAuth.Password), []byte(user.Password)) //验证（对比）
	if err != nil {
		return false, err
	}
	return true, err
}

func (u *UserAuthServiceImpl) GetLoginCode(username string) error {
	//TODO 检验账号是否合法
	//生成六位随机验证码发送
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	logging.Logger.Debug(code)
	//将验证码通过邮件发送给用户
	mail := email.NewEmail()
	//设置发送方的邮箱
	mail.From = common.Sender
	// 设置接收方的邮箱
	mail.To = []string{username}
	//设置主题
	mail.Subject = common.EmailSubject
	//设置文件发送的内容
	mail.Text = []byte("您的验证码为 " + code + " 有效期15分钟，请不要告诉他人哦！")
	//设置服务器相关的配置
	err := mail.Send(common.EmailServerAddr, smtp.PlainAuth("", common.Sender, common.EmailAuthorizationCode, common.EmailHost))
	if err != nil {
		logging.Logger.Error(err)
		return err
	}
	// 将验证码存入redis，设置过期时间为15分钟
	common.GetRedisUtil().SetEx(common.CodeKey+username, code, common.CodeExpireTime, time.Second)
	return nil
}
