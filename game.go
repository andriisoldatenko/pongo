package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	screen tcell.Screen
	ball   Ball
}

func (g *Game) Run() {

	defStyle := tcell.StyleDefault.Background(tcell.ColorDefault).Foreground(tcell.ColorDefault)
	g.screen.SetStyle(defStyle)

	for {
		width, height := g.screen.Size()
		g.ball.CheckEdges(width, height)

		g.screen.Clear()

		g.ball.Update()

		g.screen.SetContent(g.ball.X, g.ball.Y, g.ball.Display(), nil, tcell.StyleDefault)
		time.Sleep(40 * time.Millisecond)
		g.screen.Show()
	}

}
