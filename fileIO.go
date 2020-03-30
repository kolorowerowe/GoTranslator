package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func readFile(path string) string {
	data, err := ioutil.ReadFile(path) /* handle error */
	if err != nil {
		log.Print("Cannot read file: ", err)
		setStatus("ERROR: cannot read file")
	}
	return string(data)

}

func saveFile(name string, content string) {
	message := []byte(content)
	err := ioutil.WriteFile("result/"+name, message, 0644)
	if err != nil {
		log.Print("Cannot save file: ", err)
		setStatus("ERROR: cannot save file")
	}

	if shouldOpenExplorer() {
		path, _ := os.Getwd()
		_ = exec.Command(`explorer`, `/select,`, path+`\result\`+name).Run()
	}

}

func shouldOpenExplorer() bool {
	shouldOpenExplorerCheckbox, _ := root.SelectById("should_open_explorer_checkbox")
	value, _ := shouldOpenExplorerCheckbox.GetValue()
	return value.Bool()
}
