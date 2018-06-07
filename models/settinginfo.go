package models

/*
SettingInfo :strcutrue for setting information
*/
type SettingInfo struct {
	Gamma       float64 // gamma value
	LightSource string  // light source info
	PatchSize   struct {
		H int
		V int
	}
	StdPatchSavePath   string // file save path of standard Macbeth chart
	DevPatchSavePath   string // file save path of device Macbeth chart
	DeiceQEDataPath    string // QE data path
	WhitePixelDataPath string // White pixel data path
}
