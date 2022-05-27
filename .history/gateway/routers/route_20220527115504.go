package routers

import (
	"net/http"
	"tiktok/base/logger"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RunServer(mode string) {
	// gin set release mode
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// TODO:r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// regist the api router
	registRouter(r)

	pprof.Register(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	zap.L().Fatal(r.Run().Error())
}

// registRouter regist the api router
func registRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	{
		// basic apis
		apiRouter.POST("/user/register/",)
		apiRouter.POST("/user/login/")
		apiRouter.GET("/feed/").Use()
		apiRouter.GET("/user/").Use()
		publish := apiRouter.Group("/publish")
		{
			publish.POST("/action")
			publish.GET("/list/")

		}
		favorite := apiRouter.Group("/favorite")
		{
			favorite.POST("/action/")
			favorite.GET("/list/")
			// 
		}
		comment := apiRouter.Group("/comment")
		{
			comment.POST("/action/")
			comment.GET("/list/")
		}
		relation := apiRouter.Group("/relation")
		{
			relation.POST("/action/")
			relation.GET("/follower/list/")
			relation.GET("/follow/list/")
		}
	}

	// extra apis - II

}