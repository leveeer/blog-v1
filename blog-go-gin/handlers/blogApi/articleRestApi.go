package blogApi

import (
	"blog-go-gin/common"
	conf "blog-go-gin/config"
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
	"os"
	"path"
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
	currentPage, err := strconv.Atoi(ctx.Query("currentPage"))
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	articles, err := ArticleService.GetArticleList(page.IPage{Current: currentPage})
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
	currentPage, err := strconv.Atoi(ctx.Query("currentPage"))
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	archiveInfo, err := ArticleService.GetArchiveList(&page.IPage{Current: currentPage})
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

func (c *ArticleRestApi) GetArticleOptions(ctx *gin.Context) {
	articleOptions, err := ArticleService.GetArticleOptions()
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetArticleOptionsFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:          pb.Response_ResponseBeginIndex,
		Code:           pb.ResultCode_SuccessOK,
		ServerTime:     time.Now().Unix(),
		ArticleOptions: articleOptions,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *ArticleRestApi) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	_, ok := ctx.GetPostForm("index")
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusBadRequest, common.InvalidRequestParams)
		return
	}
	picExpandedName := path.Ext(file.Filename)
	newFilename := strconv.FormatInt(time.Now().Unix(), 10) + picExpandedName
	savePath := conf.GetConf().QiNiu.ImageSavePath
	_, err = os.Stat(savePath)
	if !os.IsExist(err) {
		if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
			panic(err)
		}
	}
	dst := path.Join(savePath, newFilename)
	_ = ctx.SaveUploadedFile(file, dst)
	key, err := ArticleService.UploadImage(dst)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusBadRequest, common.UploadImageFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:       pb.Response_ResponseBeginIndex,
		Code:        pb.ResultCode_SuccessOK,
		ServerTime:  time.Now().Unix(),
		Message:     "上传成功",
		UploadImage: &pb.ScImage{Key: key},
	}
	if ok {
		c.ProtoBuf(ctx, http.StatusOK, data)
	} else {
		c.RespSuccess(ctx, http.StatusOK, common.SuccessOK, data)
	}
}

func (c *ArticleRestApi) AddArticle(ctx *gin.Context) {
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
	err = ArticleService.AddArticle(request.Article)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.AddArticleFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Message:    "添加成功",
	}
	c.ProtoBuf(ctx, http.StatusOK, data)

}

func (c *ArticleRestApi) GetAdminArticles(ctx *gin.Context) {
	type CsAdminArticle struct {
		Current   int64  `json:"current" form:"current"`
		Size      int32  `json:"size" form:"size"`
		Keywords  string `json:"keywords" form:"keywords"`
		IsDelete  int32  `json:"isDelete" form:"isDelete"`
		IsPublish int32  `json:"isPublish" form:"isPublish"`
	}
	var csAdminArticle CsAdminArticle
	err := ctx.ShouldBind(&csAdminArticle)
	if err != nil {
		logging.Logger.Error(err)
		c.ProtoBufFail(ctx, http.StatusOK, common.InvalidRequestParams)
		return
	}
	logging.Logger.Debug(csAdminArticle)

	adminArticle, err := ArticleService.GetAdminArticle(&pb.CsAdminArticles{
		Current:   csAdminArticle.Current,
		Size:      csAdminArticle.Size,
		Keywords:  csAdminArticle.Keywords,
		IsDelete:  csAdminArticle.IsDelete,
		IsPublish: csAdminArticle.IsPublish,
	})
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetArticlesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:        pb.Response_ResponseBeginIndex,
		Code:         pb.ResultCode_SuccessOK,
		ServerTime:   time.Now().Unix(),
		AdminArticle: adminArticle,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *ArticleRestApi) GetArticleInfoById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	articleInfo, err := ArticleService.GetUpdateArticleInfoById(id)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetArticleByIdFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:             pb.Response_ResponseBeginIndex,
		Code:              pb.ResultCode_SuccessOK,
		ServerTime:        time.Now().Unix(),
		UpdateArticleInfo: articleInfo,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *ArticleRestApi) UpdateArticle(ctx *gin.Context) {
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
	err = ArticleService.UpdateArticle(request.Article)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.UpdateArticleFail)
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

func (c *ArticleRestApi) UpdateArticleStatus(ctx *gin.Context) {
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
	logging.Logger.Debug(request.ArticleStatus)
	err = ArticleService.UpdateArticleStatus(request.ArticleStatus)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.UpdateArticleFail)
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

func (c *ArticleRestApi) DeleteArticles(ctx *gin.Context) {
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
	logging.Logger.Debug(request.ArticleIds)
	err = ArticleService.DeleteArticles(request.ArticleIds)
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.DeleteArticleFail)
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

func (c *ArticleRestApi) UpdateArticleTop(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
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
	logging.Logger.Debug(request.ArticleTop)
	err = ArticleService.UpdateArticleTop(id, int8(request.ArticleTop.IsTop))
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.UpdateArticleFail)
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
