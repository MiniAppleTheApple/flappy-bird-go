package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type DeadScene struct {
	highest_score int
}

func NewDeadScene(highest_score int) DeadScene {
	return DeadScene{
		highest_score: highest_score,
	}
}
func (scene DeadScene) Draw(game Game, screen *ebiten.Image) {
	DrawText(screen, fmt.Sprintf("Your Score is %v", int(game.score)), 24, HEIGHT/2)
	DrawText(screen, "Press space", WIDTH/2-200, HEIGHT-200)
	if scene.highest_score < int(game.score) {
		DrawText(screen, "Highest Score", WIDTH/2-240, 200)
	}
}

func (scene DeadScene) Update(game *Game) error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		bird := DefaultBird(
			bird_resource,
		)
		game.score = 0
		game.scene = NewGameScene(bird)
	}
	return nil
}
