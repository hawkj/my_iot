module github.com/hawkj/my_iot/iot_server

go 1.21.0

replace github.com/hawkj/my_iot/common v0.0.0 => ../common

require (
	github.com/hawkj/my_iot/common v0.0.0
	github.com/segmentio/kafka-go v0.4.47
	gopkg.in/yaml.v3 v3.0.1
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
)
