package blogApi

import (
	"blog-go-gin/common"
	pb "blog-go-gin/go_proto"
	"blog-go-gin/handlers/base"
	"blog-go-gin/logging"
	"blog-go-gin/models/page"
	"blog-go-gin/service"
	"blog-go-gin/service/impl"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var (
	ArticleService service.IArticleService = impl.NewArticleServiceImpl()
)

type ArticleRestApi struct {
	base.Handler
}

func NewArticleRestApi() *ArticleRestApi {
	return &ArticleRestApi{}
}

func (c *ArticleRestApi) GetArticleList(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	//ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(request)
	var ipage page.IPage
	ipage.Current = int(request.CurrentPage)
	articles, err := ArticleService.GetArticleList(ipage)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetArticlesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:       pb.Response_ResponseBeginIndex,
		Code:        pb.ResultCode_SuccessOK,
		ServerTime:  time.Now().Unix(),
		ArticleList: articles,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *ArticleRestApi) GetArticleById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	articleInfo, err := ArticleService.GetArticleById(id)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetArticleByIdFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:       pb.Response_ResponseBeginIndex,
		Code:        pb.ResultCode_SuccessOK,
		ServerTime:  time.Now().Unix(),
		ArticleInfo: articleInfo,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *ArticleRestApi) GetArticleArchives(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logging.Logger.Error(err)
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
	}
	logging.Logger.Debug(request)

	iPage := &page.IPage{
		Current: int(request.CurrentPage),
	}
	archiveInfo, err := ArticleService.GetArchiveList(iPage)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetArticleArchivesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:       pb.Response_ResponseBeginIndex,
		Code:        pb.ResultCode_SuccessOK,
		ServerTime:  time.Now().Unix(),
		ArchiveInfo: archiveInfo,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}
