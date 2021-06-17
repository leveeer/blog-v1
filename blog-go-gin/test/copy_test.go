package test

import (
	"blog-go-gin/common"
	"blog-go-gin/config"
	"blog-go-gin/dao"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/logging"
	"blog-go-gin/models/model"
	"testing"
	"time"
)

func init() {
	config.GetConf()
	dao.InitMysql()
	common.InitRedis()
}

func TestCopyField(t *testing.T) {
	a1 := &pb.Article{
		Id:             30,
		UserId:         1,
		CategoryID:     19,
		ArticleCover:   "ecfbrehbvguherf",
		ArticleTitle:   "测试",
		ArticleContent: "测试内容",
		CreateTime:     time.Now().Unix(),
		UpdateTime:     time.Now().Unix(),
		IsTop:          true,
		IsPublish:      true,
		IsDelete:       false,
		IsOriginal:     true,
		ClickCount:     9999,
		CollectCount:   8888,
		CategoryName:   "测试分类",
		Tags: []*pb.Tag{
			{
				TagName: "测试tag1",
			}, {
				TagName: "测试tag2",
			},
		},
	}
	a2 := &model.Article{}

	err := common.Copy(&a2, a1).Do()
	if err != nil {
		return
	}

	logging.Logger.Info(a2)

}
