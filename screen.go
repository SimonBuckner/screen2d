package screen2d

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Drawer represents an item that can be drawn on the screen
type Drawer interface {
	Draw()
}

// UpdateFunc is the function that will be called during the update phase
type UpdateFunc func(ticks uint32)

// DrawFunc is the function that will be called during the draw phase
type DrawFunc func()

// Screen represents a 2d window as shown on the screen
type Screen struct {
	top        int32
	left       int32
	width      int32
	height     int32
	title      string
	elapsed    float32
	rend       *sdl.Renderer
	wind       *sdl.Window
	keyb       *KBState
	updateFunc UpdateFunc
	drawFunc   DrawFunc
}

// NewScreen returns a newly initialisd screen in Windowed mode
func NewScreen(width, height int, title string) (*Screen, error) {
	s := &Screen{
		// top:    int32(top),
		// left:   int32(left),
		width:      int32(width),
		height:     int32(height),
		title:      title,
		keyb:       NewKBState(),
		drawFunc:   drawStub,
		updateFunc: updateStub,
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "0")

	{
		wind, err := sdl.CreateWindow(title, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, s.width, s.height, sdl.WINDOW_SHOWN)
		if err != nil {
			return nil, err
		}
		s.wind = wind
	}

	{
		rend, err := sdl.CreateRenderer(s.wind, -1, sdl.RENDERER_ACCELERATED)
		if err != nil {
			return nil, err
		}
		s.rend = rend
	}

	return s, nil
}

// Destroy cleans up resources
func (s *Screen) Destroy() {
	if s.rend != nil {
		s.rend.Destroy()
	}

	if s.wind != nil {
		s.wind.Destroy()
	}
}

// SetBackgroundColor sets the background color drawn on the screen
func (s *Screen) SetBackgroundColor(r, b, g, a uint8) {
	s.rend.SetDrawColorArray(r, g, b, a)
}

// Rend returns the SDL Renderer
func (s *Screen) Rend() *sdl.Renderer {
	return s.rend
}

// SetUpdateFunc sets the function that will be called during the update phase
func (s *Screen) SetUpdateFunc(f UpdateFunc) {
	s.updateFunc = f
}

// SetDrawFunc sets the function that will be called during the draw phase
func (s *Screen) SetDrawFunc(f DrawFunc) {
	s.drawFunc = f
}

// Run starts the main event loop
func (s *Screen) Run() {
	// s.FPSCount = 0
	// s.FPS = 0
	// s.FPSStart = time.Now()

	var loopStart time.Time

	for {
		loopStart = time.Now()
		// if s.Stopping {
		// 	return
		// }

		s.keyb.Refresh()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_CLOSE {
					return
				}
				// case *sdl.KeyboardEvent:
				// 	s.FireKeyboardEvent(e)
			}
		}

		s.rend.Clear()
		s.updateFunc(sdl.GetTicks())
		s.drawFunc()
		s.rend.Present()
		// sdl.Delay(1)
		// s.loopStart = float32(time.Since(frameStart).Seconds())
		// s.FPSCount++
		// if time.Since(s.FPSStart) > time.Second {
		// 	s.FPS = s.FPSCount
		// 	s.FPSStart = time.Now()
		// 	s.FPSCount = 0
		// 	fmt.Printf("Frames per second : %d\n", s.FPS)
		// }
		s.elapsed = float32(time.Since(loopStart).Seconds())
	}
}

func updateStub(ticks uint32) {
	sdl.Delay(16)
}

func drawStub() {

}
