package jobhandler

import (
	"context"
	"encoding/json"
	"github.com/hawkj/my_iot/common/constants"
	"github.com/hawkj/my_iot/common/function"
	"github.com/hawkj/my_iot/common/pkg/queue"
	"github.com/hawkj/my_iot/common/struct"
	"github.com/hawkj/my_iot/raspi/device"
	"github.com/hawkj/my_iot/raspi/pkg/common"
	"log"
	"time"
)

func Bme280(ctx context.Context, g *common.Global) {
	i2c, bmp, err := device.GetBME280(ctx, commoncons.BME280I2CAddress)
	defer i2c.Close()
	if err != nil {
		panic(err)
	}

	mqttClient, err := queue.GetEmqClient(g.Config.SiteInfo.Name, g.Config.Emq.BrokerAddress, g.Config.Emq.Username, g.Config.Emq.Password)
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
		err = queue.SendMessage(mqttClient, commonfunc.GetEmqDeviceUploadTopic(commoncons.DeviceBME280), string(jsonData))
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("SendMessage Done!")
	}
}

func Bme280Mock(ctx context.Context, g *common.Global) {
	mqttClient, err := queue.GetEmqClient(g.Config.SiteInfo.Name, g.Config.Emq.BrokerAddress, g.Config.Emq.Username, g.Config.Emq.Password)
	if err != nil {
		panic(err)
	}
	for {
		data, err := device.GetBME280Data4Test(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		mqttMesage := commonstruct.MqttMessage{}
		mqttMesage.MsgType = commoncons.MqttMsgTypeDeviceData
		mqttMesage.Data = data
		jsonData, err := json.Marshal(mqttMesage)
		if err != nil {
			log.Println(err)
			continue
		}
		err = queue.SendMessage(mqttClient, commonfunc.GetEmqDeviceUploadTopic(commoncons.DeviceBME280), string(jsonData))
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("SendMessage Done!")
		time.Sleep(time.Second * 60)
	}
}
