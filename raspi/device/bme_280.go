package device

import (
	"context"
	"fmt"
	"github.com/hawkj/my_iot/raspi/pkg/common"
	"log"
	"time"

	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
)

func Bme280() {
	// 创建一个 Channel 用于传输传感器数据
	dataChannel := make(chan common.BME280)

	// 启动传感器数据读取服务
	go Bme280Service(dataChannel, 0x77, time.Second*5)

	// 主程序循环读取传感器数据
	for {
		// 从 Channel 中读取传感器数据
		data := <-dataChannel

		// 处理传感器数据
		fmt.Printf("Temperature: %.2f°C, Pressure: %.2f Pa, Humidity: %.2f%%\n",
			data.Temperature, data.Pressure, data.Humidity)

		// 等待一段时间再次读取
		time.Sleep(5 * time.Second)
	}
}

// Bme280Service 传感器数据读取服务

func Bme280Service(dataChannel chan<- common.BME280, address uint8, waitTime time.Duration) {
	i2c, err := i2c.NewI2C(address, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	bmp, err := bsbmp.NewBMP(bsbmp.BME280, i2c)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 读取传感器数据
		temperature, err := bmp.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
		if err != nil {
			log.Println("Error reading temperature:", err)
			continue
		}

		pressure, err := bmp.ReadPressurePa(bsbmp.ACCURACY_STANDARD)
		if err != nil {
			log.Println("Error reading pressure:", err)
			continue
		}

		supported, humidity, err := bmp.ReadHumidityRH(bsbmp.ACCURACY_STANDARD)
		if err != nil {
			log.Println("Error reading humidity:", err)
			continue
		}

		if supported == false {
			humidity = 0.0
		}
		// 将传感器数据发送到 Channel
		dataChannel <- common.BME280{
			Temperature: temperature,
			Pressure:    pressure,
			Humidity:    humidity,
		}
		// 等待一段时间再次读取
		time.Sleep(waitTime)
	}
}

func GetBME280(ctx context.Context, address uint8) (*i2c.I2C, *bsbmp.BMP, error) {
	i2c, err := i2c.NewI2C(address, 1)
	if err != nil {
		return nil, nil, err
	}

	bmp, err := bsbmp.NewBMP(bsbmp.BME280, i2c)
	if err != nil {
		return nil, nil, err
	}
	return i2c, bmp, nil
}

func GetBME280Data(ctx context.Context, bmp *bsbmp.BMP) (common.BME280, error) {
	result := common.BME280{}
	temperature, err := bmp.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		return result, err
	}

	pressure, err := bmp.ReadPressurePa(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		return result, err
	}

	supported, humidity, err := bmp.ReadHumidityRH(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		return result, err
	}

	if supported == false {
		humidity = 0.0
	}
	result.Temperature = temperature
	result.Pressure = pressure
	result.Humidity = humidity
	return result, nil
}
