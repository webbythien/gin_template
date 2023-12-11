package route

import (
	"admin-panel/v1/app/controller"
	"github.com/gin-gonic/gin"
	"net/http/pprof"
)

func ApiRoute(e *gin.Engine) {
	e.GET("/posts", controller.GetPost)
	e.GET("/debug/pprof/", func(context *gin.Context) {
		pprof.Index(context.Writer, context.Request)
	})
	e.GET("/debug/pprof/profile", func(context *gin.Context) {
		pprof.Profile(context.Writer, context.Request)
	})
	e.GET("/debug/pprof/trace", func(context *gin.Context) {
		pprof.Trace(context.Writer, context.Request)
	})
}

func ChatRoute(e *gin.Engine) {
	e.POST("/send_chat", controller.SendMessage)
}
