package pkg

import (
	"github.com/hawkj/my_iot/raspi/config"
	"os"
	"testing"
	"time"
)

func Test_GetEmqClient(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	configFile := currentDir + "/../../config/raspi_conf.yaml"
	c := config.GetConfig(configFile)
	testTopic := "test"

	mqttClient, err := GetEmqClient(c.SiteInfo.Name, c.Emq.BrokerAddress)
	if err != nil {
		t.Error(err)
	}
	// 在测试开始时订阅主题
	if token := mqttClient.Subscribe(testTopic, 0, nil); token.Wait() && token.Error() != nil {
		t.Error(token.Error())
	}

	err = SendMessage(mqttClient, testTopic, "test_message")
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Minute * 10)
}
