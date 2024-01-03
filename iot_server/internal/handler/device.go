package handler

import (
	"fmt"
	"github.com/hawkj/my_iot/common/api"
	commonerr "github.com/hawkj/my_iot/common/error"
	"github.com/hawkj/my_iot/iot_server/internal/context/request_context"
	"github.com/hawkj/my_iot/iot_server/internal/service"
)

func DeviceInfo(c *requestcontext.CommonContext) {
	siteID := c.SiteID
	deviceCode := c.GinContext.Query("device_code")
	fmt.Println("--------------------")
	fmt.Println(siteID)
	fmt.Println(deviceCode)
	fmt.Println("--------------------")
	if siteID == "" || deviceCode == "" {
		commonapi.ApiError(c.GinContext, commonerr.ErrParams.ErrorMsg)
		return
	}
	deviceInfo, err := service.GetDeviceData(c.GinContext.Request.Context(), c.Global.Redis, deviceCode, siteID)
	if err != nil {
		commonapi.ApiError(c.GinContext, err)
		return
	}
	commonapi.ApiOk(c.GinContext, deviceInfo)
}
