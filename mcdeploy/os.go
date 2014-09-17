package mcdeploy

import (
	"runtime"
)

func GetOS() string {

	switch runtime.GOOS {

	case "linux":

		return "linux"

	case "darwin":

		return "mac"

	case "windows":

		return "windows"

	default:

		return "os not found"

	}

	return "OS not found"

}
