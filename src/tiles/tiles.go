package tiles

import (
	// "fmt"
	"github.com/hajimehoshi/ebiten"
	img "image"
)

// Define your types, functions, and variables related to tiles here
// this file should create, load tiles, add methods to manipulate tiles, instatiate tiles,...

// Define the Tile struct
type Tile struct {
	Id             int                    // id of the tile
	Coordinates    Point                  // coordinates of the tile
	Actions        Actions                // movement of the tile
	Width, Height  int                    // size of the tile
	Sprite         Sprite                 // sprite of the tile
	Type           int                    // type of the tile: 0 = background, 1 = foreground, 2 = player, 3 = enemy, 4 = item, 5 = projectile
	Hitbox         Hitbox                 // hitbox of the tile
	Visible        bool                   // flag referring to the status of the tile
	Acessible      bool                   // flag indicating if the tile is acessible by pathfinding
	ResourceType   string                 // type of resource, if applicable
	Status         string                 // status of the tile: "idle", "moving", "attacking", "dead", "on fire", "frozen"
	TriggerEvent   string                 // Identifier for trigger or event associated with the tile
	InteractionPts []Point                // List of interaction points on the tile
	Durability     int                    // Durability or health of the tile
	Owner          string                 // Owner or affiliation of the tile
	CustomData     map[string]interface{} // Additional custom data
}

type Actions struct {
	Vx, Vy   int
	Standing bool // flag referring to the status of the tile
	Blocking bool // flag referring to the status of the tile
}

type Hitbox struct {
	OffsetX, OffsetY int
	Width, Height    int
}

type Point struct {
	X, Y int
}

// Define the Texture struct
type Sprite struct {
	ImageSheet *ebiten.Image // pointer to the image of the sprite sheet
	Index      int
	Variation  int // variation of the tile appearance
	Rotation   int // rotation of the clockwise of a tile: 0 = 0 degrees, 1 = 90 degrees, 2 = 180 degrees, 3 = 270 degrees
	Mirror     int // mirror of the tile: 0 = no mirror, 1 = horizontal mirror, 2 = vertical mirror, 3 = horizontal and vertical mirror

}

// Image returns the ebiten.Image of the sprite.
// SpriteByIndex with the sprite's ImageSheet, Index, and dimensions of the tile (default 16x16).
func (s *Sprite) Image() *ebiten.Image {
	img := SpriteByIndex(s.ImageSheet, s.Index, 16, 16)

	if Sprite.Mirror == 0 {
		// Do not rotate the image
	} else if Sprite.Mirror == 1 {
		// Rotate the image 90 degrees clockwise
		img = ebitenutil.RotateImage(img, 90)
	} else if Sprite.Mirror == 2 {
		// Rotate the image 180 degrees clockwise
		img = ebitenutil.RotateImage(img, 180)
	} else if Sprite.Mirror == 3 {
		// Rotate the image 270 degrees clockwise
		img = ebitenutil.RotateImage(img, 270)
	}
}

type Dimensions struct {
	Width, Height int // size of the tile
}

type Movement struct {
	Vx, Vy   int
	Standing bool // flag referring to the status of the tile
	Blocking bool // flag referring to the status of the tile
}

func (t *Tile) Size() (int, int) {
	return t.Width, t.Height
}

// NewTile creates and initializes a new Tile instance with the image set to nil.
func NewTile() Tile {
	return Tile{
		Id:             0, // id of the tile
		Coordinates:    Point{X: 0, Y: 0}, // coordinates of the tile
		Actions:        Actions{Vx: 0, Vy: 0, Standing: false, Blocking: false}, // movement of the tile
		Width:          0, // size of the tile
		Height:         0, // size of the tile
		Sprite:         Sprite{}, // sprite of the tile
		Type:           0, // type of the tile: 0 = background, 1 = foreground, 2 = player, 3 = enemy, 4 = item, 5 = projectile
		Hitbox:         Hitbox{OffsetX: 0, OffsetY: 0, Width: 0, Height: 0}, // hitbox of the tile
		Visible:        false, // flag referring to the status of the tile
		Acessible:      false, // flag indicating if the tile is acessible by pathfinding
		ResourceType:   "", // type of resource, if applicable
		Status:         "", // status of the tile: "idle", "moving", "attacking", "dead", "on fire", "frozen"
		TriggerEvent:   "", // Identifier for trigger or event associated with the tile
		InteractionPts: []Point{}, // List of interaction points on the tile
		Durability:     0, // Durability or health of the tile
		Owner:          "", // Owner or affiliation of the tile
		CustomData:     make(map[string]interface{}), // Additional custom data
	}
}


type SpriteSheet struct {
	Image                 *ebiten.Image
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
