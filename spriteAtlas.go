package screen2d

import (
	"github.com/veandco/go-sdl2/sdl"
)

// SpriteAtlas comprises multiple images tiled into a larger image. These images can be retrieved via their coordinates
type SpriteAtlas struct {
	pitch      int32
	tileWidth  int32
	tileHeight int32
	rend       *sdl.Renderer
	surf       *sdl.Surface
	tex        *sdl.Texture
}

// NewSpriteAtlas returns a new SpriteAtlas
func NewSpriteAtlas(rend *sdl.Renderer) *SpriteAtlas {
	sa := &SpriteAtlas{
		rend:       rend,
		pitch:      0,
		tileWidth:  0,
		tileHeight: 0,
	}
	return sa
}

// LoadRGBAPixels loads the Sprite using an array of 32-bit pixels containing RGBA values
func (sa *SpriteAtlas) LoadRGBAPixels(pixels []int, pixelsPitch, tilePitch, tileHeight int32) error {
	w := pixelsPitch
	h := int32(len(pixels)) / w
	sa.tileWidth = tilePitch
	sa.tileHeight = tileHeight
	sa.pitch = pixelsPitch * 4

	{
		surf, err := RGBAPixels2Surface(pixels, w, h)
		if err != nil {
			return err
		}
		sa.surf = surf
	}
	{
		tex, err := Surface2Texture(sa.rend, sa.surf)
		if err != nil {
			return err
		}
		sa.tex = tex
	}
	return nil
}

// DrawTileAt draws a specific tile at the given location
func (sa *SpriteAtlas) DrawTileAt(tileX, tileY, screenX, screenY int32, scale float32) {
	if sa.tex == nil {
		return
	}
	srcRect := &sdl.Rect{
		X: sa.tileWidth * tileX,
		Y: sa.tileHeight * tileY,
		W: sa.tileWidth,
		H: sa.tileHeight,
	}

	dstRect := &sdl.Rect{
		X: screenX,
		Y: screenY,
		W: int32(float32(sa.tileWidth) * scale),
		H: int32(float32(sa.tileHeight) * scale),
	}

	sa.rend.Copy(sa.tex, srcRect, dstRect)
}
