package views

import (
	"github.com/asaskevich/EventBus"
	"github.com/therecipe/qt/widgets"
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

	return obj
}
