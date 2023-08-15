// The graphics functions for the game.
// Includes the draw function.
package main

import (
	"fmt"
	"go_game_jumper/src/tiles" // import the tiles package
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// ScreenData represents the dimensions of the game screen and the size of the tiles used in the game.
type ScreenData struct {
	ScreenWidth  int // The width of the game screen in characters.
	ScreenHeight int // The height of the game screen in characters.
	TileWidth    int // The width of a tile in pixels.
	TileHeight   int // The height of a tile in pixels
}

// NewScreenData returns a new ScreenData struct with default values for the screen and tile dimensions.
//
// The ScreenWidth and ScreenHeight fields are set to 80 and 50, respectively, which represents the number of characters that can be displayed on the screen.
// The TileWidth and TileHeight fields are set to 16, which represents the size of each tile in pixels.
//
// Returns:
//
//	A new ScreenData struct with default values.
func NewScreenData() ScreenData {
	g := ScreenData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}
	return g
}

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func (s ScreenData) GetIndexFromXY(x int, y int) int {
	return (y * s.ScreenWidth) + x
}

// GetXYFromIndex gets the X,Y TILE coordinate from a given index.
// This coordinate is logical tiles, not pixels.
func (s ScreenData) GetXYFromIndex(index int) (int, int) {
	return index % s.ScreenWidth, index / s.ScreenWidth
}

// Get the window size.
const (
	windowWidth  int = 320 // The width of the window in pixels.
	windowHeight int = 240 // The height of the window in pixels.
)

var a = tiles.NewTileSlice(1)

// the player
// var player t.Tile =  t.NewTile( nil, 50, 50, 0, 0, 16, 16, true, true)

var player tiles.Tile = tiles.Tile{
	SpriteSheetImage: nil,
	X:                50,
	Y:                50,
	Vx:               0,
	Vy:               0,
	Width:            16,
	Height:           16,
	Standing:         true,
	Blocking:         true,
}

// Draw draws the game screen.
// It is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	println("Draw step: ", g.frameCount)

	// Game rendering.
	screen.Fill(color.Black) // Clear the screen to avoid artifacts.

	// Draw the player's image at its X and Y coordinates.
	opts.GeoM.Reset()

	screen.DrawImage(player.SpriteSheetImage, opts)
	screen.DrawImage(a[0].SpriteSheetImage, opts)

	// Draw the tiles from the 'a' slice.
	for i := 0; i < len(a); i++ {
		tile := a[i] // Get the current tile

		// Draw the tile's image at its X and Y coordinates.
		opts.GeoM.Reset()
		opts.GeoM.Translate(float64(tile.X), float64(tile.Y))
		screen.DrawImage(tile.SpriteSheetImage, opts)
	}

	// Print stats on the screen
	statsText := fmt.Sprintf("Stats:\nX: %d", player.X)
	// Display FPS and dropped frames on the screen.
	statsText += fmt.Sprintf("\nFPS: %0.1f\nDropped Frames: %d\nFrame=%d\ncycle=%d", ebiten.CurrentTPS(), g.droppedFrames, g.cycles, g.frameCount)
	ebitenutil.DebugPrint(screen, statsText)

	// Increment the frame count for the next frame.
	g.frameCount++
}

// DrawScreenBorder draws a border around the game screen.
func DrawScreenBorder() {

}
