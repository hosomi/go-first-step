package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"golang.org/x/image/draw"
)

func main() {
	flag.Parse()
	args := flag.Args()

	f, err := os.Open(args[0])
	if err != nil {
		fmt.Println("open:", err)
		return
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("decode:", err)
		return
	}

	fso, err := os.Create(args[1])
	if err != nil {
		fmt.Println("create:", err)
		return
	}
	defer fso.Close()

	rct := img.Bounds()

	dst := image.NewRGBA(image.Rect(0, 0, rct.Dx()/2, rct.Dy()/2))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, rct, draw.Over, nil)

	jpeg.Encode(fso, dst, &jpeg.Options{Quality: 100})
}
