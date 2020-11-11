package main

import (
	"fmt"
	"image"
	"os"

	"image/draw"
	"image/jpeg"
)

func main() {

	lowerImagePath := "material/100x100.jpg"
	upperImagePath := "material/50x50.jpg"
	outputImagePath := "superposition-out.jpg"

	lowerImageOpen, err := os.Open(lowerImagePath)
	if err != nil {
		fmt.Println(err)
	}

	upperImageOpen, err := os.Open(upperImagePath)
	if err != nil {
		fmt.Println(err)
	}

	lowerImage, _, err := image.Decode(lowerImageOpen)
	if err != nil {
		fmt.Println(err)
	}

	upperImage, _, err := image.Decode(upperImageOpen)
	if err != nil {
		fmt.Println(err)
	}

	upperStartPoint := image.Point{25, 25}
	upperRect := image.Rectangle{upperStartPoint, upperStartPoint.Add(upperImage.Bounds().Size())}
	lowerRect := image.Rectangle{image.Point{0, 0}, lowerImage.Bounds().Size()}

	rgba := image.NewRGBA(lowerRect)
	draw.Draw(rgba, lowerRect, lowerImage, image.Point{0, 0}, draw.Src)
	draw.Draw(rgba, upperRect, upperImage, image.Point{0, 0}, draw.Over)

	out, err := os.Create(outputImagePath)
	if err != nil {
		fmt.Println(err)
	}

	var opt jpeg.Options
	opt.Quality = 85
	jpeg.Encode(out, rgba, &opt)
}
