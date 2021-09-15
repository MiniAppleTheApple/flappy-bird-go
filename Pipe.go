package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Pipe struct {
	x, y, width, height, move_speed float64
	is_passed                       bool
	image                           *ebiten.Image
}

const INITIAL_PIPE_GRAVITY float64 = 2.0

func DefaultPipe(image *ebiten.Image) *Pipe {
	width, height := image.Size()
	return &Pipe{
		x:          0,
		y:          0,
		width:      float64(width),
		height:     float64(height),
		is_passed:  false,
		move_speed: INITIAL_PIPE_GRAVITY,
		image:      image,
	}
}

func (pipe *Pipe) SetX(x float64) *Pipe {
	pipe.x = x
	return pipe
}

func (pipe *Pipe) SetY(y float64) *Pipe {
	pipe.y = y
	return pipe
}

func (pipe *Pipe) SetWidth(width float64) *Pipe {
	pipe.width = width
	return pipe
}

func (pipe *Pipe) SetHeight(height float64) *Pipe {
	pipe.height = height
	return pipe
}

func (pipe *Pipe) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pipe.x, pipe.y)
	op.GeoM.Scale(1, 1)
	screen.DrawImage(pipe.image, op)
}

func (pipe *Pipe) Update(bird *Bird, game *Game) {
	if pipe.IsCollisionWithBird(bird) {
		game.Die()
	}
	pipe.x += pipe.move_speed
}

func (pipe *Pipe) IsOutOfScreen() bool {
	return pipe.x > float64(HEIGHT)
}

func (pipe *Pipe) IsCollisionWithBird(bird *Bird) bool {
	return pipe.x < bird.x+bird.width &&
		pipe.x+pipe.width > bird.x &&
		pipe.y < bird.y+bird.height &&
		pipe.y+pipe.height > bird.y
}
func (pipe *Pipe) PassThrough() {
	pipe.is_passed = true
}
