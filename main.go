package main

import (
	"image"
	"image/jpeg"
	"os"

	"golang.org/x/image/draw"
)

func main() {
	origFile, err := os.Open("testdata/test.jpeg")
	if err != nil {
		// We do want to handle errors because it's handy to see where things can go wrong, but no need to get fancy in a spike
		panic(err)
	}

	orig, _, err := image.Decode(origFile) // We don't care about the image format for the spike
	if err != nil {
		panic(err)
	}

	// this next code block is pretty messy and there's a lot of unnecessary repetition, but that's okay in a spike, we'll fix that up in the prod code
	// in a nutshell this code crops the image into a square around the center of the original image
	square := orig
	if orig.Bounds().Dx() < orig.Bounds().Dy() {
		subImager := orig.(interface {
			SubImage(r image.Rectangle) image.Image
		})

		adj := (orig.Bounds().Dy() - orig.Bounds().Dx()) / 2
		square = subImager.SubImage(image.Rect(orig.Bounds().Min.X, orig.Bounds().Min.Y+adj, orig.Bounds().Max.X, orig.Bounds().Max.Y-adj))
	} else if orig.Bounds().Dx() > orig.Bounds().Dy() {
		subImager := orig.(interface {
			SubImage(r image.Rectangle) image.Image
		})

		adj := (orig.Bounds().Dx() - orig.Bounds().Dy()) / 2
		square = subImager.SubImage(image.Rect(orig.Bounds().Min.X, orig.Bounds().Min.Y+adj, orig.Bounds().Max.X, orig.Bounds().Max.Y-adj))
	}

	// Here we do the actual resizing
	thumb := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.BiLinear.Scale(thumb, image.Rect(0, 0, 100, 100), square, square.Bounds(), draw.Src, nil)

	// we write to a file to see the results of our work
	thumbFile, err := os.Create("testout/test_thumb.jpeg")
	if err != nil {
		panic(err)
	}
	defer thumbFile.Close()

	// hardcoded to jpeg here for the spike but we'll want to keep the original encoding in the prod code
	if err := jpeg.Encode(thumbFile, thumb, nil); err != nil {
		panic(err)
	}
}
