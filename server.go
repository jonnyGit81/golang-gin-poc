package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jonny.marcello/golang-gin-poc/controller"
	"github.com/jonny.marcello/golang-gin-poc/middlewares"
	"github.com/jonny.marcello/golang-gin-poc/service"
	"github.com/tpkeeper/gin-dump"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogFile() {
	f, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func setupHtmlTemplates(server *gin.Engine) {
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")
}

func main() {

	setupLogFile()

	//server := gin.Default()

	server := gin.New()

	setupHtmlTemplates(server)

	//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	/*
		server.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		server.POST("/posts", func(ctx *gin.Context) {
			//ctx.JSON(200, videoController.Save(ctx))
			video, err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
			}else {
				ctx.JSON(http.StatusOK, video)
			}
		})
	*/

	// endpoint implement basic auth midleware, aunthentication is required
	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoutes.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/posts", func(ctx *gin.Context) {
			//ctx.JSON(200, videoController.Save(ctx))
			video, err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
			} else {
				ctx.JSON(http.StatusOK, video)
			}
		})
	}

	// public endpoint, no authorization required
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

	server.Run(":" + port)
}
