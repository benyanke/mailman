// Handles the onscreen layout

package layout

import (
//	"fmt"
	"github.com/jroimartin/gocui"
	"log"
//    "strconv"
)

func layout(g *gocui.Gui) error {

	// Hardcoded config vars
	var showPaneHeaders bool = true

	// TODO: Abstract the layout to an array of possible
	// layouts w/ config options to select which is used

	maxX, maxY := g.Size()
	// TODO: Add debug-level logging for max size params

	// Enable mouse support once it is done for clicking
	// TODO: Make this a config option
	// g.Mouse = true

	// Formula for the x value of the folder/message seperator
	// Currently, 20% of screen width
	var verticalSplitPoint int = int(0.2 * float32(maxX))

	// Formula for the y value of the mail body seperator
	// Currently 40% of screen height
	var bodySplitPoint int = maxY - int(0.4*float32(maxY))

	// TODO: This min doesn't seem to be fully working
	// Need to add checks below, because it causes inverted coords for panes below
	// var bodyAbsoluteMin int = 20
	// if bodySplitPoint < bodyAbsoluteMin {
	//	bodySplitPoint = bodyAbsoluteMin
	//	fmt.Println("BSMIN")
	//}

	// TODO: add these as debugging log options
	// fmt.Println("bs:" + strconv.Itoa(bodySplitPoint))
	// fmt.Println("vs:" + strconv.Itoa(verticalSplitPoint))

	if v, err := g.SetView("folders", 1, 1, verticalSplitPoint, bodySplitPoint); err != nil {
		// Handle errors
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set settings for pane
		if showPaneHeaders {
			v.Title = "Folders"
		}
		v.Editable = true
		// TODO: Make this configurable
		v.Wrap = true
	}

	if v, err := g.SetView("messages", verticalSplitPoint, 1, maxX-1, bodySplitPoint); err != nil {
		// Handle errors
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set settings for pane
		if showPaneHeaders {
			v.Title = "Messages"
		}
		v.Editable = true
		// TODO: Make this configurable
		v.Wrap = true
	}

	if v, err := g.SetView("body", 1, bodySplitPoint, maxX-1, maxY-1); err != nil {
		// Handle errors
		if err != gocui.ErrUnknownView {
			return err
		}

		// Set settings for pane
		if showPaneHeaders {
			v.Title = "Body"
		}

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
