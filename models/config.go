package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
)

type Config struct {
	Port       string `json:"port"`
	DbPath     string `json:"dbPath"`
	SigningKey string `json:"signingKey"`
}

var Conf Config

func init() {
	confData, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(confData, &Conf)
	if err != nil {
		fmt.Println("error while decoding config file")
		log.Fatal(err)
	}
}
