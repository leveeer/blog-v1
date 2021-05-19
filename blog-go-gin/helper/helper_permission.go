package helper

import (
	"blog-go-gin/logging"
	"encoding/json"
	"errors"
	"fmt"
)

// req url-path for admin
const (
	ReqPathGroupForAdmin = "/api/admin"
	// 管理员登录or登出(特殊处理, 不需要对应编码)
	ReqPathAdminLogin        = ReqPathGroupForAdmin + "/login"
	ReqPathAddAdmin          = ReqPathGroupForAdmin + "/add_admin"
	ReqPathModifyAdmin       = ReqPathGroupForAdmin + "/modify_admin"
	ReqPathModifyAdminOwnPwd = ReqPathGroupForAdmin + "/modify_admin_own_pwd"
)

// req url-path for user
const (
	ReqPathGroupForUser = "/api/user"

	ReqPathUserLogin     = ReqPathGroupForUser + "/login"      // 用户登录or登出(特殊处理, 不需要对应编码)
	ReqPathUserResetPwd  = ReqPathGroupForUser + "/reset_pw"   // 重置密码
	ReqPathUserBindPhone = ReqPathGroupForUser + "/bind_phone" // 绑定手机
)

func LoginHideHandler(body *[]byte) (err error) {
	type postPara struct {
		LoginType string `json:"login_type"`
		LoginID   string `json:"login_id"`
		LoginPW   string `json:"login_pw"`
	}
	var para postPara
	if err = json.Unmarshal(*body, &para); err != nil {
		return errors.New(fmt.Sprintf("json unmarshal fail. err: %v", err))
	}
	if len(para.LoginID) > 0 {
		para.LoginID = "******"
	}
	if len(para.LoginPW) > 0 {
		para.LoginPW = "******"
	}
	*body, err = json.Marshal(para)
	if err != nil {
		return errors.New(fmt.Sprintf("json marshal failed. err: %v", err))
	}
	return nil
}

var HideSensitiveFuncHandler = map[string]func(*[]byte) error{
	// 管理员登录
	ReqPathAdminLogin: LoginHideHandler,
	// 添加管理员
	ReqPathAddAdmin: func(body *[]byte) (err error) {
		type postPara struct {
			UserName          string `json:"user_name"`
			NickName          string `json:"nick_name"`
			PhoneNumber       string `json:"phone_number"`
			Pwd               string `json:"password"`
			PermissionGroupID uint   `json:"permission_group_id"`
		}
		var para postPara
		if err := json.Unmarshal(*body, &para); err != nil {
			return errors.New(fmt.Sprintf("json unmarshal fail. err: %v", err))
		}
		if len(para.PhoneNumber) > 0 {
			para.PhoneNumber = "******"
		}
		if len(para.Pwd) > 0 {
			para.Pwd = "******"
		}
		*body, err = json.Marshal(para)
		if err != nil {
			return errors.New(fmt.Sprintf("json marshal failed. err: %v", err))
		}
		return nil
	},
	// 修改管理员信息
	ReqPathModifyAdmin: func(body *[]byte) (err error) {
		type postPara struct {
			UserID            uint   `json:"user_id"`
			NickName          string `json:"nick_name"`
			PhoneNumber       string `json:"phone_number"`
			Pwd               string `json:"password"`
			PermissionGroupID uint   `json:"permission_group_id"`
		}
		var para postPara
		if err := json.Unmarshal(*body, &para); err != nil {
			return errors.New(fmt.Sprintf("json unmarshal fail. err: %v", err))
		}
		if len(para.PhoneNumber) > 0 {
			para.PhoneNumber = "******"
		}
		if len(para.Pwd) > 0 {
			para.Pwd = "******"
		}
		*body, err = json.Marshal(para)
		if err != nil {
			return errors.New(fmt.Sprintf("json marshal failed. err: %v", err))
		}
		return nil
	},
	// 修改管理员归属密码
	ReqPathModifyAdminOwnPwd: func(body *[]byte) (err error) {
		type postPara struct {
			OldPwd string `json:"old_pwd"`
			NewPwd string `json:"new_pwd"`
		}
		var para postPara
		if err := json.Unmarshal(*body, &para); err != nil {
			return errors.New(fmt.Sprintf("json unmarshal fail. err: %v", err))
		}
		if len(para.OldPwd) > 0 {
			para.OldPwd = "******"
		}
		if len(para.NewPwd) > 0 {
			para.NewPwd = "******"
		}
		*body, err = json.Marshal(para)
		if err != nil {
			return errors.New(fmt.Sprintf("json marshal failed. err: %v", err))
		}
		return nil
	},
	// 下面是用户级别的API过滤, 也需要隐藏关键信息
	// 玩家登录
	ReqPathUserLogin: LoginHideHandler,
	// 重置密码
	ReqPathUserResetPwd: func(body *[]byte) (err error) {
		type postPara struct {
			PhoneNum string `json:"phone_num"`
			Password string `json:"password"`
			SMSCode  string `json:"sms_code"`
		}
		var para postPara
		if err = json.Unmarshal(*body, &para); err != nil {
			return errors.New(fmt.Sprintf("json unmarshal fail. err: %v", err))
		}
		if len(para.PhoneNum) > 0 {
			para.PhoneNum = "******"
		}
		if len(para.Password) > 0 {
			para.Password = "******"
		}
		*body, err = json.Marshal(para)
		if err != nil {
			return errors.New(fmt.Sprintf("json marshal failed. err: %v", err))
		}
		return nil
	},
	// 绑定手机
	ReqPathUserBindPhone: func(body *[]byte) (err error) {
		type postPara struct {
			PhoneNum string `json:"phone_num"`
			Password string `json:"password"`
			SMSCode  string `json:"sms_code"`
		}
		var para postPara
		if err = json.Unmarshal(*body, &para); err != nil {
			return errors.New(fmt.Sprintf("json unmarshal fail. err: %v", err))
		}
		if len(para.PhoneNum) > 0 {
			para.PhoneNum = "******"
		}
		if len(para.Password) > 0 {
			para.Password = "******"
		}
		*body, err = json.Marshal(para)
		if err != nil {
			return errors.New(fmt.Sprintf("json marshal failed. err: %v", err))
		}
		return nil
	},
}

func HideSensitiveInfo(body *[]byte, path string) (err error) {
	if body != nil && len(*body) != 0 {
		if f, need := HideSensitiveFuncHandler[path]; need {
			if err := f(body); err != nil {
				logging.Logger.Debugf("path: %s, err: %v", path, err)
				*body = []byte("******")
			}
		}
		return nil
	}
	return errors.New("body is empty")
}
