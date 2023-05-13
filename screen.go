package screen2d

import (
	"fmt"

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
	counter     *Counter
	updateFunc  UpdateFunc
	drawFunc    DrawFunc
	keyDownFunc KeyboardEventFunc
	keyUpFunc   KeyboardEventFunc
}

func init() {
	fmt.Println("init")
	err := sdl.Init(uint32(sdl.INIT_EVERYTHING))
	if err != nil {
		panic(err)
	}
}

// NewScreen returns a newly initialisd screen in Windowed mode
func NewScreen(width, height int, title string, hints ...ScreenHint) (*Screen, error) {
	s := &Screen{
		width:       int32(width),
		height:      int32(height),
		title:       title,
		close:       false,
		keyb:        NewKBState(),
		counter:     &Counter{},
		drawFunc:    drawStub,
		updateFunc:  updateStub,
		keyDownFunc: keyEventStub,
		keyUpFunc:   keyEventStub,
	}

	for _, hint := range hints {
		hint(s)
	}

	{
		wind, err := sdl.CreateWindow(title, int32(sdl.WINDOWPOS_CENTERED), int32(sdl.WINDOWPOS_CENTERED), s.width, s.height, uint32(sdl.WINDOW_SHOWN))
		if err != nil {
			return nil, err
		}
		s.wind = wind
	}

	{
		rend, err := sdl.CreateRenderer(s.wind, -1, uint32(sdl.RENDERER_ACCELERATED))
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
	s.counter.Start()
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
	s.counter.Start()
	for {
		s.counter.FrameStart()
		if s.close {
			return
		}

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
		s.updateFunc(sdl.GetTicks(), s.counter.FrameElapsed)
		s.drawFunc()
		s.rend.Present()
		s.keyb.Refresh()
		s.counter.FrameEnd()
	}
}

// Close causes the Screen to close
func (s *Screen) Close() {
	s.close = true
}

func (s *Screen) despatchKeyboardEvent(e *sdl.KeyboardEvent) {
	// if e.Repeat != 0 {
	// 	return
	// }
	switch e.Type {
	case sdl.KEYDOWN:
		s.keyDownFunc(e)
	case sdl.KEYUP:
		s.keyUpFunc(e)
	}
}

// Default stub functions to reduce number of nil tests
func drawStub()                                {}
func updateStub(ticks uint32, elapsed float32) {}
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

// GetKBState allows access to the current state of the keyboard
func (s *Screen) GetKBState() *KBState {
	return s.keyb
}

// Screen Configuration

// ScreenHint represtents configuration options when creating a new screen
type ScreenHint func(*Screen)

// SetVSync enables or disables virtical sync
func SetVSync(enabled bool) func(*Screen) {
	return func(s *Screen) {
		if enabled {
			sdl.SetHint(string(sdl.HINT_RENDER_VSYNC), "1")
		}
	}
}

// ScreenScalingQuality represent the scaling method
type ScreenScalingQuality int

const (
	// ScreenScalingNearestPixel - nearest pixel sampling
	ScreenScalingNearestPixel ScreenScalingQuality = iota
	// ScreenScalingLinear - linear filtering (supported by OpenGL and Direct3D)
	ScreenScalingLinear
	// ScreenScalingAnistropic - anisotropic filtering (supported by Direct3D)
	ScreenScalingAnistropic
)

// SetScalingQuality sets the scaling quality of the screen
func SetScalingQuality(quality ScreenScalingQuality) func(*Screen) {
	return func(s *Screen) {
		switch quality {
		case ScreenScalingLinear:
			sdl.SetHint(string(sdl.HINT_RENDER_SCALE_QUALITY), "1")
		case ScreenScalingAnistropic:
			sdl.SetHint(string(sdl.HINT_RENDER_SCALE_QUALITY), "2")
		default:
			sdl.SetHint(string(sdl.HINT_RENDER_SCALE_QUALITY), "0")
		}
	}
}
