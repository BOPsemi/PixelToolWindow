package views

import (
	"time"

	"github.com/asaskevich/EventBus"
	"github.com/therecipe/qt/widgets"
)

const (
	nodataImagePath = "/Users/kazufumiwatanabe/go/src/PixelToolWindow/data/NoDataImage.png"
)

/*
ImageViewIdentifier : vierwer identifier
*/
type ImageViewIdentifier int

const (
	StdImageViewer ImageViewIdentifier = iota
	DevImageViewer
)

/*
MainWindow :main window
*/
type MainWindow struct {

	// --- Image Viewer ---
	stdCCImageView *ImageViewer // standard Macbeth color chart
	devCCImageView *ImageViewer // device Macbeth color chart

	// --- load buttons ---
	stdImageLoadButton *widgets.QPushButton // image load button for standard
	devImageLoadButton *widgets.QPushButton // image load button for device

	// --- message box ---
	messageBox *widgets.QTextEdit // message box

	// widget
	Cell *widgets.QWidget
}

/*
NewMainWindow : initializer of main window
*/
func NewMainWindow(bus EventBus.Bus) *MainWindow {
	obj := new(MainWindow)

	obj.Cell = widgets.NewQWidget(nil, 0)

	// imageViewer initialize
	obj.stdCCImageView = NewImageViewer(nodataImagePath, 0.5)
	obj.devCCImageView = NewImageViewer(nodataImagePath, 0.5)

	// button
	obj.stdImageLoadButton = widgets.NewQPushButton2("Image Load", obj.Cell)
	obj.devImageLoadButton = widgets.NewQPushButton2("Image Load", obj.Cell)

	// message button
	initlog := "Logging started" + "  :  " + time.Now().Format(time.ANSIC)
	obj.messageBox = widgets.NewQTextEdit2(initlog, obj.Cell)

	// group
	stdGroup := widgets.NewQGroupBox2("Standard Macbeth Color Chart", obj.Cell)
	devGroup := widgets.NewQGroupBox2("Device Macbeth Color Chart", obj.Cell)

	// layout
	stdLayout := widgets.NewQVBoxLayout()
	stdLayout.AddWidget(obj.stdCCImageView.Cell, 0, 0)
	stdLayout.AddWidget(obj.stdImageLoadButton, 0, 0)
	stdGroup.SetLayout(stdLayout)

	devLayout := widgets.NewQVBoxLayout()
	devLayout.AddWidget(obj.devCCImageView.Cell, 0, 0)
	devLayout.AddWidget(obj.devImageLoadButton, 0, 0)
	devGroup.SetLayout(devLayout)

	// layout
	layout := widgets.NewQGridLayout2()
	layout.AddWidget(stdGroup, 0, 0, 0)
	layout.AddWidget(devGroup, 0, 1, 0)
	layout.AddWidget3(obj.messageBox, 1, 0, 1, 3, 0)

	// apply layout
	obj.Cell.SetLayout(layout)

	// action connection
	obj.stdImageLoadButton.ConnectClicked(func(checked bool) {
		obj.reloadImage("/Users/kazufumiwatanabe/go/src/PixelToolWindow/data/std_macbeth_chart.png", 0.5, StdImageViewer)
		bus.Publish("main:message", "Standard Macbeth Color Chart was reloded")
	})
	obj.devImageLoadButton.ConnectClicked(func(checked bool) {
		obj.reloadImage("/Users/kazufumiwatanabe/go/src/PixelToolWindow/data/std_macbeth_chart.png", 0.5, DevImageViewer)
		bus.Publish("main:message", "Device Macbeth Color Chart was reloded")
	})

	return obj
}

// reloadImage
func (mm *MainWindow) reloadImage(path string, scale float64, identifier ImageViewIdentifier) {
	switch identifier {
	case StdImageViewer:
		mm.stdCCImageView.SetImageView(path, scale)
		mm.stdCCImageView.Cell.Repaint()

	case DevImageViewer:
		mm.devCCImageView.SetImageView(path, scale)
		mm.devCCImageView.Cell.Repaint()
	}
}
