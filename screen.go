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
type UpdateFunc func(ticks uint32, elapsed float32)

// DrawFunc is the function that will be called during the draw phase
type DrawFunc func()

// KeyboardEventFunc is the function that will be called when a keyboard event occurs
type KeyboardEventFunc func(e *sdl.KeyboardEvent)

// Screen represents a 2d window as shown on the screen
type Screen struct {
	top         int32
	left        int32
	width       int32
	height      int32
	title       string
	close       bool
	rend        *sdl.Renderer
	wind        *sdl.Window
	keyb        *KBState
	updateFunc  UpdateFunc
	drawFunc    DrawFunc
	keyDownFunc KeyboardEventFunc
	keyUpFunc   KeyboardEventFunc
}

// NewScreen returns a newly initialisd screen in Windowed mode
func NewScreen(width, height int, title string) (*Screen, error) {
	s := &Screen{
		width:       int32(width),
		height:      int32(height),
		title:       title,
		close:       false,
		keyb:        NewKBState(),
		drawFunc:    drawStub,
		updateFunc:  updateStub,
		keyDownFunc: keyEventStub,
		keyUpFunc:   keyEventStub,
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

// SetKeyDownFunc sets the function that will be called whena key is pressed
func (s *Screen) SetKeyDownFunc(f KeyboardEventFunc) {
	s.keyDownFunc = f
}

// SetKeyUpFunc sets the function that will be called when a key is released
func (s *Screen) SetKeyUpFunc(f KeyboardEventFunc) {
	s.keyUpFunc = f
}

// Run starts the main event loop
func (s *Screen) Run() {
	var loopStart time.Time
	var elapsed float32

	for {
		loopStart = time.Now()
		if s.close {
			return
		}
		s.keyb.Refresh()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_CLOSE {
					return
				}
			case *sdl.KeyboardEvent:
				s.despatchKeyboardEvent(e)
			}
		}

		s.rend.Clear()
		s.updateFunc(sdl.GetTicks(), elapsed)
		s.drawFunc()
		s.rend.Present()

		sdl.Delay(1)
		elapsed = float32(time.Since(loopStart).Seconds())
	}
}

// Close causes the Screen to close
func (s *Screen) Close() {
	s.close = true
}

func (s *Screen) despatchKeyboardEvent(e *sdl.KeyboardEvent) {
	if e.Repeat != 0 {
		return
	}
	switch e.Type {
	case sdl.KEYDOWN:
		s.keyDownFunc(e)
	case sdl.KEYUP:
		s.keyUpFunc(e)
	}
}

// Default stub functions to reduce number of nil tests
func updateStub(ticks uint32, elapsed float32) {}
func drawStub()                                {}
func keyEventStub(e *sdl.KeyboardEvent)        {}

// ClearUpdateFunc clears the function that will be called during the update phase
func (s *Screen) ClearUpdateFunc() {
	s.updateFunc = updateStub
}

// ClearDrawFunc clears the function that will be called during the draw phase
func (s *Screen) ClearDrawFunc() {
	s.drawFunc = drawStub
}

// ClearKeyDownFunc clears the function that will be called whena key is pressed
func (s *Screen) ClearKeyDownFunc() {
	s.keyDownFunc = keyEventStub
}

// ClearKeyUpFunc clears the function that will be called when a key is released
func (s *Screen) ClearKeyUpFunc() {
	s.keyUpFunc = keyEventStub
}

// ClearFuncs clears all update functions
func (s *Screen) ClearFuncs() {
	s.updateFunc = updateStub
	s.drawFunc = drawStub
	s.keyDownFunc = keyEventStub
	s.keyUpFunc = keyEventStub
}
