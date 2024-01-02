package commonstruct

type BME280 struct {
	Temperature float32
	Pressure    float32
	Humidity    float32
	Timestamp   int64
}

type DeviceUploadMessage struct {
	Topic       string      `json:"topic"`
	MqttMessage MqttMessage `json:"mqtt_message"`
	ClientID    string      `json:"clientid"`
}

type MqttMessage struct {
	MsgType string      `json:"msg_type"`
	Data    interface{} `json:"data"`
}
