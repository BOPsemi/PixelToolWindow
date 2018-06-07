package views

import (
	"PixelToolWindow/models"

	"github.com/asaskevich/EventBus"
	"github.com/therecipe/qt/widgets"
)

/*
SideWindow :Side window
*/
type SideWindow struct {
	// --- Enviroment Setting group ---
	gammaAdjuster       *SliderInput         // slider object for gamma correction
	lightSourceSelector *ComboBoxSelector    // light source selector
	patchSizeInput      *PixelSizeInputField // patch size setting

	// --- file save group ---
	stdPatchSave *SavePathField // standard Macbeth Patch save point
	devPatchSave *SavePathField // Simulated Macbeth Pathc save point

	// --- input file group ---
	deviceQEData   *InputField // device QE file input
	whitePixelData *InputField // white pixel file input

	// Apply button
	button *widgets.QPushButton // apply button

	// Cell
	Cell *widgets.QWidget
}

/*
NewSideWindow :initializer of SideWindow
*/
func NewSideWindow(bus EventBus.Bus) *SideWindow {
	obj := new(SideWindow)

	// initialize widgets
	obj.Cell = widgets.NewQWidget(nil, 0)

	// initalize button
	obj.button = widgets.NewQPushButton2("Apply", obj.Cell)

	// initialize each gourp
	envGroup := obj.setupEnvGroup()
	fileSaveGroup := obj.setFileSaveGroup()
	inputFileGroup := obj.setInputFileGroup()

	// layout
	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(envGroup, 0, 0)
	layout.AddWidget(fileSaveGroup, 0, 0)
	layout.AddWidget(inputFileGroup, 0, 0)
	layout.AddWidget(obj.button, 0, 0)

	// apply layout
	obj.Cell.SetLayout(layout)

	// action connection
	obj.button.ConnectClicked(func(checked bool) {
		info := new(models.SettingInfo)

		info.Gamma = obj.gammaAdjuster.Value
		info.LightSource = obj.lightSourceSelector.SelectedItem
		info.PatchSize.H = obj.patchSizeInput.HorizontalSize
		info.PatchSize.V = obj.patchSizeInput.VerticalSize
		info.StdPatchSavePath = obj.stdPatchSave.textLabelForPath.Text()
		info.DevPatchSavePath = obj.devPatchSave.textLabelForPath.Text()
		info.DeiceQEDataPath = obj.deviceQEData.textField.Text()
		info.WhitePixelDataPath = obj.whitePixelData.textField.Text()

		bus.Publish("sideWin:settingInfo", info)
	})

	return obj
}

// Enviroment setting group
func (sw *SideWindow) setupEnvGroup() *widgets.QGroupBox {
	sw.gammaAdjuster = NewSliderInput("Gamma", 0.24)
	sw.lightSourceSelector = NewComboBoxSelector("Light Source", []string{"D65", "D50", "Ill-A"})
	sw.patchSizeInput = NewPixelSizeInputField("Patch Size", 100, 100)

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(sw.gammaAdjuster.Cell, 0, 0)
	layout.AddWidget(sw.lightSourceSelector.Cell, 0, 0)
	layout.AddWidget(sw.patchSizeInput.Cell, 0, 0)

	group := widgets.NewQGroupBox2("Simulation Enviroment Setting", nil)
	group.SetLayout(layout)

	return group
}

// file save group
func (sw *SideWindow) setFileSaveGroup() *widgets.QGroupBox {
	sw.stdPatchSave = NewSavePathField("Std Patch Save")
	sw.devPatchSave = NewSavePathField("Dev Patch Save")

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(sw.stdPatchSave.Cell, 0, 0)
	layout.AddWidget(sw.devPatchSave.Cell, 0, 0)

	group := widgets.NewQGroupBox2("File Save Setting", nil)
	group.SetLayout(layout)

	return group
}

// input file group
func (sw *SideWindow) setInputFileGroup() *widgets.QGroupBox {
	sw.deviceQEData = NewInputField("Device QE", "Device QE raw data")
	sw.whitePixelData = NewInputField("White Pixel", "White pixel raw data")

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(sw.deviceQEData.Cell, 0, 0)
	layout.AddWidget(sw.whitePixelData.Cell, 0, 0)

	group := widgets.NewQGroupBox2("Input File information", nil)
	group.SetLayout(layout)

	return group
}
