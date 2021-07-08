package common

import (
	conf "blog-go-gin/config"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/rs/xid"
)

var _qiniuUtil *QiNiuUtil

type QiNiuUtil struct{}

func NewQiNiuUtil() *QiNiuUtil {
	return &QiNiuUtil{}
}

func init() {
	_qiniuUtil = NewQiNiuUtil()
}

func GetQiNiuUtil() *QiNiuUtil {
	return _qiniuUtil
}

func GetQiNiuConfig() *conf.QiNiu {
	return &conf.GetConf().QiNiu
}

func (*QiNiuUtil) UploadQiNiu(localFilePath string) (string, error) {
	qiNiuConfig := GetQiNiuConfig()
	accessKey := qiNiuConfig.AccessKey
	secretKey := qiNiuConfig.SecretKey
	bucket := qiNiuConfig.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = setQiNiuArea(qiNiuConfig.Area)
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	key := "blog-go-gin/" + xid.New().String()
	err := formUploader.PutFile(ctx, &ret, upToken, key, localFilePath, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return ret.Key, nil
}

func setQiNiuArea(area string) *storage.Region {
	var zone *storage.Region
	switch area {
	case "huadong":
		zone = &storage.ZoneHuadong
	case "huabei":
		zone = &storage.ZoneHuabei
	case "huanan":
		zone = &storage.ZoneHuanan
	case "beimei":
		zone = &storage.ZoneBeimei
	case "xinjiapo":
		zone = &storage.ZoneXinjiapo
	}
	return zone
}
