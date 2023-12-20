package commonstruct

type BME280 struct {
	Temperature float32 `json:"temperature"`
	Pressure    float32 `json:"pressure"`
	Humidity    float32 `json:"humidity"`
	Timestamp   int64   `json:"timestamp"`
}

type DeviceUploadMessage struct {
	Topic    string `json:"topic"`
	Payload  string `json:"payload"`
	ClientID string `json:"clientid"`
}

type MqttMessage struct {
	MsgType string      `json:"msg_type"`
	Data    interface{} `json:"data"`
}
