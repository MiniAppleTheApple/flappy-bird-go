package main

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	score float32
	scene Scene
}

func (game *Game) Update() error {
	return game.scene.Update(game)
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.scene.Draw(*game, screen)
}

func (game *Game) Layout(width, height int) (int, int) {
	return width, height
}

// func (game *Game) ReStart(bird) {
// 	ioutil.WriteFile("highestscore.txt",game.score)
// 	game.scene = NewGameScene(bird)
// }
func (game *Game) Die() {
	content, err := ioutil.ReadFile("highestscore.txt")
	if err != nil {
		log.Fatal(err)
	}
	number, err := strconv.Atoi(string(content))
	if err != nil {
		log.Fatal(err)
	}
	game.scene = NewDeadScene(number)
}
