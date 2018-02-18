package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	file, err := os.Open("sample.jpg")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	img = resize.Resize(36, 0, img, resize.Lanczos3)
	hist := getHistgram(img)

	max := 0

	for _, val := range hist {
		if val > max {
			max = val
		}
	}
	r, g, b := int2rgb(max)
	colorCode := "#" + dec2hex(r, 2) + dec2hex(g, 2) + dec2hex(b, 2)

	fmt.Println(colorCode)

}

func getHistgram(img image.Image) []int {
	hist := make([]int, 1000000)

	rect := img.Bounds()
	for i := 0; i < rect.Max.Y; i++ {
		for j := 0; j < rect.Max.X; j++ {
			r, g, b, _ := img.At(j, i).RGBA()
			i := rgb2int(int(r), int(g), int(b))
			hist[i]++
		}
	}

	return hist
}

func dec2hex(n, beam int) string {
	hex := ""
	str := "0123456789abcdef"
	for i := 0; i < beam; i++ {
		m := n & 0xf
		hex = string(str[m]) + hex
		n -= m
		n >>= 4
	}

	return hex
}

func rgb2int(r, g, b int) int {
	return (((r >> 5) << 6) | ((g >> 5) << 3) | ((b >> 5) << 0))
}

func int2rgb(i int) (r, g, b int) {
	return ((i >> 6 & 0x7) << 5) + 16, ((i >> 3 & 0x7) << 5) + 16, ((i >> 0 & 0x7) << 5) + 16
}
