package commonfunc

import "fmt"

func GetEmqDeviceUploadTopic(deviceName string) string {
	return fmt.Sprintf("device_upload_%s", deviceName)
}
