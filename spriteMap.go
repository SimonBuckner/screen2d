package screen2d

// SpriteMap contains a map of individual Sprites that can be retrieved using an SpriteMapKey
type SpriteMap struct {
	sprites map[SpriteMapKey]*Sprite
}

// SpriteMapKey is used to retrieve a specific Sprite from the SpriteMap
type SpriteMapKey int

// NewSpriteMap returns an empty SpriteMap
func NewSpriteMap() *SpriteMap {
	a := &SpriteMap{
		sprites: make(map[SpriteMapKey]*Sprite),
	}
	return a
}

// AddSprite adds a Sprite to the SpriteMap against the specified SpriteMapKey
func (a *SpriteMap) AddSprite(key SpriteMapKey, sprite *Sprite) {
	a.sprites[key] = sprite
}

// GetSprite gets the Sprite at the SpriteMap position specified by the SpriteMapKey
func (a *SpriteMap) GetSprite(key SpriteMapKey) *Sprite {
	// TODO Clean up FMT
	// st := time.Now
	s := a.sprites[key]
	// rt := time.Since(st)
	// fmt.Printf("RT %v\n", rt)
	return s
}
