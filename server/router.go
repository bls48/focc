package server

import (
	"focc/api"
	"focc/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 跨域, 登录状态
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("videos", api.CreatVideo)
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("list", api.ListVideo)
		v1.PUT("video/:id", api.UpdateVideo)
		v1.GET("videodel/:id", api.DeleteVideo)
		v1.POST("upload/token", api.UploadToken)
		v1.POST("ping", api.Ping)
		v1.GET("rank/daily", api.DailyRank)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe) //个人资料
			auth.DELETE("user/logout", api.UserLogout)
		}

	}
	return r
}
