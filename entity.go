package screen2d

// Entity is a item that can be displayed on the screen and tested for collision
type Entity struct {
	// Virtaul game position
	X, Y, Z int32
	Scale   float32
	Sprite  *Sprite
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
		X:     0,
		Y:     0,
		Z:     0,
		Scale: 1.0,
	}
	return e
}

// GetBox returns the Emntity Box
func (e *Entity) GetBox() Box {
	return Box{
		X1: e.X,
		Y1: e.Y,
		X2: int32(e.Sprite.w),
		Y2: int32(e.Sprite.h),
	}
}

// Draw the Entity
func (e *Entity) Draw() {
	e.Sprite.Draw()
}

// SetPos sets the position of the Entity
func (e *Entity) SetPos(x, y, z int32) {
	e.X = x
	e.Y = y
	e.Z = z
}

// SetSprite sets the Sprite the Entity will display
func (e *Entity) SetSprite(sprite *Sprite) {
	e.Sprite = sprite
}