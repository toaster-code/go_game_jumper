// main package implements a simple game in Go using the Ebiten game library.
// The game is a platformer where the player must jump over obstacles to reach the end of the level.
package main

import (
	"go_game_jumper/src/tiles" // import the tiles package
	"go_game_jumper/src/tools" // import the tools package
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const blockSize = float64(16) // The resolution of a voxel.

var world Physics = Physics{
	maxGroundSpeed:   int(2),
	jumpImpulse:      int(3.125 * blockSize), // 3.125 * blockSize = 50
	gravity:          int(0.625 * blockSize),
	airControl:       int(0.625 * blockSize),
	groundFriction:   int(0.0625 * blockSize),
	airDragCoeff:     int(9),
	terminalVelocity: int(4 * blockSize),
	Standing:         true,
}

// Define DrawImageOptions
var opts = &ebiten.DrawImageOptions{}

func main() {

	println(tiles.Check_if_working)
	println(tools.Max(1, 5))
	var err error

	var tileSetImage *ebiten.Image

	tileSetImage, _, err = ebitenutil.NewImageFromFile("./res/Tilemaps/Png Files/wood_moss_alt_tileset_2.png", ebiten.FilterDefault)
	// Test_image, _, _ := ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	player.SpriteSheetImage, _, _ = ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	// Extract textures from the tileset image.
	img := tiles.SpriteByIndex(tileSetImage, 5, 16, 16) // wallImage

	a[0] = tiles.NewTile(img, 0, 0, 0, 0, 16, 16, true, true)
	// a[1] = tiles.NewTile(player.Image, 150, 150, 0, 0, 16, 16, true, true)
	// a[2] = tiles.NewTile(wallImage, 250, 400, 0, 0, 16, 16, true, true)

	game := NewGame()

	// Set the maximum TPS to 60
	ebiten.SetMaxTPS(10)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowResizable(true)

	// Create a new window
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Mario")

	opts.GeoM.Scale(2, 2)

	// define start position of the player
	// opts.GeoM.Translate(float64(player.X), float64(player.Y))

	if err != nil {
		log.Fatal(err)
	}

	// Start the game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
