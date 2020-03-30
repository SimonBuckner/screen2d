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

// NewSpriteFromnRGBAPixels returns a new Sprite
func NewSpriteFromnRGBAPixels(screen *Screen, pixels []int, pitch int32) (*Sprite, error) {
	s := &Sprite{
		rend:  screen.Rend(),
		w:     pitch,
		h:     int32(len(pixels)) / pitch,
		pitch: pitch,
	}
	{
		surf, err := RGBAPixels2Surface(screen, pixels, s.w, s.h)
		if err != nil {
			return nil, err
		}
		s.surf = surf
	}
	{
		tex, err := Surface2Texture(screen, s.surf)
		if err != nil {
			return nil, err
		}
		s.tex = tex
	}
	return s, nil
}

// DrawAt the Sprite onto the screen
func (s *Sprite) DrawAt(x, y, z int32, scale float32) {
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
