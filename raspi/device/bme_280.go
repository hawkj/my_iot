package device

import (
	"context"
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	commonstruct "github.com/hawkj/my_iot/common/struct"
	"math/rand"
	"time"
)

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

func GetBME280Data(ctx context.Context, bmp *bsbmp.BMP) (commonstruct.BME280, error) {
	result := commonstruct.BME280{}
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
	result.Timestamp = time.Now().Unix()
	return result, nil
}

func GetBME280Data4Test(ctx context.Context) (commonstruct.BME280, error) {
	result := commonstruct.BME280{}
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	result.Temperature = float32(randomGenerator.Intn(50))
	result.Pressure = 10.0
	result.Humidity = 10.0
	result.Timestamp = time.Now().Unix()
	return result, nil
}
