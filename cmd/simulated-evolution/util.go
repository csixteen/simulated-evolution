package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
)

type SpritesSheet struct {
	sprites pixel.Picture
	frames  []pixel.Rect
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func NewSpritesSheet(path string, side float64) *SpritesSheet {
	ss, err := loadPicture(path)
	if err != nil {
		panic(err)
	}

	sheet := &SpritesSheet{
		sprites: ss,
	}

	for x := ss.Bounds().Min.X; x < ss.Bounds().Max.X; x += side {
		for y := ss.Bounds().Min.Y; y < ss.Bounds().Max.Y; y += side {
			sheet.frames = append(sheet.frames, pixel.R(x, y, x+side, y+side))
		}
	}

	return sheet
}
