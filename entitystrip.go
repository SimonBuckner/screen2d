package screen2d

// StripDirection represents the direction in which the Sprites will be drawn.
type StripDirection int

const (
	// Left2RightStrip draws from left to right
	Left2RightStrip StripDirection = iota
	// Top2BottomStrip draws from top to bootm
	Top2BottomStrip
)

// EntityStrip is a strip of items that can be displayed on the screen and tested for collision
type EntityStrip struct {
	// Virtaul game position
	X, Y, Z   float32
	Scale     float32
	Sprites   []*Sprite
	size      int
	direction StripDirection
	// ProjectXYFunc ProjectXYFunc
	projectXYFunc ProjectXYFunc
}

// NewEntityStrip returns a new EntityStrip
// func NewEntityStrip(size int, direction StripDirection) *EntityStrip {
// 	e := &EntityStrip{
// 		X:                0,
// 		Y:                0,
// 		Z:                0,
// 		Scale:            1.0,
// 		size:             size,
// 		direction:        direction,
// 		Sprites:          make([]*Sprite, size),
// 		ProjectXYFunc: scaleScreenXY,
// 	}
// 	return e
// }

// GetBox returns the Entity Box
func (e *EntityStrip) GetBox() Box {
	return Box{
		X1: int32(e.X),
		Y1: int32(e.Y),
		X2: (e.Sprites[0].w * int32(e.size)) + int32(e.X),
		Y2: int32(e.Sprites[0].h),
	}
}

// SetPos sets the position of the Entity
func (e *EntityStrip) SetPos(x, y, z int32) {
	e.X = float32(x)
	e.Y = float32(y)
	e.Z = float32(z)
}

// SetSprite sets the Sprite the Entity will display
func (e *EntityStrip) SetSprite(index int, sprite *Sprite) {
	if index < e.size {
		e.Sprites[index] = sprite
	}
}

// Draw is the default draw method
func (e *EntityStrip) Draw() {
	if e.Sprites == nil {
		return
	}

	x, y := e.projectXYFunc(e.X, e.Y, e.Scale)
	for _, s := range e.Sprites {
		if s != nil {
			s.DrawAt(x, y, e.Scale)
			x += s.w
		}
	}
}

// SetProjectXYFunc overrides the default virutal game to screen coord calculation
func (e *EntityStrip) SetProjectXYFunc(f ProjectXYFunc) {
	e.projectXYFunc = f
}

// ClearProjectXYFunc restores the default virutal game to screen coord calculation
func (e *EntityStrip) ClearProjectXYFunc() {
	e.projectXYFunc = defaultProjectXY
}
