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

func saveFile(name string, content string) {
	message := []byte(content)
	err := ioutil.WriteFile("result/"+name, message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
