package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func ReadFile(path string) string {
	data, err := ioutil.ReadFile(path) /* handle error */
	if err != nil {
		log.Print("Cannot read file: ", err)
		//app.SetStatus("ERROR: cannot read file")
	}
	return string(data)

}

func SaveFile(name string, content string, shouldOpenExplorerCheckbox bool, shouldOpenBrowser bool) {
	message := []byte(content)
	err := ioutil.WriteFile("result/"+name, message, 0644)
	if err != nil {
		log.Print("Cannot save file: ", err)
		//app.SetStatus("ERROR: cannot save file")
	}

	path, _ := os.Getwd()
	fullPath := path + `\result\` + name

	if shouldOpenExplorerCheckbox {
		openExplorer(fullPath)
	}

	if shouldOpenBrowser {
		openBrowser(fullPath)
	}
}

func openExplorer(url string) {
	_ = exec.Command(`explorer`, `/select,`, url).Run()
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
