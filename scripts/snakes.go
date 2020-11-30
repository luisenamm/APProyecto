package main

import (
	"errors"
	"image/color"
	"log"
	"time"
	"unicode"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

var cherrys = widget.NewEntry()
var enemies = widget.NewEntry()
var content = widget.NewForm(widget.NewFormItem("Cherry's", cherrys), widget.NewFormItem("Enemies", enemies))

func main() {
	a := app.New()
	w := a.NewWindow("SNAKES NIC")
	validate(w)
	w.Resize(fyne.NewSize(340, 240)) //X Y
	w.SetFixedSize(true)
	w.SetContent(canvas.NewRasterWithPixels(rgbGradient))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func new() {
	g := fyne.CurrentApp().NewWindow("Hello")
	g.SetContent(widget.NewLabel("Edu!"))
	g.Resize(fyne.NewSize(340, 240))
	g.Show()
	g.CenterOnScreen()
}

func validate(w fyne.Window) {
	dialog.ShowCustomConfirm("SetUp", "Start", "Cancel", content, func(b bool) {
		if b {
			log.Println("Info", cherrys.Text, enemies.Text)
			if cherrys.Text == "" || enemies.Text == "" {
				validate(w)
				dialog.ShowError(errors.New("You are missing data"), w)
			} else {
				if !isInt(cherrys.Text) || !isInt(enemies.Text) {
					validate(w)
					dialog.ShowError(errors.New("Insert a number"), w)
				} else {
					prog := dialog.NewProgress("Progress", "Nearly there...", w)
					prog.Show()
					num := 0.0
					for num < 1.0 {
						time.Sleep(50 * time.Millisecond)
						prog.SetValue(num)
						num += 0.01
					}
					prog.Hide()
					prog.SetValue(1)
					new()
				}
			}
		}
	}, w)
}

func rgbGradient(x, y, w, h int) color.Color {
	g := int(float32(x) / float32(w) * float32(255))
	b := int(float32(y) / float32(h) * float32(255))

	return color.NRGBA{uint8(255 - b), uint8(g), uint8(b), 0xff}
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
