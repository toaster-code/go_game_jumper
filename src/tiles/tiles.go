package tiles

import (
	img "image"

	"github.com/hajimehoshi/ebiten"
)

// Define your types, functions, and variables related to tiles here
// this file should create, load tiles, add methods to manipulate tiles, instatiate tiles,...

var Check_if_working int = 255

// Define the Tile struct
type Tile struct {
	SpriteSheetImage *ebiten.Image // pointer to the image of the tile
	X, Y             int           // embbed the X, Y coordinates of the tile
	Vx, Vy           int           // embbed the speed at X, Y coordinates of the tile
	Width, Height    int           // embbed the size of the tile
	Standing         bool          // flag referring to the status of the tile.
	Blocking         bool          // flag referring to the status of the tile.
}

// NewTile creates and initializes a new Tile instance with the image set to nil.
func NewTile(image *ebiten.Image, x int, y int, vx int, vy int, width int, height int, standing bool, blocking bool) Tile {
	return Tile{
		SpriteSheetImage: image,
		X:                x,
		Y:                y,
		Vx:               vx,
		Vy:               vy,
		Width:            width,
		Height:           height,
		Standing:         standing,
		Blocking:         blocking,
	}
}

// creates a slice of tiles
func NewTileSlice(size int) []Tile {
	return make([]Tile, size)
}

func (t *Tile) numSprites() int {
	return countPixels(t.SpriteSheetImage)/256
}

func countPixels(img *ebiten.Image) int {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	return width * height
}

func (t *Tile) SpriteByIndex(index int) *ebiten.Image {
	// size := t.SpriteSheetImage.Bounds().Size()
	// size.X
	return SpriteByIndex(t.SpriteSheetImage, index, t.Width, t.Height)
}

func SpriteByIndex(image *ebiten.Image, index int, tileWidth, tileHeight int) *ebiten.Image {
	// Define the coordinates for the subimage
	minX := index * tileWidth
	minY := 0
	maxX := minX + tileWidth
	maxY := minY + tileHeight
	return image.SubImage(img.Rect(minX, minY, maxX, maxY)).(*ebiten.Image)
}
