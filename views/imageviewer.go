package views

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

/*
ImageViewer :image viewer
*/
type ImageViewer struct {
	reader *gui.QImageReader            // image reader
	pixmap *gui.QPixmap                 // pixel map
	item   *widgets.QGraphicsPixmapItem // pixel item
	scene  *widgets.QGraphicsScene      // graphic scene

	Cell *widgets.QGraphicsView
}

/*
NewImageViewer :initializer of image viewer
*/
func NewImageViewer(path string, scale float64) *ImageViewer {
	obj := new(ImageViewer)

	// initialize widgets
	obj.Cell = widgets.NewQGraphicsView(nil)

	if path != "" {
		obj.reader = gui.NewQImageReader3(path, core.NewQByteArray())
	} else {
		obj.reader = gui.NewQImageReader()
	}

	// make scene
	obj.makeScene(scale)

	return obj
}

// make scene
func (iv *ImageViewer) makeScene(scale float64) {
	iv.pixmap = gui.QPixmap_FromImageReader(iv.reader, core.Qt__AutoColor)
	iv.item = widgets.NewQGraphicsPixmapItem2(iv.pixmap, nil)

	// set scale
	var viewScale float64
	if scale < 0 {
		viewScale = 1.0
	} else {
		viewScale = scale
	}
	iv.item.SetScale(viewScale)

	// add item to scene
	iv.scene = widgets.NewQGraphicsScene(nil)
	iv.scene.AddItem(iv.item)

	// create image view
	iv.Cell.SetScene(iv.scene)
}
