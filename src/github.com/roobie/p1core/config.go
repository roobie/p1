package p1core;

import (
  "log"
  "os"
  "io/ioutil"
  "encoding/json"
)

type Config struct {
  Mappings []Mapping `json:"mappings"`
}

type Mapping struct {
  Hostname string `json:"hostname"`
  IPV4 string `json:"ipv4"`
  Port string `json:"port"`
}

func GetConfig() Config{
  jsonFile, err := os.Open("../config.json")
  if err != nil {
    log.Fatal(err)
  }
  defer jsonFile.Close()
  bytes, _ := ioutil.ReadAll(jsonFile)
  var config Config
  json.Unmarshal(bytes, &config)
  return config
}
