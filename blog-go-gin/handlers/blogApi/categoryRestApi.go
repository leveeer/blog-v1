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
	"net/http"
	"strconv"
	"time"
)

var (
	CategoryService service.ICategoryService = impl.NewCategoryServiceImpl()
)

type CategoryRestApi struct {
	base.Handler
}

func NewCategoryRestApi() *CategoryRestApi {
	return &CategoryRestApi{}
}

func (c *CategoryRestApi) GetCategories(ctx *gin.Context) {
	categories, err := CategoryService.GetCategories()
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetCategoriesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:      pb.Response_ResponseBeginIndex,
		Code:       pb.ResultCode_SuccessOK,
		ServerTime: time.Now().Unix(),
		Categories: categories,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}

func (c *CategoryRestApi) GetArticleByCategoryID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("categoryId"))
	currentPage, _ := strconv.Atoi(ctx.Query("currentPage"))
	articles, err := ArticleService.GetArticleByCategoryID(id, &page.IPage{Current: currentPage})
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

func (c *CategoryRestApi) GetAdminCategories(ctx *gin.Context) {
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

	adminCategories, err := CategoryService.GetAdminCategory(&pb.CsCondition{
		Current:  condition.Current,
		Size:     condition.Size,
		Keywords: condition.Keywords,
	})
	if err != nil {
		c.ProtoBufFail(ctx, http.StatusOK, common.GetArticlesFail)
		return
	}
	data := &pb.ResponsePkg{
		CmdId:           pb.Response_ResponseBeginIndex,
		Code:            pb.ResultCode_SuccessOK,
		ServerTime:      time.Now().Unix(),
		AdminCategories: adminCategories,
	}
	c.ProtoBuf(ctx, http.StatusOK, data)
}
