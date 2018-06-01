package main

import (
	"PixelTool_Desktop/views"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	centralWidget := widgets.NewQWidget(nil, 0)

	sliderInput := views.NewSliderInput("Gamma", 0.24)
	layout := widgets.NewQVBoxLayout2(centralWidget)
	layout.AddWidget(sliderInput.Cell, 0, 0)

	/*
		// initialize input field
		deviceQEInputField := views.NewInputField("DeviceQE", "File Path")
		whitePixelInputField := views.NewInputField("Whipte Pixel", "File Path")

		// add input field to central widget
		layout := widgets.NewQVBoxLayout2(centralWidget)
		layout.AddWidget(deviceQEInputField.Cell, 0, 0)
		layout.AddWidget(whitePixelInputField.Cell, 0, 0)

		centralWidget.SetLayout(layout)
	*/

	centralWidget.Show()
	app.Exec()
}
