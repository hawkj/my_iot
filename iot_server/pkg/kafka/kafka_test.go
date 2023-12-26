package pkg

import (
	"fmt"
	"github.com/hawkj/my_iot/iot_server/config"
	"os"
	"testing"
	"time"
)

func Test_kafka(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	configFile := currentDir + "/../../config/iot_server_conf.yaml"
	c := config.GetConfig(configFile)
	fmt.Println(c.Kafka)
	//testTopic := "test"

	time.Sleep(time.Minute * 10)
}
