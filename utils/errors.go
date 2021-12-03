package utils

import (
	"fmt"

	"github.com/labstack/gommon/log"
)

func CheckError(err error) {
	if err != nil {
		data := fmt.Sprintf("ERROR DETECTED: %s", err.Error())
		log.Error(data)
	} else {
		return
	}
}
