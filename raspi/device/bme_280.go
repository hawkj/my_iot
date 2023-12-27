package device

import (
	"context"
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	commonstruct "github.com/hawkj/my_iot/common/struct"
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
	return result, nil
}
