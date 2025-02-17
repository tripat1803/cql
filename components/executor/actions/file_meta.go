package actions

import (
	"encoding/json"
	"log"
	"os"
	"slices"
)

func GetConfig() []byte {
	bytes, err := os.ReadFile(config)
	if err != nil {
		log.Fatalln("File cannot be read.")
	}

	return bytes
}

func WriteToConfig(bytes []byte) {
	err := os.WriteFile(config, bytes, os.ModePerm)
	if err != nil {
		log.Fatalln("Unable to create config")
	}
}

func ValidateConfig() {
	if _, err := os.Stat(config); err != nil {
		_, createErr := os.Create(config)
		if createErr != nil {
			log.Fatalln("Unable to create config file")
		}

		jsonStructure := new(JsonCategoryStructure)
		jsonStructure.Category = FILES
		jsonStructure.Keys = make([]string, 0)

		bytes, jsonError := json.Marshal(jsonStructure)
		if jsonError != nil {
			log.Fatalln("Cannot encode.")
		}

		err := os.WriteFile(config, bytes, os.ModePerm)
		if err != nil {
			log.Fatalln("Unable to create config")
		}
	}
}

func AddConfigKey(key string) {
	config := new(JsonCategoryStructure)

	bytes := GetConfig()
	json.Unmarshal(bytes, &config)
	config.Keys = append(config.Keys, key)

	newBytes, err := json.Marshal(config)
	if err != nil {
		log.Fatalln("Error encoding value")
	}
	WriteToConfig(newBytes)
}

func DeleteConfigKey(key string) {
	config := new(JsonCategoryStructure)

	bytes := GetConfig()
	json.Unmarshal(bytes, &config)
	index := slices.Index(config.Keys, key)
	if index >= 0 {
		config.Keys = append(config.Keys[:index], config.Keys[index+1:]...)

		newBytes, err := json.Marshal(config)
		if err != nil {
			log.Fatalln("Error encoding value")
		}
		WriteToConfig(newBytes)
	}
}

func GetConfigFileKeys() []string {
	config := new(JsonCategoryStructure)

	bytes := GetConfig()
	json.Unmarshal(bytes, &config)

	return config.Keys
}
