package screen2d

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Sprite represents a single object on the screen
type Sprite struct {
	w, h, pitch int32
	rend        *sdl.Renderer
	surf        *sdl.Surface
	tex         *sdl.Texture
}

// NewSprite returns a new Sprite
func NewSprite(rend *sdl.Renderer) *Sprite {
	s := &Sprite{
		rend:  rend,
		w:     0,
		h:     0,
		pitch: 0,
	}
	return s
}

// LoadRGBAPixels loads the Sprite using an array of 32-bit pixels containing RGBA values
func (s *Sprite) LoadRGBAPixels(pixels []int, pitch int32) error {
	s.w = pitch
	s.h = int32(len(pixels)) / pitch
	s.pitch = pitch

	{
		surf, err := RGBAPixels2Surface(pixels, s.w, s.h)
		if err != nil {
			return err
		}
		s.surf = surf
	}
	{
		tex, err := Surface2Texture(s.rend, s.surf)
		if err != nil {
			return err
		}
		s.tex = tex
	}
	return nil
}

// DrawAt the Sprite onto the screen
func (s *Sprite) DrawAt(x, y int32, scale float32) {
	if s.tex == nil {
		return
	}

	dstRect := &sdl.Rect{
		X: x,
		Y: y,
		W: int32(float32(s.w) * scale),
		H: int32(float32(s.h) * scale),
	}
	s.rend.Copy(s.tex, nil, dstRect)
}

// GetPitch returns the Sprite Pitch
func (s *Sprite) GetPitch() int32 {
	return int32(s.pitch)
}

// GetPixels returns the Sprite Pixels
func (s *Sprite) GetPixels() []byte {
	return s.surf.Pixels()
}
