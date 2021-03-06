package main

import (
	"PixelToolWindow/views"
	"os"

	"github.com/asaskevich/EventBus"
	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Notification
	bus := EventBus.New()

	// main window initializer
	window := widgets.NewQMainWindow(nil, 0)
	window.Resize2(1920, 400)
	window.SetWindowTitle("Pixel Tool Desktop")

	// central widget
	centralWidget := widgets.NewQWidget(nil, 0)
	topWindow := views.NewTopWindow(bus)
	layout := widgets.NewQVBoxLayout2(centralWidget)
	layout.AddWidget(topWindow.Cell, 0, 0)

	// apply layout
	window.SetCentralWidget(centralWidget)

	// show window
	window.Show()

	app.Exec()
}
