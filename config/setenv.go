package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func SetEnv(configFilePath string) Config {
	file, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal("Could not read config file: ", err)
	}
	conf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	if err := json.Unmarshal(conf, &config); err != nil {
		log.Fatal(err)
	}
	if len(config.MongoURI) == 0 {
		log.Fatal("Configuration could not be read")
	}
	fmt.Println("Configuration found ")
	return config
	//	fmt.Println("Config: ", config)
	//	if err := os.Setenv("mongo_username", config.MongoUsername); err != nil {
	//		log.Fatal(err)
	//	}
	//	if err := os.Setenv("mongo_password", config.MongoPassword); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if err := os.Setenv("mongo_database", config.MongoDatabase); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if err := os.Setenv("mongo_collection", config.MongoCollection); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	if err := os.Setenv("js_file_path", config.JSFilePath); err != nil {
	//		log.Fatal(err)
	//	}

}
