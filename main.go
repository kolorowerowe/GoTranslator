package main

import (
	"fmt"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"path/filepath"
	"syscall"
)

var root *sciter.Element
var rootSelectorErr error
var w *window.Window
var windowErr error

const WindowWidth = 800
const WindowHeight = 450

func init() {

	rect := sciter.NewRect(50, 100, WindowWidth, WindowHeight)
	w, windowErr = window.New(sciter.SW_TITLEBAR|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_GLASSY,
		rect)

	if windowErr != nil {
		fmt.Println("Can not create new window: ", windowErr)
		return
	}

	htmlPath, err := filepath.Abs("translatorView.html")
	if err != nil {
		log.Fatal(err)
	}

	htmlLoadErr := w.LoadFile(htmlPath)
	if htmlLoadErr != nil {
		fmt.Println("Can not load html in the screen: ", htmlLoadErr.Error())
		return
	}

	root, rootSelectorErr = w.GetRootElement()
	if rootSelectorErr != nil {
		fmt.Println("Can not select root element")
		return
	}

	// registering methods
	w.DefineFunction("closeWindow", closeApplication)

	w.SetTitle("GoTranslator")

	setStatus("READY")
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

func main() {

	readFileButton, _ := root.SelectById("read_file_button")

	readFileButton.OnClick(func() {
		setStatus("READING")
		inputPath := readInputPath(root)
		inputFileContent := readFile(inputPath)

		setStatus("TRANSLATING...")
		outputFileContent := translate(inputFileContent)
		outputFileName := markdownToHtmlName(inputPath)

		setStatus("SAVING...")
		saveFile(outputFileName, outputFileContent)

		setStatus("DONE")

	})

	w.Show()
	w.Run()
}

func setStatus(message string) {
	statusText, _ := root.SelectById("status_text")
	_ = statusText.SetText(message)
}
