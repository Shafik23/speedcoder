package main

import (
	"./fetcher"
	"fmt"
	"gopkg.in/qml.v1"
	//"gopkg.in/qml.v1/gl/2.0"
	"os"
)

var filename = "main.qml"

func GuiMain() {
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	qml.RegisterTypes("GoExtensions", 1, 0, []qml.TypeSpec{{
		Init: func(sm *Snippet, obj qml.Object) {
			sm.Object = obj

			sm.Code = fetcher.GetCodeSnippet("net", "java", 200, 300)
		},
	}})

	component, err := engine.LoadFile(filename)
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
