package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

type DominantColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func main() {
	file, err := os.Open("sample.jpg")
	defer file.Close()
	if err != nil {
		log.Fatalf("open err: %v", err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("decode err: %v", err)
	}

	img = resize.Resize(36, 0, img, resize.Lanczos3)
	rgba := getDominantColor(img)

	fmt.Println(rgba)

}

func getDominantColor(img image.Image) DominantColor {
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

	return DominantColor{
		Red:   uint8(r / count),
		Green: uint8(g / count),
		Blue:  uint8(b / count),
	}
}
