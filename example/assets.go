package main

// Bitmap reqpresents a basic bitmap
type Bitmap struct {
	Pitch  int
	Pixels []int
}

var alienSprCYA = &Bitmap{
	Pitch: 16,
	Pixels: []int{
		0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0x0,
		0x0, 0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0,
		0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0x0,
		0x0, 0x0, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0x0, 0x0,
		0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0,
		0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0,
	},
}

var alienSprCYB = &Bitmap{
	Pitch: 16,
	Pixels: []int{
		0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0,
		0x0, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0x0,
		0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0,
		0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0x0, 0x0,
		0x0, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0x0, 0x0, 0xFFFFFFFF, 0x0, 0xFFFFFFFF, 0x0,
	},
}
