package blogApi

import (
	"blog-go-gin/handlers/base"
	"github.com/gin-gonic/gin"
)

type TagRestApi struct {
	base.Handler
}

func (r *TagRestApi) GetTagList(ctx *gin.Context) {

}
