package main

import (
	"PixelToolWindow/views"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	centralWidget := widgets.NewQWidget(nil, 0)

	pixelInput := views.NewPixelSizeInputField("Patch size", 100, 100)
	layout := widgets.NewQVBoxLayout2(centralWidget)
	layout.AddWidget(pixelInput.Cell, 0, 0)

	/*
		comboBox := views.NewComboBoxSelector("Light Source", []string{"D65", "D50", "ill-A"})
		layout := widgets.NewQVBoxLayout2(centralWidget)
		layout.AddWidget(comboBox.Cell, 0, 0)
	*/

	/*
		fileOutputField := views.NewFileOutputField("File Save Path")
		layout := widgets.NewQVBoxLayout2(centralWidget)
		layout.AddWidget(fileOutputField.Cell, 0, 0)
	*/
	/*
		pathField := views.NewSavePathField("Standard Patch")
		layout := widgets.NewQVBoxLayout2(centralWidget)
		layout.AddWidget(pathField.Cell, 0, 0)
	*/
	/*
		sliderInput := views.NewSliderInput("Gamma", 0.24)
		layout := widgets.NewQVBoxLayout2(centralWidget)
		layout.AddWidget(sliderInput.Cell, 0, 0)
	*/

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
