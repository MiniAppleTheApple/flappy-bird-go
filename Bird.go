package main

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const INITIAL_BIRD_GRAVITY float64 = 0.25
const INITIAL_BIRD_MOVE_SPEED float64 = 5.0
const INITIAL_BIRD_VELOCITY float64 = 0

type Bird struct {
	x, y, width, height, velocity, move_speed, gravity float64
	image                                              *ebiten.Image
}

func DefaultBird(image *ebiten.Image) *Bird {
	width, height := image.Size()
	return &Bird{
		x:          float64((WIDTH / 2) + width),
		y:          float64((HEIGHT / 2) + height),
		width:      float64(width),
		height:     float64(height),
		velocity:   INITIAL_BIRD_VELOCITY,
		move_speed: INITIAL_BIRD_MOVE_SPEED,
		gravity:    INITIAL_BIRD_GRAVITY,
		image:      image,
	}
}

func (bird *Bird) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(bird.x, bird.y)
	op.GeoM.Scale(1, 1)
	screen.DrawImage(bird.image, op)
}

func (bird *Bird) Update() {
	bird.move_speed = 5.0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		bird.move_speed += 0.05
		bird.velocity = -(bird.gravity + bird.move_speed)
	}
	bird.gravity = 0.5
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		bird.gravity += 0.05
	}
	bird.velocity += bird.gravity
	bird.y += bird.velocity
}

func (bird Bird) GetXPosition() float64 {
	return bird.x
}

func (bird Bird) GetYPosition() float64 {
	return bird.y
}

func (bird *Bird) SetX(x float64) *Bird {
	bird.x = x
	return bird
}

func (bird *Bird) SetY(y float64) *Bird {
	bird.y = y
	return bird
}

func (bird *Bird) SetHeight(height float64) *Bird {
	bird.height = height
	return bird
}

func (bird *Bird) SetWidth(width float64) *Bird {
	bird.width = width
	return bird
}

func (bird Bird) IsOutOfScreen() bool {
	return bird.y > float64(WIDTH)-bird.width
}

func (bird Bird) IsPassedThrough(pipe *Pipe) bool {
	return pipe.x > bird.x
}
