package web

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/logging"
	"blog-go-gin/models/page"
	"blog-go-gin/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	ArticleService = &service.ArticleService{}
)

type ArticleRestApi struct {
	base.Handler
}

func (c *ArticleRestApi) GetArticleList(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	logging.Logger.Debug(string(body))

	if err != nil {
		logging.Logger.Error(err)
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
	}
	logging.Logger.Debug(request.CmdId)
	logging.Logger.Debug(request.CurrentPage)
	var ipage page.IPage
	ipage.Current = int(request.CurrentPage)
	articles, err := ArticleService.GetArticleList(ipage)
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetArticlesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:       pb.Response_ResponseBeginIndex,
		Code:        pb.ResultCode_SuccessOK,
		ServerTime:  time.Now().Unix(),
		ArticleList: articles,
	}
	c.WriteWithProtoBuf(ctx, http.StatusOK, data)
	c.RespSuccess(ctx, http.StatusOK, common.SuccessOK, &articles)
}

func (c *ArticleRestApi) GetArticleById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	article, err := ArticleService.GetArticleById(id)
	if err != nil {
		c.RespFailWithDesc(ctx, http.StatusOK, common.GetArticleByIdFail)
		return
	}
	c.RespSuccess(ctx, http.StatusOK, common.SuccessOK, &article)
}
