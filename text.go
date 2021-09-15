package main

import (
	"image/color"
	"io/ioutil"
	"log"

	"golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/opentype"
)

func LoadFontByPath(path string, size float64) font.Face {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	typography, err := opentype.Parse(content)

	_font, err := opentype.NewFace(typography, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	return _font
}
func DrawText(screen *ebiten.Image, _text string, x, y int) {
	text.Draw(screen, _text, press_start_2p, x, y, color.White)
}
