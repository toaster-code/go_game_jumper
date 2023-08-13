package tiles

import (
	"github.com/hajimehoshi/ebiten"
)

// Define your types, functions, and variables related to tiles here



// NewTile creates and initializes a new Tile instance with the image set to nil.
func NewTile(image *ebiten.Image, x, y, vx, vy, width, height int, standing, blocking bool) Tile {
	return Tile{
		Image:    image,
		X:        x,
		Y:        y,
		Vx:       vx,
		Vy:       vy,
		Width:    width,
		Height:   height,
		Standing: standing,
		Blocking: blocking,
	}
}

// Define the Tile struct
type Tile struct {
	Image         *ebiten.Image
	X, Y          int
	Vx, Vy        int
	Width, Height int
	Standing      bool
	Blocking      bool
}

