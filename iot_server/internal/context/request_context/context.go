package requestcontext

import (
	"github.com/gin-gonic/gin"
	"github.com/hawkj/my_iot/iot_server/config"
	"github.com/hawkj/my_iot/iot_server/pkg/common"
)

type CommonContext struct {
	GinContext *gin.Context
	Global     *common.Global
	SiteID     string
	Config     *config.Config
}
