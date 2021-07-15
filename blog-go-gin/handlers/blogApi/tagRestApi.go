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

type TagRestApi struct {
	base.Handler
}

func NewTagRestApi() *TagRestApi {
	return &TagRestApi{}
}

var (
	TagService service.ITagService = impl.NewTagServiceImpl()
)

func (c *TagRestApi) GetTags(ctx *gin.Context) {
	tags, err := TagService.GetTags()
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetTagsFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Tags:       tags,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *TagRestApi) GetArticlesByTagID(ctx *gin.Context) {
	tagId, _ := strconv.Atoi(ctx.Param("tagId"))
	currentPage, _ := strconv.Atoi(ctx.Query("currentPage"))
	articles, err := ArticleService.GetArticleByTagID(tagId, &page.IPage{Current: currentPage})
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetArticleByCategoryIDFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:                   pb.Response_ResponseBeginIndex,
		Code:                    pb.ResultCode_SuccessOK,
		ServerTime:              time.Now().Unix(),
		ArticlesByCategoryOrTag: articles,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *TagRestApi) GetAdminTags(ctx *gin.Context) {
	type CsCondition struct {
		Current  int64  `json:"current" form:"current"`
		Size     int32  `json:"size" form:"size"`
		Keywords string `json:"keywords" form:"keywords"`
	}
	var condition CsCondition
	err := ctx.ShouldBind(&condition)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(condition)

	adminTags, err := TagService.GetAdminTags(&pb.CsCondition{
		Current:  condition.Current,
		Size:     condition.Size,
		Keywords: condition.Keywords,
	})
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetTagsFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		AdminTags:  adminTags,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *TagRestApi) AddOrUpdateTag(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(request)
	err = TagService.AddOrUpdateTag(request.CsTag)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.AddOrUpdateTagFail)
		return
	}
	message := "更新成功"
	if request.CsTag.Id == 0 {
		message = "新增成功"
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    message,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *TagRestApi) DeleteTag(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	request := &pb.RequestPkg{}
	err = proto.Unmarshal(body, request)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(request.TagIds)
	err = TagService.DeleteTag(request.TagIds)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.DeleteTagFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    "删除成功",
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}
