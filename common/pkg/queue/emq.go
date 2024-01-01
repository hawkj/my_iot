package queue

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func GetEmqClient(clientID string, brokerAddresses, username, password string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().AddBroker(brokerAddresses).
		SetClientID(clientID).
		SetUsername(username). // 设置用户名
		SetPassword(password)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}

func SendMessage(client mqtt.Client, topic, message string) error {
	// 发布消息
	token := client.Publish(topic, 1, false, message)

	// 等待发布操作完成
	token.Wait()

	// 检查发布是否成功
	if token.Error() != nil {
		return token.Error()
	}
	return nil
}
