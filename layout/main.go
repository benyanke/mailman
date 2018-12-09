// Handles the onscreen layout

package layout

import (
	"github.com/jroimartin/gocui"
	"log"
)

func layout(g *gocui.Gui) error {

	maxX, maxY := g.Size()
	// TODO: Add debug-level logging for max size params

	// Enable mouse support once it is done for clicking
	// g.Mouse = true

	if v, err := g.SetView("side", 1, 1, int(0.2*float32(maxX)), maxY-1); err != nil {
		// Handle errors
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set settings for pane
		// TODO: Make this configurable
		v.Title = "Folders"
		v.Editable = true
		// TODO: Make this configurable
		v.Wrap = true
	}

	if v, err := g.SetView("main", int(0.2*float32(maxX)), 1, maxX-1, maxY-1); err != nil {
		// Handle errors
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set settings for pane
		// TODO: Make this configurable
		v.Title = "Messages"
		v.Editable = true
		// TODO: Make this configurable
		v.Wrap = true
	}

	if v, err := g.SetView("cmdline", 1, maxY-5, maxX-1, maxY-1); err != nil {
		// Handle errors
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set settings for pane
		// TODO: Make this configurable
		v.Title = "Content"
		v.Editable = false
		// TODO: Make this configurable
		v.Wrap = true
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
