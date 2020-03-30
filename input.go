package screen2d

import "github.com/veandco/go-sdl2/sdl"

// KBState represents the state of the keyboard
type KBState struct {
	state  []uint8
	state1 []uint8
}

// NewKBState fsceney
func NewKBState() *KBState {
	kb := &KBState{}

	kb.state = sdl.GetKeyboardState()
	kb.state1 = make([]uint8, len(kb.state))

	return kb
}

const (
	// KEYUP represents the Key Up state
	KEYUP uint8 = 0
	// KEYDOWN represents the Key Down or pressed state
	KEYDOWN uint8 = 1
)

// Refresh the keyboard state
func (kb *KBState) Refresh() {
	for i, v := range kb.state {
		kb.state1[i] = v
	}
	kb.state = sdl.GetKeyboardState()
}

// OnKeyDown returns true when the key state changes from up to down
func (kb *KBState) OnKeyDown(key uint8) bool {
	return kb.state[key] == KEYDOWN && kb.state1[key] == KEYUP
}

// OnKeyUp returns true when the key state changes from up to down
func (kb *KBState) OnKeyUp(key uint8) bool {
	return kb.state[key] == KEYUP && kb.state1[key] == KEYDOWN
}

// IsKeyDown returns true if the key state is down
func (kb *KBState) IsKeyDown(key uint8) bool {
	return kb.state[key] == KEYDOWN
}

// IsKeyUp returns true if the key state is up
func (kb *KBState) IsKeyUp(key uint8) bool {
	return kb.state[key] == KEYUP
}
