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
	CommentService service.ICommentService = impl.NewCommentServiceImpl()
)

type CommentRestApi struct {
	base.Handler
}

func NewCommentRestApi() *CommentRestApi {
	return &CommentRestApi{}
}

func (c *CommentRestApi) GetComments(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("articleId"))
	currentPage, _ := strconv.Atoi(ctx.Query("currentPage"))
	commentInfo, err := CommentService.GetComments(id, &page.IPage{Current: currentPage})
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetCommentsFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:       pb.Response_ResponseBeginIndex,
		Code:        pb.ResultCode_SuccessOK,
		ServerTime:  time.Now().Unix(),
		CommentInfo: commentInfo,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *CommentRestApi) AddComment(ctx *gin.Context) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		logging.Logger.Error(err)
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
	var message string
	var errMessage common.ErrorCode
	if request.CsComment.ReplyId != 0 {
		//回复
		message = "回复成功"
		errMessage = common.AddReplyFail
	} else {
		message = "评论成功"
		errMessage = common.AddCommentFail
	}
	err = CommentService.AddComment(request.CsComment)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, errMessage)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    message,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *CommentRestApi) GetReplies(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("commentId"))
	currentPage, _ := strconv.Atoi(ctx.Query("currentPage"))
	logging.Logger.Debug(id)
	logging.Logger.Debug(currentPage)
	replies, err := CommentService.GetReplies(id, &page.IPage{Current: currentPage})
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetRepliesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		ReplyList:  replies,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *CommentRestApi) LikeComment(ctx *gin.Context) {
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
	logging.Logger.Debug(request.LikeComment)
	err = CommentService.LikeComment(request.LikeComment.CommentId, request.LikeComment.UserId)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.LikeCommentFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *CommentRestApi) GetAdminComments(ctx *gin.Context) {
	type Condition struct {
		Current  int64  `json:"current" form:"current"`
		Size     int32  `json:"size" form:"size"`
		Keywords string `json:"keywords" form:"keywords"`
		IsDelete int32  `json:"isDelete" form:"isDelete"`
	}
	var condition Condition
	err := ctx.ShouldBind(&condition)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(condition)

	adminComment, err := CommentService.GetAdminComments(&pb.CsCondition{
		Current:  condition.Current,
		Size:     condition.Size,
		Keywords: condition.Keywords,
		IsDelete: condition.IsDelete,
	})
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.GetCommentsFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:         pb.Response_ResponseBeginIndex,
		Code:          pb.ResultCode_SuccessOK,
		ServerTime:    time.Now().Unix(),
		AdminComments: adminComment,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *CommentRestApi) UpdateCommentStatus(ctx *gin.Context) {
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
	logging.Logger.Debug(request.CommentStatus)
	err = CommentService.UpdateCommentStatus(request.CommentStatus)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.UpdateCommentFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    "更新成功",
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *CommentRestApi) DeleteComment(ctx *gin.Context) {
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
	logging.Logger.Debug(request.CommentIds)
	err = CommentService.DeleteComments(request.CommentIds)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.DeleteCommentFail)
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
