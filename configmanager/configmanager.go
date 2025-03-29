package configmanager

import (
	"fmt"
	"os"

	"ApiRouter/apperrors"

	"gopkg.in/yaml.v3"
)

type configmanager struct {
	config map[string]map[string]interface{}
}

var (
	isInitialized bool = false
	instance      configmanager
)

func Initialize() {
	instance = configmanager{}
	instance.config = map[string]map[string]interface{}{}
	isInitialized = true
}

func AddYamlConfigFile(id string, file string) *apperrors.AppError {
	if !isInitialized {
		Initialize()
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return apperrors.NewAppError("E1000", fmt.Errorf("unable to add config file to config manager : %s \n %w", file, err))
	}
	config := map[string]interface{}{}
	err1 := yaml.Unmarshal(data, &config)
	if err1 != nil {
		return apperrors.NewAppError("E1000", fmt.Errorf("unable to parse yaml file : %s \n %w  ", file, err1))
	}
	instance.config[id] = config
	return nil

}

func GetString(id string, config string) (string, *apperrors.AppError) {
	if !isInitialized {
		Initialize()
	}
	data, ok := instance.config[id][config]
	if !ok {
		return "", apperrors.NewAppError("E1000", fmt.Errorf("unable to find value for config tag %s , in config file %s ", config, id))

	}
	datacast, ok := data.(string)
	if !ok {
		return "", apperrors.NewAppError("E1001", fmt.Errorf("not a string value found in config file id : %s , for config tag : %s, found %v instead ", id, config, data))
	}
	return datacast, nil

}

func GetInt(id string, config string) (int, *apperrors.AppError) {
	if !isInitialized {
		Initialize()
	}
	data, ok := instance.config[id][config]
	if !ok {
		return -1, apperrors.NewAppError("E1000", fmt.Errorf("unable to find value for config tag %s , in config file %s ", config, id))
	}
	datacast, ok := data.(int)
	if !ok {
		return -1, apperrors.NewAppError("E1001", fmt.Errorf("not a string value found in config file id : %s , for config tag : %s, found %v instead ", id, config, data))
	}
	return datacast, nil

}
