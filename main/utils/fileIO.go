package utils

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func ReadFile(path string) string {
	data, err := ioutil.ReadFile(path) /* handle error */
	if err != nil {
		log.Print("Cannot read file: ", err)
		//app.SetStatus("ERROR: cannot read file")
	}
	return string(data)

}

func SaveFile(name string, content string, shouldOpenExplorerCheckbox bool) {
	message := []byte(content)
	err := ioutil.WriteFile("result/"+name, message, 0644)
	if err != nil {
		log.Print("Cannot save file: ", err)
		//app.SetStatus("ERROR: cannot save file")
	}

	if shouldOpenExplorerCheckbox {
		path, _ := os.Getwd()
		_ = exec.Command(`explorer`, `/select,`, path+`\result\`+name).Run()
	}

}
