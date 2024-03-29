package main

import (
	"errors"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"time"
)

func addLog(tag string, logString string, level int) {
	// only show formatted logs if debugLevel (app config value with a default value of 2)
	// is greater than level parameter
	if debugLevel >= level {
		fmt.Println(time.Now().Format("[ 2006-01-02 15:04:05 ]") + " [ " + tag + " ] [ " + logString + " ]")
	}
}

func getAppConfig(configKey string) (interface{}, error) {
	// Parse app config json file and return error if exists
	appConfigMap, appConfigMapErr := gabs.ParseJSONFile(appConfigPath)
	if appConfigMapErr != nil {
		return nil, appConfigMapErr
	}

	// Try to find the provided key on the map
	for appConfigKey, appConfigValue := range appConfigMap.ChildrenMap() {
		if configKey == appConfigKey {
			return appConfigValue.Data(), nil
		}
	}
	// if provided key does not exist, return an error
	return nil, errors.New("No " + configKey + " app config key found")
}

func getDeviceID() string {
	// Parse app config json file and return error if exists
	appConfigMap, appConfigMapErr := gabs.ParseJSONFile(barbaraIDPath)
	if appConfigMapErr != nil {
		return ""
	}

	// Try to find the provided key on the map
	for appConfigKey, appConfigValue := range appConfigMap.ChildrenMap() {
		if "id" == appConfigKey {
			return appConfigValue.Data().(string)
		}
	}

	// if provided key does not exist, return an empty string
	return ""
}

func initMessage() {
	addLog(logTag, "=================================", 2)
	addLog(logTag, "App name: "+logTag, 2)
	addLog(logTag, "App version: "+version, 2)
	addLog(logTag, "Device ID: "+getDeviceID(), 2)
	addLog(logTag, "=================================", 2)
}

func interfaceToBool(providedInterface interface{}) (bool, error) {
	if providedInterface != nil {
		switch providedInterface.(type) {
		case bool:
			return providedInterface.(bool), nil
		default:
			return false, errors.New("not a bool interface")
		}
	}
	return false, errors.New("nil interface")
}

func interfaceToInt(providedInterface interface{}) (int, error) {
	if providedInterface != nil {
		switch providedInterface.(type) {
		case float64:
			return int(providedInterface.(float64)), nil
		default:
			return -1, errors.New("not a float64 interface")
		}
	}
	return -1, errors.New("nil interface")
}

func interfaceToString(providedInterface interface{}) (string, error) {
	if providedInterface != nil {
		switch providedInterface.(type) {
		case string:
			return fmt.Sprintf("%v", providedInterface), nil
		default:
			return "", errors.New("not a string interface")
		}
	}
	return "", errors.New("nil interface")
}
