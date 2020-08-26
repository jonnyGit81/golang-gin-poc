package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jonny.marcello/golang-gin-poc/entity"
	"github.com/jonny.marcello/golang-gin-poc/service"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) entity.Video
}

type videoController struct {
	service service.VideoService
}

func (c *videoController) FindAll() []entity.Video {
	return c.service.FindAll()
}

func (v *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return v.service.Save(video)
}

func New(videoService service.VideoService) VideoController {
	return &videoController{
		service: videoService,
	}
}
