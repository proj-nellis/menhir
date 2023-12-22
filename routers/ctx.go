package routers

import (
	"projnellis.com/menhir/app"
)

type RouteContext struct {
	App *app.App
}

// func (r_ctx *RouteContext) SessionMiddleware(ctx *gin.Context) {
// 	middleware.InnerSessionMiddleware(ctx, r_ctx.App)
// }

func (ctx *RouteContext) GenerateSnowflakeId() string {
	return (*ctx.App.Snowflakes).Generate().String()
}
