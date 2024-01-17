package middleware

import (
	"github.com/swclabs/swipe-api/internal/config"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func GinMiddleware(a *gin.Engine) {
	a.Use(CORSMiddleware())
	store := cookie.NewStore([]byte("secret"))
	a.Use(sessions.Sessions("mysession", store))
}

func Sentry(a *gin.Engine) {
	if config.StageStatus != "dev" {
		a.Use(sentrygin.New(sentrygin.Options{
			Repanic:         true,
			WaitForDelivery: true,
		}))
	}
}
