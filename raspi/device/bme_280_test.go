package device

import (
	"github.com/d2r2/go-bsbmp"
	"github.com/d2r2/go-i2c"
	"github.com/hawkj/my_iot/common/constants"
	"log"
	"testing"
	"time"
)

// go test -run Test_Bme280
func Test_Bme280(t *testing.T) {
	// 创建I2C连接
	i2c, err := i2c.NewI2C(commoncons.BME280I2CAddress, 1) // 0x77是BME280的I2C地址，1是I2C总线号
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	// 创建BME280实例
	bmp, err := bsbmp.NewBMP(bsbmp.BME280, i2c)
	if err != nil {
		log.Fatal(err)
	}

	// 读取BME280的数据
	temperature, err := bmp.ReadTemperatureC(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	pressure, err := bmp.ReadPressurePa(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	supported, humidity, err := bmp.ReadHumidityRH(bsbmp.ACCURACY_STANDARD)
	if err != nil {
		log.Fatal(err)
	}

	// 打印读取到的数据
	if supported {
		log.Printf("Temperature: %.2f°C\n", temperature)
		log.Printf("Pressure: %.2f Pa\n", pressure)
		log.Printf("Humidity: %.2f%%\n", humidity)
	} else {
		log.Println("Humidity reading not supported by the sensor.")
	}

	// 注意：BME280读取数据可能需要一些时间，请根据传感器的规格表设置适当的延时

	// 为了让程序运行一段时间，可以添加一个sleep
	time.Sleep(1 * time.Second)
}
