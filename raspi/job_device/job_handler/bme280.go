package jobhandler

import (
	"context"
	"encoding/json"
	commoncons "github.com/hawkj/my_iot/common/constants"
	commonfunc "github.com/hawkj/my_iot/common/function"
	commonpkg "github.com/hawkj/my_iot/common/pkg"
	"github.com/hawkj/my_iot/raspi/device"
	"github.com/hawkj/my_iot/raspi/pkg/common"
	"log"
)

func Bme280(ctx context.Context, g *common.Global) {
	address := uint8(0x77)
	i2c, bmp, err := device.GetBME280(ctx, address)
	defer i2c.Close()
	if err != nil {
		panic(err)
	}

	mqttClient, err := commonpkg.GetEmqClient(g.Config.SiteInfo.Name, g.Config.Emq.BrokerAddress)
	if err != nil {
		panic(err)
	}
	for {
		data, err := device.GetBME280Data(ctx, bmp)
		if err != nil {
			log.Println(err)
			continue
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
			continue
		}
		err = commonpkg.SendMessage(mqttClient, commonfunc.GetEmqDeviceUploadTopic(commoncons.DeviceBME280), string(jsonData))
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
