package main

import (
	"fmt"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"path/filepath"
)

var root *sciter.Element
var rootSelectorErr error
var w *window.Window
var windowErr error

func init() {

	rect := sciter.NewRect(0, 100, 300, 350)
	w, windowErr = window.New(sciter.SW_TITLEBAR|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_GLASSY,
		rect)

	if windowErr != nil {
		fmt.Println("Can not create new window: ", windowErr)
		return
	}

	htmlPath, err := filepath.Abs("template.html")
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
	w.SetTitle("GoTranslator")
}

func main() {

	addbutton, _ := root.SelectById("add")

	out1, errout1 := root.SelectById("output1")
	if errout1 != nil {
		log.Fatal("failed to bound output 1 ", errout1.Error())
	}

	addbutton.OnClick(func() {
		input1 := readInput1(root)
		output := processValue(input1)
		_ = out1.SetText(fmt.Sprint(output))
	})

	w.Show()
	w.Run()
}
