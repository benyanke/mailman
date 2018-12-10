package main

import (
	"fmt"
	"github.com/benyanke/mailman/config"
	"github.com/benyanke/mailman/imap"
	"github.com/benyanke/mailman/layout"
	"github.com/jroimartin/gocui"
	"log"
)

func notmain() {
	// Working proof-of-concept for fetching imap mailboxes and mail
	imap.Test()
}

func main() {
	layout.Run()

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layoutFunc)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layoutFunc(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	var sizeX int = 40
	var sizeY int = 10
	if v, err := g.SetView("hello", maxX/2-sizeX, maxY/2, maxX/2+sizeX, maxY/2+sizeY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Reading config from "+config.GetConfigDir())
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
