// Handles the onscreen layout

package layout

import (
	"github.com/jroimartin/gocui"
	"log"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// Enable mouse support once it is done for clicking
	// g.Mouse = true

	// Test view

	if v, err := g.SetView("side", 1, 1, int(0.2*float32(maxX)), maxY-5); err != nil {
		// Handle errors
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set settings for pane
		v.Title = "Folders"
		v.Editable = true
		v.Wrap = true
	}

	if _, err := g.SetView("main", int(0.2*float32(maxX)), -1, maxX, maxY-5); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	if _, err := g.SetView("cmdline", -1, maxY-5, maxX, maxY); err != nil &&
		err != gocui.ErrUnknownView {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func Run() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	return
}
