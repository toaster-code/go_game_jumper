package tiles

import (
	img "image"

	"github.com/hajimehoshi/ebiten"
)

// Define your types, functions, and variables related to tiles here
// this file should create, load tiles, add methods to manipulate tiles, instatiate tiles,...

// Define the Tile struct
type Tile struct {
	SpriteSheetImage *ebiten.Image // pointer to the image of the sprite sheet
	X, Y             int           // embbed the X, Y coordinates of the tile
	Vx, Vy           int           // embbed the speed at X, Y coordinates of the tile
	Width, Height    int           // embbed the size of the tile
	Standing         bool          // flag referring to the status of the tile
	Blocking         bool          // flag referring to the status of the tile
	SpriteIndex      int           // index of the sprite in the sprite sheet
}

// NewTile creates and initializes a new Tile instance with the image set to nil.
func NewTile(img *ebiten.Image, x int, y int, vx int, vy int, width int, height int, standing bool, blocking bool) Tile {
	return Tile{
		SpriteSheetImage: img,
		X:           x,
		Y:           y,
		Vx:          vx,
		Vy:          vy,
		Width:       width,
		Height:      height,
		Standing:    standing,
		Blocking:    blocking,
	}
}

func (t *Tile) SpriteSheet() SpriteSheet {

	return
}

// NewTileSlice creates a slice of tiles
func NewTileSlice(size int) []Tile {
	return make([]Tile, size)
}

type SpriteSheet struct {
	Image *ebiten.Image
	TileWidth, TileHeight int
}

// Check if the sprite sheet is valid by checking its dimensions
func (t *SpriteSheet) isValid() bool {
	// the size must be 
}

// Initialize a sprite sheet from an image file
func LoadImage(imagePath string, tileWidth int, tileHeight int) (*ebiten.Image, error) {
	return ebiten.NewImageFromFile(imagePath)
}

func (t *Tile) numSprites() int {
	return countPixels(t.Image) / 256
}

func countPixels(img *ebiten.Image) int {
	bounds := img.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.Y
	return width * height
}

func (t *Tile) SpriteByIndex(index int) *ebiten.Image {
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
