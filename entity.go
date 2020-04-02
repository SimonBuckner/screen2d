package screen2d

// CalcScreenXYFunc allows for custom conversion of virtaul Game X/Y coords to Screen coords
type CalcScreenXYFunc func(x, y, scale float32) (tX, tY int32)

// Entity is a item that can be displayed on the screen and tested for collision
type Entity struct {
	// Virtaul game position
	X, Y, Z          float32
	Scale            float32
	Sprite           *Sprite
	calcScreenXYFunc CalcScreenXYFunc
}

// NewEntity returns a new Entity
func NewEntity() *Entity {
	e := &Entity{
		X:     0,
		Y:     0,
		Z:     0,
		Scale: 1.0,
	}
	return e
}

// NewEntityWithTransform returns a new Entity
func NewEntityWithTransform() *Entity {
	e := &Entity{
		X:                0,
		Y:                0,
		Z:                0,
		Scale:            1.0,
		calcScreenXYFunc: calcScreenXY,
	}
	return e
}

// GetBox returns the Emntity Box
func (e *Entity) GetBox() Box {
	return Box{
		X1: int32(e.X),
		Y1: int32(e.Y),
		X2: int32(e.Sprite.w),
		Y2: int32(e.Sprite.h),
	}
}

// SetPos sets the position of the Entity
func (e *Entity) SetPos(x, y, z int32) {
	e.X = float32(x)
	e.Y = float32(y)
	e.Z = float32(z)
}

// SetSprite sets the Sprite the Entity will display
func (e *Entity) SetSprite(sprite *Sprite) {
	e.Sprite = sprite
}

// Draw is the default draw method
func (e *Entity) Draw() {
	if e.Sprite == nil {
		return
	}
	x, y := e.calcScreenXYFunc(e.X, e.Y, e.Scale)
	e.Sprite.DrawAt(x, y, e.Scale)
}

// SetCalcScreenXYFunc overrides the default virutal game to screen coord calculation
func (e *Entity) SetCalcScreenXYFunc(f CalcScreenXYFunc) {
	e.calcScreenXYFunc = f
}

// ClearCalcScreenXYFunc restores the default virutal game to screen coord calculation
func (e *Entity) ClearCalcScreenXYFunc() {
	e.calcScreenXYFunc = calcScreenXY
}

func calcScreenXY(x, y, scale float32) (tX, tY int32) {
	return int32(x * scale), int32(y * scale)
}
