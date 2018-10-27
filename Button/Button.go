package main

import (
	"fmt"
	
	. "github.com/lxn/walk/declarative"
)

func main() {
	MainWindow {
		Title:	"Button",
		MinSize:	Size { 600, 400 },
		Layout:	VBox {},
		Children:
			[]Widget {
				PushButton {
					Text:	"BTN",
					OnClicked:
						func() {
							fmt.Println("Clicked")
						},
				},
			},
	}.Run()
}
