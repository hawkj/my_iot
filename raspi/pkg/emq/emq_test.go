package pkg

import (
	"fmt"
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
	for i := 0; i < 10000; i++ {
		time.Sleep(time.Second * 1)
		err = SendMessage(mqttClient, testTopic, fmt.Sprintf("test_msg_%d", i))
		if err != nil {
			t.Error(err)
		}
		fmt.Println("send: " + fmt.Sprintf("test_msg_%d", i))
	}

	time.Sleep(time.Minute * 10)
}
