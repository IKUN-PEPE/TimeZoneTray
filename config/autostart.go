package config

import (
	"os"

	"golang.org/x/sys/windows/registry"
)

const registryKey = `Software\Microsoft\Windows\CurrentVersion\Run`
const appName = "ZoneTray"

// SetAutostart 设置或取消开机自启动
func SetAutostart(enabled bool) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, registryKey, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	if enabled {
		execPath, err := os.Executable()
		if err != nil {
			return err
		}
		return k.SetStringValue(appName, execPath)
	} else {
		return k.DeleteValue(appName)
	}
}
