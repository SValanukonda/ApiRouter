package main

import (
	"ApiRouter/configmanager"
	"fmt"
)

func main() {
	configmanager.Initialize()
	configmanager.AddYamlConfigFile("test", "./test.yaml")
	configval, err := configmanager.GetString("test", "string1")
	if err != nil {
		panic(err)
	}
	fmt.Println(configval)
}
