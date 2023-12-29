package commonstruct

type BME280 struct {
	Temperature float32
	Pressure    float32
	Humidity    float32
	Timestamp   int64
}
