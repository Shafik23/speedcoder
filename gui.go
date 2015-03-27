package main

import (
	"./fetcher"
	"fmt"
	"gopkg.in/qml.v1"
	//"gopkg.in/qml.v1/gl/2.0"
	"os"
)

var qmlFile = "main.qml"
var keyword string
var language string

func GuiMain(lang string, keyw string) {
	keyword = keyw
	language = lang

	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
}

func run() error {
	engine := qml.NewEngine()

	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(snip *Snippet, obj qml.Object) {
			snip.Object = obj
			snip.Code = fetcher.GetCodeSnippet(keyword, language, 200, 300)
		},
	}})

	component, err := engine.LoadFile(qmlFile)
	if err != nil {
		return err
	}

	win := component.CreateWindow(nil)
	win.Set("x", 560)
	win.Set("y", 320)
	win.Show()
	win.Wait()
	return nil
}

type Snippet struct {
	qml.Object
	Code string
}

func (r *Snippet) Paint(p *qml.Painter) {
	fmt.Println("Paint called!")
}
