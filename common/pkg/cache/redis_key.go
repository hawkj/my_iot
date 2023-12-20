package commoncache

import "fmt"

// keys & fields
const H_Device = "H_Device:%s:%s"
const H_F_Device_DeviceData = "device-data"

func GetDeviceCacheKey(siteID, deviceCode string) string {
	return fmt.Sprintf(H_Device, siteID, deviceCode)
}
