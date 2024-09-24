// author: zfy  date: 2024/9/24
package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gofer go")

	w.SetContent(widget.NewLabel("Gofer go go"))
	w.ShowAndRun()

	fmt.Println("close!")
}
