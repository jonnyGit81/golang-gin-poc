package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jonny.marcello/golang-gin-poc/entity"
	"github.com/jonny.marcello/golang-gin-poc/service"
	"github.com/jonny.marcello/golang-gin-poc/validators"
	"net/http"
)

type VideoController interface {
	FindAll() []entity.Video
	//Save(ctx *gin.Context) entity.Video
	Save(ctx *gin.Context) (entity.Video, error)
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}

/*
func (v *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return v.service.Save(video)
}
*/

var validate *validator.Validate

func (c *controller) Save(ctx *gin.Context) (entity.Video, error) {

	validate = validator.New()
	validate.RegisterValidation("my-custom-validate-tag", validators.ValidateCoolTitle)

	var video entity.Video
	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		return video, err
	}

	err = validate.Struct(video)

	if err == nil {
		c.service.Save(video)
	}

	return video, err
}

func New(videoService service.VideoService) VideoController {
	return &controller{
		service: videoService,
	}
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
