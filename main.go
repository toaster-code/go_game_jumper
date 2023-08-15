package main

// tips for importing packages:
// the import of local packages is relative to the current package, and should refer to the subdirectory name

import (
	"fmt"
	"image/color"
	"go_game_jumper/src/tiles" // import the tiles package
	"go_game_jumper/src/tools" // import the tools package
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

var a = tiles.NewTileSlice(1)

// the player
// var player t.Tile =  t.NewTile( nil, 50, 50, 0, 0, 16, 16, true, true)

var player tiles.Tile = tiles.Tile{
	Image:    nil,
	X:        50,
	Y:        50,
	Vx:       0,
	Vy:       0,
	Width:    16,
	Height:   16,
	Standing: true,
	Blocking: true,
}

// Get the window size.
const (
	windowWidth  int = 320 // The width of the window in pixels.
	windowHeight int = 240 // The height of the window in pixels.
)

// Define the variables for the player position, velocity, and impulse.
// Note that precision using integers is considering 1 decimal place.
// This is to simulate sub-pixel precision using integers.

type Physics = struct {
	maxGroundSpeed   int  // The maximum horizontal speed when on the ground.
	jumpImpulse      int  // The upward impulse when jumping.
	gravity          int  // The downward acceleration due to gravity.
	airControl       int  // The horizontal deceleration when in the air.
	groundFriction   int  // The horizontal deceleration when on the ground.
	airDragCoeff     int  // The coefficient of air drag.
	terminalVelocity int  // The maximum downward speed.
	Standing         bool // Whether the player is on the ground or not.
}

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

// Game implements ebiten.Game interface.
type Game struct {
	cycles        int
	frameCount    int
	droppedFrames int
}

type GameData struct {
	ScreenWidth  int
	ScreenHeight int
	TileWidth    int
	TileHeight   int
}

// NewGameData creates a new GameData struct
func NewGameData() GameData {
	g := GameData{
		ScreenWidth:  80,
		ScreenHeight: 50,
		TileWidth:    16,
		TileHeight:   16,
	}
	return g
}

// Define DrawImageOptions
var opts = &ebiten.DrawImageOptions{}

func NewImageFromImage(source *ebiten.Image) (*ebiten.Image, error) {
	// Create a new image with the same size as the source image.
	newImage, err := ebiten.NewImageFromImage(source, ebiten.FilterDefault)
	if err != nil {
		newImage = nil
	}
	return newImage, err
}

// This function is called every frame.
func (g *Game) Update(screen *ebiten.Image) error {

	println("Update step: ", g.cycles)


	// Game logic goes here.

	// Move the rectangle based on the arrow keys.
	if player.Standing {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			player.Vx = -world.maxGroundSpeed
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.Vx = +world.maxGroundSpeed
		} else {
			// If no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
			if player.Vx > 0 {
				player.Vx -= world.groundFriction
			} else if player.Vx < 0 {
				player.Vx += world.groundFriction
			}
		}
	} else {
		// Apply gravity when in air.
		player.Vy += world.gravity

		// Saturate fall (terminal velocity):
		if player.Vy > world.terminalVelocity {
			player.Vy = world.terminalVelocity
		}

		// Handle jumping mechanics.
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			player.Vy = -world.jumpImpulse
			player.Standing = false
		}

		// Allow air control to add a smaller effect to the horizontal velocity while in the air.
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			player.Vx -= world.airControl
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.Vx += world.airControl
		}

		// Allow air friction when no key is pressed.
		// if no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
		if ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.Vx *= world.airDragCoeff
		}
	}

	// Update the position based on velocity.
	player.X += player.Vx
	player.Y += player.Vy

	// Perform collision detection to prevent the rectangle from moving outside the window frame.
	if player.X < 0 { // If the player is outside the left edge of the window, move it back inside and set the velocity to zero.
		player.X = 0
		player.Vx = 0
	}
	if player.X > windowWidth-player.Width { // If the player is outside the right edge of the window, move it back inside and set the velocity to zero.
		player.X = windowWidth - player.Width
		player.Vx = 0
	}
	if player.Y < 0 { // If the player is outside the top edge of the window, move it back inside and set the velocity to zero.
		player.Y = 0
		player.Vy = 0
	}
	if player.Y > windowHeight-player.Height { // If the player is outside the bottom edge of the window, move it back inside and set the velocity to zero.
		player.Y = windowHeight - player.Height
		player.Vy = 0
		player.Standing = true
	}

	// define position of the player
	opts.GeoM.Translate(float64(player.Vx), float64(player.Vy))

	g.cycles++

	// No errors occurred, return nil (zero-value).
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	println("Draw step: ", g.frameCount)

	// Game rendering.
	screen.Fill(color.Black) // Clear the screen to avoid artifacts.

	// Draw the player's image at its X and Y coordinates.
	opts.GeoM.Reset()

	screen.DrawImage(player.Image, opts)
	screen.DrawImage(a[0].Image, opts)

	// Print stats on the screen
	statsText := fmt.Sprintf("Stats:\nX: %d", player.X)
	// Display FPS and dropped frames on the screen.
	statsText += fmt.Sprintf("\nFPS: %0.1f\nDropped Frames: %d\nFrame=%d\ncycle=%d", ebiten.CurrentTPS(), g.droppedFrames, g.cycles, g.frameCount)
	ebitenutil.DebugPrint(screen, statsText)

    // Draw the tiles from the 'a' slice.
    for i := 0; i < len(a); i++ {
        tile := a[i] // Get the current tile

        // Draw the tile's image at its X and Y coordinates.
        opts.GeoM.Reset()
        opts.GeoM.Translate(float64(tile.X), float64(tile.Y))
        screen.DrawImage(tile.Image, opts)
    }

	g.frameCount++
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

// NewGame creates a new Game Object and initializes the data
// This is a pretty solid refactor candidate for later
func NewGame() *Game {
	g := &Game{
		frameCount:    0,
		cycles:        0,
		droppedFrames: 0,
	}
	return g
}



func main() {

	println(tiles.Check_if_working)
	println(tools.Max(1, 5))
	var err error

	// var wallImage *ebiten.Image

	// wallImage, _, err = ebitenutil.NewImageFromFile("./res/Tilemaps/Png Files/wood_moss_alt_tileset_2.png", ebiten.FilterDefault)
	Test_image, _, _ := ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	player.Image, _, _ = ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	a[0] = tiles.NewTile(Test_image, 0, 0, 0, 0, 16, 16, true, true)
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
