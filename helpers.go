package screen2d

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Color meets the color.Color interface required for CreateRGBSurface
type Color struct {
	R, G, B, A uint32
}

// RGBA meets the color.Color interface required for CreateRGBSurface
func (c Color) RGBA() (r, g, b, a uint32) {
	return c.R, c.G, c.B, c.A
}

// HexColorToRGBA converts a colour stored in an int to RGBA values
func HexColorToRGBA(color int) *Color {
	return &Color{
		R: uint32((color & 0xFF000000) >> 24),
		G: uint32((color & 0x00FF0000) >> 16),
		B: uint32((color & 0x0000FF00) >> 8),
		A: uint32(color & 0x000000FF),
	}
}

// new1PTexture returns a new texture comprising a single pixel
func new1PTexture(rend *sdl.Renderer, r, g, b, a uint8) *sdl.Texture {
	tex, _ := rend.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, 1, 1)
	tex.SetBlendMode(sdl.BLENDMODE_ADD)
	pixels := make([]byte, 4)
	pixels[0] = r
	pixels[1] = g
	pixels[2] = b
	pixels[3] = a

	tex.Update(nil, pixels, 4)
	return tex
}

// RGBAPixels2Surface takes an array of RGBA pixel data and returns a Surface
func RGBAPixels2Surface(rgbaData []int, w, h int32) (*sdl.Surface, error) {
	if int32(len(rgbaData)) != w*h {
		return nil, fmt.Errorf("bitmap does not have the correct number of pixels for surface (%d: %d*%d", len(rgbaData), w, h)
	}

	surf, err := sdl.CreateRGBSurfaceWithFormat(0, w, h, 32, sdl.PIXELFORMAT_RGBA8888)
	if err != nil {
		return nil, err
	}

	pixels := surf.Pixels()
	j := 0
	for i := 0; i < len(rgbaData); i++ {
		j = i * 4
		pixels[j+0] = byte((rgbaData[i] & 0xFF000000) >> 24)
		pixels[j+1] = byte((rgbaData[i] & 0x00FF0000) >> 16)
		pixels[j+2] = byte((rgbaData[i] & 0x0000FF00) >> 8)
		pixels[j+3] = byte((rgbaData[i] & 0x000000FF))
	}
	return surf, nil
}

// Surface2Texture takes a Surface returnes a Texxture
func Surface2Texture(rend *sdl.Renderer, surf *sdl.Surface) (*sdl.Texture, error) {
	return rend.CreateTextureFromSurface(surf)
}

// Box discribes the bounding corners of a Box
type Box struct {
	X1, Y1 int32
	X2, Y2 int32
}

// // Hitter represents an item that can be checked to see if it has hit another item
// type Hitter interface {
// 	GetBox() Box
// 	GetPitch() int32
// 	GetPixels() []byte
// }

// // CheckBoxHit checks if any part of the two EntitState boxes overlap
// func CheckBoxHit(entity1, entity2 Hitter) bool {

// 	r1 := entity1.GetBox()
// 	r2 := entity2.GetBox()

// 	// Too far left or right
// 	if r1.X1 > r2.X2 || r1.X2 < r2.X2 {
// 		return false
// 	}

// 	// Top high or low
// 	if r1.Y1 > r2.Y2 || r1.Y2 < r2.Y1 {
// 		return false
// 	}

// 	return true
// }

// // CheckPixelHit checks if any pixels in the two EntityStates overlap
// func CheckPixelHit(entity1, entity2 Hitter, miss *sdl.Color) bool {

// 	b1 := entity1.GetBox()
// 	p1 := entity1.GetPitch()

// 	pixels1 := entity1.GetPixels()

// 	b2 := entity2.GetBox()
// 	p2 := entity1.GetPitch()
// 	pixels2 := entity2.GetPixels()

// 	i := int32(0)
// 	for i < int32(len(pixels1)) {
// 		px := b1.X1 + (i % p1)
// 		py := b1.Y1 + (i / p1)
// 		if px >= b2.X1 && px < b2.X2 && py >= b2.Y1 && py < b2.Y2 {
// 			x := px - b2.X1
// 			y := py - b2.Y1
// 			i2 := (y * p2) + (x * 4)
// 			if pixels1[i] != miss.R && pixels2[i2] != miss.R {
// 				return true
// 			}
// 			i++
// 			if pixels1[i] != miss.G && pixels2[i2] != miss.G {
// 				return true
// 			}
// 			i++
// 			if pixels1[i] != miss.B && pixels2[i2] != miss.B {
// 				return true
// 			}
// 			i++
// 			// No point checking alpha
// 			i++
// 		}
// 	}
// 	return false
// }
