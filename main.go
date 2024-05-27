package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/canvas"
)

var wys *widget.Label = widget.NewLabel("")

type myTheme struct {
	fyne.Theme
}

// var le int
var sings []string
var numbers []float64
var licz string
var wynn float64

// setting a size of a font in a calculator
func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 30 // Default text size
	}
	return m.Theme.Size(name)
}

func main() {
	//background := canvas.NewImageFromFile("/home/wojtyla/programowanie/proj/background.jpg")
	//background.FillMode = canvas.ImageFillStretch

	ap := app.New()
	win := ap.NewWindow("calculator")
	ap.Settings().SetTheme(&myTheme{theme.DefaultTheme()})
	wynn = 0

	je := widget.NewButton("1", func() { number(1) })
	dw := widget.NewButton("2", func() { number(2) })
	sum := widget.NewButton("+", func() { sign("+") })
	diff := widget.NewButton("-", func() { sign("-") })
	tim := widget.NewButton("*", func() { sign("*") })
	tre := widget.NewButton("3", func() { number(3) })
	fo := widget.NewButton("4", func() { number(4) })
	fiv := widget.NewButton("5", func() { number(5) })
	six := widget.NewButton("6", func() { number(6) })
	sev := widget.NewButton("7", func() { number(7) })
	eigh := widget.NewButton("8", func() { number(8) })
	nine := widget.NewButton("9", func() { number(9) })
	zero := widget.NewButton("0", func() { number(0) })
	div := widget.NewButton("/", func() { sign("/") })
	c := widget.NewButton("C", func() {
		numbers = []float64{}
		sings = []string{}
		wys.Text = ""
		wys.Refresh()
	})

	eql := widget.NewButton("=", func() {

		wynn = numbers[0]
		//since when "=" is used "sign" function isnt called so, convert var "le" created in function "number" into int and store it in slice "numbers"
		le, err := strconv.ParseFloat(licz, 64)
		if err != nil {
			le = 0
		}
		numbers = append(numbers, le)
		//check sign used based on data stored in slice "sings"
		for i := 0; i < len(sings); i++ {
			if sings[i] == "+" {
				wynn = wynn + numbers[i+1]
			}
			if sings[i] == "-" {
				wynn = wynn - numbers[i+1]
			}
			if sings[i] == "*" {
				wynn = wynn * numbers[i+1]
			}
			if sings[i] == "/" {
				wynn = wynn / numbers[i+1]
			}
		}
		wys.Text = fmt.Sprintln(wynn)
		wys.Refresh()
		sings = []string{}
		numbers = []float64{}
		licz = ""

	})
	e := container.NewHBox(
		je,
		dw,
		tre,
		sum,
		tim,
	)

	win.SetContent(
		//      container.NewWithoutLayout(
		//        background,
		container.NewVBox(
			wys,
			container.NewHBox(
				e,
			),
			container.NewHBox(
				fo,
				fiv,
				six,
				diff,
				div,
			),
			container.NewHBox(
				sev,
				eigh,
				nine,
				zero,
				c,
				eql,
			),
		),
		//  ),
	)
	win.Resize(fyne.NewSize(600, 600))
	win.ShowAndRun()
}

// writing given numbers on "wys" label and storing numbers as strin in var licz
func number(x int) {
	elo := wys.Text + strconv.Itoa(x)
	wys.Text = elo
	wys.Refresh()
	licz = licz + strconv.Itoa(x)
}
func sign(k string) {
	//converting string of a number created in function "numbers" into int and adding it into slice conting numbers
	le, err := strconv.ParseFloat(licz, 64)
	if err != nil {
		le = 0
	}
	numbers = append(numbers, le)
	//appending slice with string pulled from a button and displaying a sign on a label
	elo := wys.Text + k
	sings = append(sings, k)
	wys.Text = elo
	wys.Refresh()
	licz = ""
}
