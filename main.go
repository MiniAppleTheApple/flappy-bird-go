package main

import (
	"bytes"
	"image"
	_ "image/png"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

const WIDTH int = 600
const HEIGHT int = 600

var bird_resource *ebiten.Image
var pipe_resource *ebiten.Image
var press_start_2p font.Face

func LoadImageFromPath(path string) image.Image {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func main() {
	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Flappy Bird")

	bird_image := LoadImageFromPath("duck.png")
	bird_resource = ebiten.NewImageFromImage(bird_image)
	bird := DefaultBird(
		bird_resource,
	)

	pipe_image := LoadImageFromPath("pipe.png")
	pipe_resource = ebiten.NewImageFromImage(pipe_image)

	press_start_2p = LoadFontByPath("PressStart2P-Regular.ttf", 36)

	game := Game{
		score: 0,
		scene: NewGameScene(bird),
	}

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
