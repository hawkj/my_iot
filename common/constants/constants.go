package commoncons

const (
	DeviceBME280 = "bme280"
)

const BME280I2CAddress uint8 = 0x77

const JobDefaultParam = "{}"

const DefaultConsumerGroup = "default_group"

// kafka 设备数据的topic
const (
	KafkaTopicUploadDevice = "device-upload"
)

// 设备发送给 kafak 的数据类型
const (
	MqttMsgTypeDeviceData = "device-data"
)
