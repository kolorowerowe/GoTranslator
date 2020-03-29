package main

import (
	"io/ioutil"
	"log"
)

func readFile(path string) string {
	data, err := ioutil.ReadFile(path) /* handle error */
	if err != nil {
		log.Fatal("Cannot read file: ", err)
	}

	return string(data)

}
