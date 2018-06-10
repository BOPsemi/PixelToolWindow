package views

import (
	"PixelToolWindow/models"
	"fmt"
	"time"

	"github.com/asaskevich/EventBus"
	"github.com/therecipe/qt/widgets"
)

/*
sideWin:settingInfo
main:message
*/

/*
TopWindow :top window structure
*/
type TopWindow struct {
	sideWin *SideWindow // side window
	mainWin *MainWindow // main window

	Cell *widgets.QWidget
}

/*
NewTopWindow :initializer of top window
*/
func NewTopWindow(bus EventBus.Bus) *TopWindow {
	obj := new(TopWindow)

	obj.Cell = widgets.NewQWidget(nil, 0)

	// initialize both windows
	obj.sideWin = NewSideWindow(bus)
	obj.mainWin = NewMainWindow(bus)

	// layout
	layout := widgets.NewQHBoxLayout()
	layout.AddWidget(obj.sideWin.Cell, 0, 0)
	layout.AddWidget(obj.mainWin.Cell, 0, 0)

	// apply layout
	obj.Cell.SetLayout(layout)

	// event bus subscribe
	bus.Subscribe("sideWin:settingInfo", obj.settingInfoReciever)
	bus.Subscribe("main:message", obj.messageReciever)

	return obj
}

// --- Subscriber ---
func (tw *TopWindow) settingInfoReciever(info *models.SettingInfo) {
	fmt.Println(info)
}

func (tw *TopWindow) messageReciever(message string) {
	tw.mainWin.messageBox.Append(message + "  :  " + time.Now().Format(time.ANSIC))
	tw.mainWin.messageBox.Repaint()
}
