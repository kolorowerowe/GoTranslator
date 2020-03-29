package main

import (
	"github.com/sciter-sdk/go-sciter"
	"log"
)

func readInputPath(root *sciter.Element) string {

	in1, errin1 := root.SelectById("input_path")
	if errin1 != nil {
		log.Fatal("failed to bound input 1 ", errin1.Error())
	}

	in1val, errv1 := in1.GetValue()
	if errv1 != nil {
		log.Fatal(errv1.Error())
	}

	return in1val.String()
}
