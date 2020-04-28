package screen2d

// Entity is a item that can be displayed on the screen and tested for collision
type Entity struct {
	// Virtaul game position
	X, Y, Z float32
	Scale   float32
	Sprite  *Sprite
	// ProjectXYFunc ProjectXYFunc
	Visible       bool
	projectXYFunc ProjectXYFunc
}

// GetBox returns the Emntity Box
func (e *Entity) GetBox() Box {
	return Box{
		X1: int32(e.X),
		Y1: int32(e.Y),
		X2: int32(e.X) + e.Sprite.w,
		Y2: int32(e.Y) + e.Sprite.h,
		W:  e.Sprite.w,
		H:  e.Sprite.h,
	}
}

// GetMask returns the collision mask for the underlying sprite
func (e *Entity) GetMask() []bool {
	if e.Sprite != nil {
		return e.Sprite.mask
	}
	return nil
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
	if e.Sprite == nil || e.Visible == false {
		return
	}
	x, y := e.projectXYFunc(e.X, e.Y, e.Scale)
	e.Sprite.DrawAt(x, y, e.Scale)
}

// SetProjectXYFunc overrides the default virutal game to screen coord calculation
func (e *Entity) SetProjectXYFunc(f ProjectXYFunc) {
	e.projectXYFunc = f
}

// ClearProjectXYFunc restores the default virutal game to screen coord calculation
func (e *Entity) ClearProjectXYFunc() {
	e.projectXYFunc = defaultProjectXY
}
