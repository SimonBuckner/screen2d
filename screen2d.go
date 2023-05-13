package screen2d

import (
	"fmt"
	"time"
	"unsafe"

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
	tex, _ := rend.CreateTexture(uint32(sdl.PIXELFORMAT_RGBA8888), int(sdl.TEXTUREACCESS_STATIC), 1, 1)
	tex.SetBlendMode(sdl.BlendMode(sdl.BLENDMODE_ADD))
	pixels := make([]byte, 4)
	pixels[0] = r
	pixels[1] = g
	pixels[2] = b
	pixels[3] = a

	tex.Update(nil, unsafe.Pointer(&pixels), 4)
	return tex
}

// RGBAPixels2Surface takes an array of RGBA pixel data and returns a Surface
func RGBAPixels2Surface(rgbaData []int, w, h int32) (*sdl.Surface, error) {
	if int32(len(rgbaData)) != w*h {
		return nil, fmt.Errorf("bitmap does not have the correct number of pixels for surface (%d: %d*%d", len(rgbaData), w, h)
	}

	surf, err := sdl.CreateRGBSurfaceWithFormat(0, w, h, 32, uint32(sdl.PIXELFORMAT_RGBA8888))
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

// RGBAPixels2Texture takes an array of RGBA pixel data and returns a Texture
func RGBAPixels2Texture(rend *sdl.Renderer, rgbaData []int, w, h int32) (*sdl.Texture, error) {
	if int32(len(rgbaData)) != w*h {
		return nil, fmt.Errorf("bitmap does not have the correct number of pixels for surface (%d: %d*%d", len(rgbaData), w, h)
	}

	surf, err := sdl.CreateRGBSurfaceWithFormat(0, w, h, 32, uint32(sdl.PIXELFORMAT_RGBA8888))
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
	tex, texErr := rend.CreateTextureFromSurface(surf)

	return tex, texErr
}

// RGBAPixels2Mask takes an array of RGBA pixel data and returns a collision mask
func RGBAPixels2Mask(rgbaData []int, w, h int32) ([]bool, error) {
	if int32(len(rgbaData)) != w*h {
		return nil, fmt.Errorf("bitmap does not have the correct number of pixels for surface (%d: %d*%d", len(rgbaData), w, h)
	}

	mask := make([]bool, w*h)

	for i := 0; i < len(rgbaData); i++ {
		mask[i] = !(rgbaData[i] == 0x00000000)
	}
	return mask, nil
}

// Surface2Texture takes a Surface returnes a Texxture
func Surface2Texture(rend *sdl.Renderer, surf *sdl.Surface) (*sdl.Texture, error) {
	return rend.CreateTextureFromSurface(surf)
}

// Box discribes the bounding corners of a Box
type Box struct {
	X1, Y1 int32
	X2, Y2 int32
	W, H   int32
}

// Hitter represents an item that can be checked to see if it has hit another item
type Hitter interface {
	GetBox() Box
	GetMask() []bool
}

// CheckBoxHit checks if any part of the two EntitState boxes overlap
func CheckBoxHit(entity1, entity2 Hitter) bool {

	r1 := entity1.GetBox()
	r2 := entity2.GetBox()

	// Too far left or right
	if r1.X1 > r2.X2 || r1.X2 < r2.X1 {
		return false
	}

	// Top high or low
	if r1.Y1 > r2.Y2 || r1.Y2 < r2.Y1 {
		return false
	}

	return true
}

// CheckBoxHitDebug checks if any part of the two EntitState boxes overlap
func CheckBoxHitDebug(entity1, entity2 Hitter) bool {

	r1 := entity1.GetBox()
	fmt.Printf("R1 - X1: %d   Y1: %d   X2: %d   Y2: %d\n", r1.X1, r1.Y1, r1.X2, r1.Y2)

	r2 := entity2.GetBox()
	fmt.Printf("R2 - X1: %d   Y1: %d   X2: %d   Y2: %d\n", r2.X1, r2.Y1, r2.X2, r2.Y2)

	fmt.Printf("L/R Check - ")
	// Too far left or right
	if r1.X1 > r2.X2 || r1.X2 < r2.X1 {
		fmt.Printf("Miss\n")
		return false
	}
	fmt.Printf(" On Target / H/L Check -- ")
	// Top high or low
	if r1.Y1 > r2.Y2 || r1.Y2 < r2.Y1 {
		fmt.Printf("Miss\n")
		return false
	}
	fmt.Printf("Hit\n")
	return true
}

// CheckPixelHit checks if any pixels in the two EntityStates overlap
func CheckPixelHit(entity1, entity2 Hitter) bool {
	b1 := entity1.GetBox()
	m1 := entity1.GetMask()

	b2 := entity2.GetBox()
	m2 := entity2.GetMask()

	i := int32(0)
	for i < int32(len(m1)) {
		px := b1.X1 + (i % b1.W)
		py := b1.Y1 + (i / b1.W)
		if px >= b2.X1 && px < b2.X2 && py >= b2.Y1 && py < b2.Y2 {
			x := px - b2.X1
			y := py - b2.Y1
			i2 := (y * b2.W) + x
			if m1[i] == true && m2[i2] == true {
				return true
			}
		}
		i++
	}
	return false
}

// Counter hold various runtie countners
type Counter struct {
	FPS          int
	fpsStart     time.Time
	fpsFrames    int
	FrameElapsed float32
	frameStart   time.Time
	min, max     float32
}

// Start the counters running
func (c *Counter) Start() {
	c.fpsStart = time.Now()
	c.frameStart = time.Now()
	c.fpsFrames = 0
}

// FrameStart indicates a rendering frame as started
func (c *Counter) FrameStart() {
	c.frameStart = time.Now()
}

// FrameEnd indicates the rendering frame has ended
func (c *Counter) FrameEnd() {
	c.FrameElapsed = float32(time.Since(c.frameStart).Seconds())
	c.fpsFrames++

	if time.Since(c.fpsStart).Seconds() > 1 {
		fmt.Printf("FPS - %d\n", c.fpsFrames)
		c.FPS = c.fpsFrames
		c.fpsStart = time.Now()
		c.fpsFrames = 0
	}
}
