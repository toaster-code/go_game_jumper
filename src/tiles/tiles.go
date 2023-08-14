package tiles

import (
	"github.com/hajimehoshi/ebiten"
)

// Define your types, functions, and variables related to tiles here
// this file should create, load tiles, add methods to manipulate tiles, instatiate tiles,...

var Fabio int = 255

// Define the Tile struct
type Tile struct {
	Image         *ebiten.Image // Image of the tile
	X, Y          int // embbed the X, Y coordinates of the tile
	Vx, Vy        int // embbed the speed at X, Y coordinates of the tile
	Width, Height int // embbed the size of the tile
	Standing      bool // flag referring to the status of the tile.
	Blocking      bool
}

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
