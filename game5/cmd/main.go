package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"pl.home/game5/internal/cfg"
	"pl.home/game5/internal/game"
)

func main() {

	game := game.NewGame()

	ebiten.SetWindowSize(cfg.WindowWidth, cfg.WindowHeigh)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
