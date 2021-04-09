package main

import (
	"io/ioutil"
	"log"
)

// This function returns the running version
// of this open inventory instance as a string
func GetRunningVersion() string {

	data, err := ioutil.ReadFile("./VERSION")
	if err != nil {
		log.Fatal(err.Error())
	}

	return string(data)
}
