package main

import (
	"./translator"
	"./utils"
	"fmt"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"path/filepath"
	"syscall"
)

var Root *sciter.Element
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

	htmlPath, err := filepath.Abs("./frontend/translatorView.html")
	if err != nil {
		log.Fatal(err)
	}

	htmlLoadErr := w.LoadFile(htmlPath)
	if htmlLoadErr != nil {
		fmt.Println("Can not load html in the screen: ", htmlLoadErr.Error())
		return
	}

	Root, rootSelectorErr = w.GetRootElement()
	if rootSelectorErr != nil {
		fmt.Println("Can not select Root element")
		return
	}

	// registering methods
	w.DefineFunction("closeWindow", closeApplication)

	w.SetTitle("GoTranslator")

	SetStatus("READY")
}

func closeApplication(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}

func main() {

	readFileButton, _ := Root.SelectById("read_file_button")

	readFileButton.OnClick(func() {
		SetStatus("READING")
		inputPath := utils.ReadInputPath(Root)
		inputFileContent := utils.ReadFile(inputPath)

		SetStatus("TRANSLATING...")
		outputFileName := utils.MarkdownToHtmlName(inputPath)
		outputFileContent := translator.Translate(inputFileContent, outputFileName)

		SetStatus("SAVING...")
		utils.SaveFile(outputFileName, outputFileContent, shouldOpenExplorer(), shouldOpenBrowser())

		SetStatus("DONE")

	})

	w.Show()
	w.Run()
}

func SetStatus(message string) {
	statusText, _ := Root.SelectById("status_text")
	_ = statusText.SetText(message)
}

func shouldOpenExplorer() bool {
	shouldOpenExplorerCheckbox, _ := Root.SelectById("should_open_explorer_checkbox")
	value, _ := shouldOpenExplorerCheckbox.GetValue()
	return value.Bool()
}

func shouldOpenBrowser() bool {
	shouldOpenBrowserCheckbox, _ := Root.SelectById("should_open_browser_checkbox")
	value, _ := shouldOpenBrowserCheckbox.GetValue()
	return value.Bool()
}
