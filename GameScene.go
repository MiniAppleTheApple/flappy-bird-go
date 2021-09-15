package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	is_pipe_generating bool
	bird               *Bird
	pipes              *[]*Pipe
}

func NewGameScene(bird *Bird) *GameScene {
	return &GameScene{
		is_pipe_generating: false,
		bird:               bird,
		pipes:              &[]*Pipe{},
	}
}
func (scene GameScene) Draw(game Game, screen *ebiten.Image) {
	DrawText(screen, strconv.Itoa(int(game.score)), WIDTH/2, HEIGHT/2)
	scene.bird.Draw(screen)
	for _, v := range *scene.pipes {
		v.Draw(screen)
	}
}
func (scene *GameScene) Update(game *Game) error {
	scene.bird.Update()

	if scene.bird.IsOutOfScreen() {
		game.Die()
	}

	new_pipes := &[]*Pipe{}

	for _, pipe := range *scene.pipes {
		if !pipe.is_passed && scene.bird.IsPassedThrough(pipe) {
			pipe.PassThrough()
			game.score += 0.5
		}
		if pipe.IsOutOfScreen() {
			continue
		}

		pipe.Update(scene.bird, game)

		*new_pipes = append(*new_pipes, pipe)
	}

	if !scene.is_pipe_generating {
		go scene.SpawnPipe()
	}

	*scene.pipes = *new_pipes
	return nil
}

func (scene *GameScene) AddPipe(pipe *Pipe) {
	*scene.pipes = append(*scene.pipes, pipe)
}

func (scene *GameScene) SpawnPipe() {
	var top float64
	var bottom float64

	top = float64(rand.Intn(300) + 50)
	bottom = top + 160

	scene.is_pipe_generating = true

	time.Sleep(time.Millisecond * 3000)

	top_pipe := DefaultPipe(pipe_resource).SetY(top - 512)
	bottom_pipe := DefaultPipe(pipe_resource).SetY(bottom)

	scene.AddPipe(top_pipe)
	scene.AddPipe(bottom_pipe)

	scene.is_pipe_generating = false
}
