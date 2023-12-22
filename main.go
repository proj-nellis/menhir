package main

// helloß
import (
	"fmt"

	"github.com/gin-gonic/gin"
	"projnellis.com/menhir/app"
	"projnellis.com/menhir/routers"
)

func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-Password")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "PATCH, POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func main() {
	app := app.Init()
	if app.Config.Mode != nil && *app.Config.Mode == "prod" || *app.Config.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(CORSMiddleware)
	r.MaxMultipartMemory = 10 << 50
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s (%s) %d --> %s\n",
			param.Method,
			param.ClientIP,
			param.StatusCode,
			param.Path,
		)
	}))
	r_ctx := &routers.RouteContext{App: &app}
	// UNAUTHORIZED ROUTES
	r.POST("/accounts", r_ctx.PostCreateAccount)
	// AUTHORIZED ROUTES
	r.Run()
}
