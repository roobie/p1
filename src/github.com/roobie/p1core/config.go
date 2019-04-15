package p1core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Mappings `json:"mappings"`
}

type Mappings struct {
	Authentication MappingItem `json:"authentication"`
	Authorization  MappingItem `json:"authorization"`
	Accounting     MappingItem `json:"accounting"`
}

type MappingItem struct {
	Hostname string `json:"hostname"`
	IPV4     string `json:"ipv4"`
	Port     string `json:"port"`
}

func GetConfig() Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(bytes, &config)
	return config
}
