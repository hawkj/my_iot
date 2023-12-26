package pkg

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GetEmqClient(clientID string, brokerAddresses string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(brokerAddresses).SetClientID(clientID)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}

func SendMessage(client mqtt.Client, topic, message string) error {
	// 发布消息
	token := client.Publish(topic, 0, false, message)

	// 等待发布操作完成
	token.Wait()

	// 检查发布是否成功
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}

//func main() {
//	// 发送消息到多个代理
//	message := "Hello, EMQ!" // 替换成你要发布的消息
//	sendMessage(clientID, message, brokerAddresses)
//
//	// 等待一段时间以接收订阅消息
//	time.Sleep(5 * time.Second)
//
//	// 等待程序终止信号
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt)
//	<-c
//}
