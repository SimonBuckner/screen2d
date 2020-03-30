package screen2d

// Atlas contains a list of Sprites that can be retrieved using an AtlasKey
type Atlas struct {
	sprites map[AtlasKey]*Sprite
}

// AtlasKey is used to retrieve a specific Sprite from the Atlas
type AtlasKey int

// NewAtlas returns an empty Atlas
func NewAtlas() *Atlas {
	a := &Atlas{
		sprites: make(map[AtlasKey]*Sprite),
	}
	return a
}

// AddSprite adds a Sprite to the Atlas against the specified AtlasKey
func (a *Atlas) AddSprite(key AtlasKey, sprite *Sprite) {
	a.sprites[key] = sprite
}

// GetSprite gets the Sprite at the Atlas position specified by the AtlasKey
func (a *Atlas) GetSprite(key AtlasKey) *Sprite {
	return a.sprites[key]
}
