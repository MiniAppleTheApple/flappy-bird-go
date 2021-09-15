package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Draw(Game, *ebiten.Image)
	Update(*Game) error
}
