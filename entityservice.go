package screen2d

import "fmt"

// EntityService builds new Entities with preset default values
type EntityService struct {
	projectXYFunc ProjectXYFunc
	scale         float32
}

// NewEntityService returns a new EntityService
func NewEntityService() *EntityService {
	es := &EntityService{
		projectXYFunc: defaultProjectXY,
		scale:         1.0,
	}
	return es
}

// NewEntity returns a new Entity with the preset defaults
func (es *EntityService) NewEntity() *Entity {
	e := &Entity{
		X:             0,
		Y:             0,
		Z:             0,
		Scale:         es.scale,
		Visible:       true,
		projectXYFunc: es.projectXYFunc,
	}
	return e
}

// NewTextEntity returns a new TextEntity with the preset defaults
func (es *EntityService) NewTextEntity() *TextEntity {
	e := &TextEntity{
		X:             0,
		Y:             0,
		Z:             0,
		Scale:         es.scale,
		Visible:       true,
		projectXYFunc: es.projectXYFunc,
	}
	return e
}

// SetXYProjection sets the default XY projection
func (es *EntityService) SetXYProjection(projectXY ProjectXYFunc) {
	es.projectXYFunc = projectXY
}

// SetScale sets the default Scale
func (es *EntityService) SetScale(scale float32) {
	es.scale = scale
}

// ProjectXYFunc definces a function that projects of virtaul Game X/Y coords to Screen coords
type ProjectXYFunc func(x, y, scale float32) (tX, tY int32)

// defaultProjectXY is the default projection used for virtual Game to Screen X/Y coords
func defaultProjectXY(x, y, scale float32) (tX, tY int32) {
	fmt.Println("defaultProjectXY")
	return int32(x * scale), int32(y * scale)
}
