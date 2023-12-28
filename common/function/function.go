package commonfunc

import "fmt"

func GetEmqDeviceUploadTopic(deviceName string) string {
	return fmt.Sprintf("device/upload/%s", deviceName)
}
