package configmanager

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type configmanager struct {
	config map[string]map[string]interface{}
}

var (
	isInitialized bool = false
	instance      configmanager
)

func initialize() {
	instance = configmanager{}
	instance.config = map[string]map[string]interface{}{}
	isInitialized = true
}

func addYamlConfigFile(id string, file string) error {
	if !isInitialized {
		initialize()
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("unable to add config file to config manager : %s \n %w", file, err)
	}
	config := map[string]interface{}{}
	err1 := yaml.Unmarshal(data, &config)
	if err1 != nil {
		return fmt.Errorf("unable to parse yaml file : %s \n %w  ", file, err1)
	}
	instance.config[id] = config
	return nil

}

func getString(id string, config string) (string, error) {
	if !isInitialized {
		initialize()
	}
	data, ok := instance.config[id][config]
	if !ok {
		return "", fmt.Errorf("unable to find value for config tag %s , in config file %s ", config, id)
	}
	datacast, ok := data.(string)
	if !ok {
		return "", fmt.Errorf("not a string value , found %v", data)
	}
	return datacast, nil

}

func getInt(id string, config string) (int, error) {
	if !isInitialized {
		initialize()
	}
	data, ok := instance.config[id][config]
	if !ok {
		return -1, fmt.Errorf("unable to find value for config tag %s , in config file %s ", config, id)
	}
	datacast, ok := data.(int)
	if !ok {
		return -1, fmt.Errorf("not a int value , found %v", data)
	}
	return datacast, nil

}
