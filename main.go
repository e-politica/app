package main

import (
	"log"
	"os/exec"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Local server manager")
	w.Resize(fyne.NewSize(500, 80))

	binaryPathLabel := widget.NewLabel("Path to the binary file of the server")
	binaryPathEntry := widget.NewEntry()

	runServerBtn, err := createRunServerBtn(binaryPathEntry)
	if err != nil {
		log.Println("Error creating run server btn:", err)
		return
	}

	grid := container.New(layout.NewAdaptiveGridLayout(2), binaryPathLabel, binaryPathEntry, runServerBtn)
	w.SetContent(grid)

	w.ShowAndRun()
}

func createRunServerBtn(binaryPathEntry *widget.Entry) (btn *widget.Button, err error) {
	powerIconOff, err := fyne.LoadResourceFromPath("./static/power-off.png")
	if err != nil {
		return
	}

	powerIconOn, err := fyne.LoadResourceFromPath("./static/power-on.png")
	if err != nil {
		return
	}

	btn = widget.NewButtonWithIcon(
		"RUN API",
		powerIconOff,
		func() {
			go func() {
				btn.Icon = powerIconOn
				err := exec.Command("./" + binaryPathEntry.Text).Run()
				if err != nil {
					log.Println("Error executing command:", err)
					btn.Icon = powerIconOff
					return
				}
			}()
		},
	)
	btn.Resize(fyne.NewSize(50, 50))

	return
}
