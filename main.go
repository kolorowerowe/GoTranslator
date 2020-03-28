package main

import (
	"log"
	"path/filepath"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, nil)
	if err != nil {
		log.Fatal(err)
	}

	htmlPath, err := filepath.Abs("template.html")
	if err != nil {
		log.Fatal(err)
	}

	err = w.LoadFile(htmlPath)
	if err != nil {
		log.Fatal(err)
	}

	w.SetTitle("Example")
	w.Show()
	w.Run()
}
