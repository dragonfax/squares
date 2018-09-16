package squares

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type BoxFit uint8

const (
	BoxFitCover BoxFit = iota
)

var _ ElementWidget = Image{}
var _ Element = &ImageElement{}

type Image struct {
	File string
	Fit  BoxFit
}

func (i Image) createElement() Element {
	ie := &ImageElement{}
	ie.widget = i
	return ie
}

type ImageElement struct {
	loaded  bool
	surface *sdl.Surface
	texture *sdl.Texture
	elementData
}

func (ie *ImageElement) load() error {
	if !ie.loaded {
		ie.loaded = true

		surface, err := img.Load(ie.widget.(Image).File)
		if err != nil {
			return err
		}
		ie.surface = surface
	}
	return nil
}

func (ie *ImageElement) layout(c Constraints) error {
	err := ie.load()
	if err != nil {
		return err
	}

	if ie.surface == nil {
		ie.size = Size{}
	} else {
		width := float64(ie.surface.W)
		height := float64(ie.surface.H)
		size := Size{Width: width, Height: height}
		ie.size = c.constrainWithRatio(size)
	}

	return nil
}

func (ie *ImageElement) render(o Offset, r *sdl.Renderer) {
	if ie.loaded && ie.texture == nil {
		texture, err := r.CreateTextureFromSurface(ie.surface)
		if err != nil {
			panic(err)
		}
		ie.texture = texture
	}

	if ie.texture != nil && ie.size.Width != 0 && ie.size.Height != 0 {
		dstRect := &sdl.Rect{X: int32(o.x), Y: int32(o.y), W: int32(ie.size.Width), H: int32(ie.size.Height)}
		r.Copy(ie.texture, nil, dstRect)
	}

}
