package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port string `json:"port"`
	DbPath string `json:"dbPath"`
}

var Conf Config

func init() {
	confData, err := ioutil.ReadFile("../config.json")
	err = json.Unmarshal(confData, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}
