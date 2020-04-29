package screen2d

// TextEntity is a item that can be displayed on the screen and tested for collision
type TextEntity struct {
	// Virtaul game position
	X, Y, Z float32
	Scale   float32
	atlas   *SpriteAtlas
	keys    map[rune]int32
	// ProjectXYFunc ProjectXYFunc
	Visible       bool
	projectXYFunc ProjectXYFunc
	text          string
	fontWidth     int32
}

// SetPos sets the position of the Entity
func (e *TextEntity) SetPos(x, y, z int32) {
	e.X = float32(x)
	e.Y = float32(y)
	e.Z = float32(z)
}

// LoadAtlas sets the Sprite the Entity will display
func (e *TextEntity) LoadAtlas(atlas *SpriteAtlas, keys map[rune]int32, fontWidth int32) {
	e.atlas = atlas
	e.keys = keys
	e.fontWidth = fontWidth
}

// Draw is the default draw method
func (e *TextEntity) Draw() {
	if e.atlas == nil || e.Visible == false {
		return
	}

	x, y := e.X, e.Y
	for _, r := range e.text {
		if tileY, ok := e.keys[r]; ok {
			newX, newY := e.projectXYFunc(x, y, e.Scale)
			e.atlas.DrawTileAt(0, tileY, newX, newY, e.Scale)
		}
		x += float32(e.fontWidth)
	}
}

// SetProjectXYFunc overrides the default virutal game to screen coord calculation
func (e *TextEntity) SetProjectXYFunc(f ProjectXYFunc) {
	e.projectXYFunc = f
}

// ClearProjectXYFunc restores the default virutal game to screen coord calculation
func (e *TextEntity) ClearProjectXYFunc() {
	e.projectXYFunc = defaultProjectXY
}

// SetText sets the text that will be displayed
func (e *TextEntity) SetText(value string) {
	e.text = value
}
