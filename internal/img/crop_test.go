package img_test

import (
	"image"
	"image/color"
	"reflect"
	"testing"
	"thumbs/internal/img"
)

type scenario struct {
	src                       image.Image
	targetWidth, targetHeight int

	want image.Image
}

func TestCrop(t *testing.T) {
	baseImg := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	size := baseImg.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			clr := color.RGBA{R: uint8(255 * x / size.X), G: uint8(255 * y / size.Y), B: 55, A: 255}
			baseImg.Set(x, y, clr)
		}
	}

	scns := map[string]scenario{
		"should not crop image with same size as target": {
			src:          baseImg.SubImage(image.Rect(0, 0, 100, 100)),
			targetWidth:  100,
			targetHeight: 100,

			want: baseImg.SubImage(image.Rect(0, 0, 100, 100)),
		},
	}

	for name, scn := range scns {
		t.Run(name, scn.testCrop)
	}
}

func (scn scenario) testCrop(t *testing.T) {
	t.Parallel()

	got := img.Crop(scn.src, scn.targetWidth, scn.targetHeight)

	// we check the size of the images to tell us if we've cropped an image of the right size
	if scn.want.Bounds() != got.Bounds() {
		t.Errorf("expected an %dx%d image but got an %dx%d image", scn.want.Bounds().Dx(), scn.want.Bounds().Dy(), got.Bounds().Dx(), got.Bounds().Dx())
	}
	
	if !reflect.DeepEqual(scn.want, got) {
		t.Errorf("the cropped image was not the same as the expected subimage")
	}
}
