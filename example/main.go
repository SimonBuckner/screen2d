package main

import (
	"github.com/SimonBuckner/screen2d"
)

type game struct {
	screen         *screen2d.Screen
	atlas          *screen2d.Atlas
	e1, e2, e3, e4 *screen2d.Entity
}

const (
	keyAlienSprCYA screen2d.AtlasKey = iota
	keyAlienSprCYB
)

func main() {
	s, err := screen2d.NewScreen(1024, 768, "Example Screen2D")
	if err != nil {
		panic(err)
	}
	defer s.Destroy()

	s.SetBackgroundColor(0, 0, 0, 0)

	g := &game{
		screen: s,
		atlas:  screen2d.NewAtlas(),
		e1:     screen2d.NewEntity(),
		e2:     screen2d.NewEntity(),
		e3:     screen2d.NewEntity(),
		e4:     screen2d.NewEntity(),
	}

	{
		s, err := screen2d.NewSpriteFromnRGBAPixels(s, alienSprCYA.Pixels, int32(alienSprCYA.Pitch))
		if err != nil {
			panic(err)
		}
		g.atlas.AddSprite(keyAlienSprCYA, s)
	}

	{
		s, err := screen2d.NewSpriteFromnRGBAPixels(s, alienSprCYB.Pixels, int32(alienSprCYB.Pitch))
		if err != nil {
			panic(err)
		}
		g.atlas.AddSprite(keyAlienSprCYB, s)
	}

	s.SetUpdateFunc(g.update)
	s.SetDrawFunc(g.draw)

	g.e1.Sprite = g.atlas.GetSprite(keyAlienSprCYA)
	g.e2.Sprite = g.atlas.GetSprite(keyAlienSprCYA)
	g.e3.Sprite = g.atlas.GetSprite(keyAlienSprCYB)
	g.e4.Sprite = g.atlas.GetSprite(keyAlienSprCYB)

	g.e1.Scale = 2.0
	g.e2.Scale = 5.0
	g.e3.Scale = 1.0
	g.e4.Scale = 1.5

	s.Run()
}

func (g *game) update(ticks uint32) {
	g.e1.SetPos(100, 100, 0)
	g.e2.SetPos(200, 200, 0)
	g.e3.SetPos(300, 300, 0)
	g.e4.SetPos(400, 400, 0)
}

func (g *game) draw() {
	g.e1.Draw()
	g.e2.Draw()
	g.e3.Draw()
	g.e4.Draw()
}
