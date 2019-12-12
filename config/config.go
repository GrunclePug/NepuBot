package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	Token, Prefix string
	config *configStruct
)

type configStruct struct {
	Token string `json:"Token"`
	Prefix string `json:"Prefix"`
}

func ReadConfig() error {
	fmt.Print("Initializing...")

	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Printf("an error occurred while reading the config, %s\n", err)
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Printf("an error occurred while unmarshalling json data, %s\n", err)
		return err
	}

	Token = config.Token
	Prefix = config.Prefix

	return nil
}