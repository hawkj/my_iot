package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hawkj/my_iot/iot_server/config"
	"github.com/hawkj/my_iot/iot_server/internal/context/request_context"
	"github.com/hawkj/my_iot/iot_server/pkg/common"
)

type CommonApiHandler func(c *requestcontext.CommonContext)

func CommonHandlerWrapper(handlerFunc CommonApiHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		siteID := c.GetHeader("Site-Id")
		ctx, _ := c.Get("CommonContext")
		commonContext := ctx.(*requestcontext.CommonContext)
		commonContext.SiteID = siteID
		handlerFunc(commonContext)
	}
}

func CommonContext(g *common.Global, conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		commonContext := &requestcontext.CommonContext{
			GinContext: c,
			Global:     g,
			Config:     conf,
		}
		c.Set("CommonContext", commonContext)
		c.Next()
	}
}
