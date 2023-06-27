package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math/cmplx"
	"math/rand"
	"os"
)

// todo
// - go routine
// - live updates
// - cursor input on image / web canvas

const maxIter = 100

var colorMapping = make(map[int]color.RGBA64)

func f(c complex128) color.RGBA64 {
	z := complex(0, 0)
	for i := 0; i < maxIter; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			clr, ok := colorMapping[i]
			if !ok {
				clr = color.RGBA64{uint16(rand.Intn(65536)), uint16(rand.Intn(65536)), 65535, 65535}
				colorMapping[i] = clr
			}
			return clr
		}
	}
	return color.RGBA64{0, 0, 0, 65535}
}

func getPoints(img *image.RGBA64) {
	for i := -1.25; i <= 1.25; i += 0.0005 {
		for j := -2.0; j <= 0.5; j += 0.0005 {
			clr := f(complex(j, i))
			img.SetRGBA64(int(j*2000+4000), int(i*2000+2500), clr)
		}
	}
}

func main() {
	img := image.NewRGBA64(image.Rect(0, 0, 5000, 5000))
	getPoints(img)
	fmt.Println("Computed!")
	out, err := os.Create("display.jpeg")
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	defer out.Close()
	jpeg.Encode(out, img, nil)
	fmt.Printf("Colors Used %v", len(colorMapping))
}
