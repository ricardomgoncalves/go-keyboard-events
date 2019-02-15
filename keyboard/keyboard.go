package keyboard

import (
	"io/ioutil"
	"strconv"
	"strings"
)
func FindKeyboardDevices() []string {
	devicePath := "/dev/input/event"
	prePath := "/sys/class/input/event"
	posPath := "/device/name"
	var devices []string

	for i := 0; i < 255; i++ {
		buff, err := ioutil.ReadFile(prePath + strconv.Itoa(i) + posPath)
		if err != nil {
			continue
		}
		if strings.Contains(strings.ToLower(string(buff)), "keyboard") {
			devices = append(devices, devicePath+strconv.Itoa(i))
		}
	}
	return devices
}

func FindKeyboardDevice() string {
	devicePath := "/dev/input/event"
	prePath := "/sys/class/input/event"
	posPath := "/device/name"

	for i := 0; i < 255; i++ {
		buff, err := ioutil.ReadFile(prePath + strconv.Itoa(i) + posPath)
		if err != nil {
			continue
		}
		if strings.Contains(strings.ToLower(string(buff)), "keyboard") {
			return devicePath + strconv.Itoa(i)
		}
	}
	return ""
}
