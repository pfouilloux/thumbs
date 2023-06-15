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

func TestSquare(t *testing.T) {
	baseImg := newGradientImage(1000, 1000)

	scns := map[string]scenario{
		"should not crop image with same size as target": {
			src:          baseImg.SubImage(image.Rect(0, 0, 100, 100)),

			want: baseImg.SubImage(image.Rect(0, 0, 100, 100)),
		},
	}

	for name, scn := range scns {
		t.Run(name, scn.testSquare)
	}
}

func (scn scenario) testSquare(t *testing.T) {
	t.Parallel()

	got := img.Square(scn.src)

	scn.assertIsExpectedSize(t, got)
	scn.assertContainsExpectedPixels(t, got)
}

func (scn scenario) assertIsExpectedSize(t testing.TB, got image.Image) {
	if scn.want.Bounds().Min != got.Bounds().Min || scn.want.Bounds().Max != got.Bounds().Max {
		t.Errorf("expected an %dx%d image but got an %dx%d image", scn.want.Bounds().Dx(), scn.want.Bounds().Dy(), got.Bounds().Dx(), got.Bounds().Dx())
	}
}

func (scn scenario) assertContainsExpectedPixels(t testing.TB, got image.Image) {
	if !reflect.DeepEqual(scn.want, got) {
		t.Errorf("the cropped image was not the same as the expected subimage")
	}
}

func newGradientImage(w, h int) *image.RGBA {
	gradient := image.NewRGBA(image.Rect(0, 0, w, h))
	size := gradient.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			gradient.Set(x, y, interpolateColour(size, x, y))
		}
	}

	return gradient
}

func interpolateColour(sz image.Point, x, y int) color.RGBA {
	return color.RGBA{
		R: uint8(255 * x / sz.X),
		G: uint8(255 * y / sz.Y),
		B: 55,
		A: 255,
	}
}
