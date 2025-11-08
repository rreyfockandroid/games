package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"pl.home/game2/conf"
	"pl.home/game2/stage"
)

func main() {
	// scene := stage.NewScene()
	game := stage.NewGame()

	ebiten.SetWindowSize(conf.WindowWidth, conf.WindowHeigh)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
