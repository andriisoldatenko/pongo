package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {

	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	ball := Ball{
		X:      1,
		Y:      1,
		Xspeed: 1,
		Yspeed: 1,
	}

	game := Game{
		screen: screen,
		ball:   ball,
	}

	go game.Run()

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}

}
