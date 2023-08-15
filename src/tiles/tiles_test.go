package tiles

import (
	"testing"

	"github.com/hajimehoshi/ebiten"
)

// TestNewTile tests the NewTile function.
// NewTile creates a new Tile with the given parameters and returns it.
// It takes an ebiten.Image, x, y, vx, vy, width, height, standing and blocking as parameters.
// It returns a Tile.
func TestNewTile(t *testing.T) {
	img, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	tile := NewTile(img, 0, 0, 0, 0, 16, 16, true, true)

	if tile.SpriteSheetImage != img {
		t.Errorf("Expected SpriteSheetImage to be %v but got %v", img, tile.SpriteSheetImage)
	}

	if tile.X != 0 {
		t.Errorf("Expected X to be 0 but got %v", tile.X)
	}

	if tile.Y != 0 {
		t.Errorf("Expected Y to be 0 but got %v", tile.Y)
	}

	if tile.Vx != 0 {
		t.Errorf("Expected Vx to be 0 but got %v", tile.Vx)
	}

	if tile.Vy != 0 {
		t.Errorf("Expected Vy to be 0 but got %v", tile.Vy)
	}

	if tile.Width != 16 {
		t.Errorf("Expected Width to be 16 but got %v", tile.Width)
	}

	if tile.Height != 16 {
		t.Errorf("Expected Height to be 16 but got %v", tile.Height)
	}

	if tile.Standing != true {
		t.Errorf("Expected Standing to be true but got %v", tile.Standing)
	}

	if tile.Blocking != true {
		t.Errorf("Expected Blocking to be true but got %v", tile.Blocking)
	}
}

// TestNewTileSlice tests the NewTileSlice function.
// NewTileSlice creates a new slice of Tiles with the given size and returns it.
// It takes size as a parameter.
// It returns a slice of Tiles.
func TestNewTileSlice(t *testing.T) {
	size := 10
	tileSlice := NewTileSlice(size)

	if len(tileSlice) != size {
		t.Errorf("Expected slice length to be %v but got %v", size, len(tileSlice))
	}
}

// TestNumSprites tests the numSprites function.
// numSprites returns the number of sprites in the Tile's sprite sheet image.
// It takes no parameters.
// It returns an int.
func TestNumSprites(t *testing.T) {
	img, _ := ebiten.NewImage(16, 16, ebiten.FilterDefault)
	tile := NewTile(img, 0, 0, 0, 0, 16, 16, true, true)

	if tile.numSprites() != 1 {
		t.Errorf("Expected numSprites to be 1 but got %v", tile.numSprites())
	}
}

// TestNumSprites tests the numSprites function. Test for 10 sprites.
// numSprites returns the number of sprites in the Tile's sprite sheet image.
// It takes no parameters.
// It returns an int.
func TestNumSprites10(t *testing.T) {
	img, _ := ebiten.NewImage(160, 16, ebiten.FilterDefault)
	tile := NewTile(img, 0, 0, 0, 0, 16, 16, true, true)

	if tile.numSprites() != 10 {
		t.Errorf("Expected numSprites to be 10 but got %v", tile.numSprites())
	}
}


// TestSpriteByIndex tests the SpriteByIndex function.
// SpriteByIndex returns the sprite at the given index in the Tile's sprite sheet image.
// It takes an index as a parameter.
// It returns an ebiten.Image.
func TestSpriteByIndex(t *testing.T) {
	img, _ := ebiten.NewImage(32, 16, ebiten.FilterDefault)
	tile := NewTile(img, 0, 0, 0, 0, 16, 16, true, true)

	sprite := tile.SpriteByIndex(1)

	if sprite.Bounds().Dx() != 16 {
		t.Errorf("Expected sprite width to be 16 but got %v", sprite.Bounds().Dx())
	}

	if sprite.Bounds().Dy() != 16 {
		t.Errorf("Expected sprite height to be 16 but got %v", sprite.Bounds().Dy())
	}
}
