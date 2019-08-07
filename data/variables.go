package data

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type MyData struct {
	Name     string `yml: "name"`
	FullName string `yml: "fullname"`
	Body     Body   `yml: "body"`
}

type Body struct {
	Height string `yml: "height"`
	Weight string `yml: "weight"`
}

func GetDataFromFile(filePath string) MyData {
	var myData MyData
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Cannot read file %s", err)
	}
	err = yaml.NewDecoder(bytes.NewReader(configFile)).Decode(&myData)
	if err != nil {
		log.Fatalf("Cannot decode the file %s", err)
	}
	return myData
}
