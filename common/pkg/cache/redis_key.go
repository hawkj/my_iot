package commoncache

import "fmt"

// keys & fields
const H_Device = "H_Device:%s:%s"
const HF_H_Device_UploadData = "upload_data"

func GetDeviceCacheKey(siteID, deviceCode string) string {
	return fmt.Sprintf(H_Device, siteID, deviceCode)
}
