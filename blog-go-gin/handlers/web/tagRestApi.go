package web

import (
	"blog-go-gin/handlers/base"
	"github.com/gin-gonic/gin"
)

type TagRestApi struct {
	base.Controller
}

func (r *TagRestApi) GetTagList(ctx *gin.Context) {

}
