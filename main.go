package main

import (
	"PixelToolWindow/models"
	"PixelToolWindow/views"
	"fmt"
	"os"

	"github.com/asaskevich/EventBus"
	"github.com/therecipe/qt/widgets"
)

func settingInfoReceiver(info *models.SettingInfo) {
	fmt.Println(info)
}

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Notification
	bus := EventBus.New()
	bus.Subscribe("sideWin:settingInfo", settingInfoReceiver)

	centralWidget := widgets.NewQWidget(nil, 0)

	/*
		sideWindow := views.NewSideWindow(bus)
		layout := widgets.NewQVBoxLayout2(centralWidget)
		layout.AddWidget(sideWindow.Cell, 0, 0)
	*/

	imageView := views.NewImageViewer("/Users/kazufumiwatanabe/go/src/PixelToolWindow/std_macbeth_chart.png", 0.5)
	//imageView := views.NewImageViewer("", 0.5)
	layout := widgets.NewQVBoxLayout2(centralWidget)
	layout.AddWidget(imageView.Cell, 0, 0)

	/*
		pixelInput := views.NewPixelSizeInputField("Patch size", 100, 100)
		layout := widgets.NewQVBoxLayout2(centralWidget)
		layout.AddWidget(pixelInput.Cell, 0, 0)
	*/

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
