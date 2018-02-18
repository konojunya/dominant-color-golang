package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("sample.jpg")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	img = resize.Resize(36, 0, img, resize.Lanczos3)
	rgb := getDominantColor(img)

	output(rgb)
}

func output(c color.RGBA) {
	var x, y int
	w := 200
	h := 100

	img := image.NewRGBA(image.Rect(x, y, w, h))
	rect := img.Rect

	// X, Yに塗っていく
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, c)
		}
	}

	file, err := os.Create("output.jpg")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	if err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100}); err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}

func getDominantColor(img image.Image) color.RGBA {
	var r, g, b, count float64

	rect := img.Bounds()
	for i := 0; i < rect.Max.Y; i++ {
		for j := 0; j < rect.Max.X; j++ {
			c := color.RGBAModel.Convert(img.At(j, i))
			r += float64(c.(color.RGBA).R)
			g += float64(c.(color.RGBA).G)
			b += float64(c.(color.RGBA).B)
			count++
		}
	}

	return color.RGBA{
		R: uint8(r / count),
		G: uint8(g / count),
		B: uint8(b / count),
	}
}
