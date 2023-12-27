package common

import "github.com/hawkj/my_iot/raspi/config"

type Global struct {
	Config *config.Config
}

type BME280 struct {
	Temperature float32
	Pressure    float32
	Humidity    float32
}
