package bootstrap

import (
	"log"
	"os"
	"strconv"
)

const (
	_defaultMaxCharBodyLogger = 500
	_defaultDebugMode         = false
)

func getDebugMode() bool {
	isDebugMode := _defaultDebugMode
	debugModeString := os.Getenv("DEBUG_MODE")

	if debugModeString != "" {
		debugMode, err := strconv.ParseBool(debugModeString)
		if err != nil {
			log.Fatalln("environment DEBUG_MODE must be bool")
		}
		isDebugMode = debugMode
	}

	return isDebugMode
}

func getMaxCharBodyLogger() int {
	maxCharBody := _defaultMaxCharBodyLogger
	maxCharBodyString := os.Getenv("LOGGER_MAX_CHAR_BODY")

	if maxCharBodyString != "" {
		maxChar, err := strconv.Atoi(maxCharBodyString)
		if err != nil {
			log.Fatalln("environment MAX_CHAR_BODY_LOGGER must be integer")
		}
		maxCharBody = maxChar
	}

	return maxCharBody
}
