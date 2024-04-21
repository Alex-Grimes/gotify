package gotify

import (
	"os/user"
	"strings"
)

func expandUser(path string) string {
	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			return "error"
		}
		return strings.Replace(path, "~", usr.HomeDir, 1)
	}
	return path
}

func GetConfig() {
	userConfigDir := expandUser("~")
}
