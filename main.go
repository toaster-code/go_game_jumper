package main

// tips for importing packages:
// the import of local packages is relative to the current package, and should refer to the subdirectory name

import (
	"fmt"
	// "image/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	t "go_game_jumper/src/tiles" // import the tiles package
	"log"
)

// GetIndexFromXY gets the index of the map array from a given X,Y TILE coordinate.
// This coordinate is logical tiles, not pixels.
func GetIndexFromXY(x int, y int) int {
	gd := NewGameData()
	return (y * gd.ScreenWidth) + x
}

// the player
var player t.Tile =  t.NewTile( nil, 50, 50, 0, 0, 16, 16, true, true)

var playerqq t.Tile = t.Tile{
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
	xBlockSize, yBlockSize int  // The resolution of a voxel.
	maxGroundSpeed         int  // The maximum horizontal speed when on the ground.
	jumpImpulse            int  // The upward impulse when jumping.
	gravity                int  // The downward acceleration due to gravity.
	airControl             int  // The horizontal deceleration when in the air.
	groundFriction         int  // The horizontal deceleration when on the ground.
	airDragCoeff           int  // The coefficient of air drag.
	terminalVelocity       int  // The maximum downward speed.
	standing               bool // Whether the player is on the ground or not.
}

const blockSize = float64(16) // The resolution of a voxel.

var world Physics = Physics{
	xBlockSize:       int(4),
	yBlockSize:       int(4),
	maxGroundSpeed:   int(2),
	jumpImpulse:      int(3.125 * blockSize), // 3.125 * blockSize = 50
	gravity:          int(0.625 * blockSize),
	airControl:       int(0.625 * blockSize),
	groundFriction:   int(0.0625 * blockSize),
	airDragCoeff:     int(9),
	terminalVelocity: int(4 * blockSize),
	standing:         true,
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

	// Move the rectangle based on the arrow keys.
	if player.standing {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			player.vx = -world.maxGroundSpeed
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.vx = world.maxGroundSpeed
		} else {
			// If no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
			if player.vx > 0 {
				player.vx -= world.groundFriction
			} else if player.vx < 0 {
				player.vx += world.groundFriction
			}
		}
	} else {
		// Apply gravity when in air.
		player.vy += world.gravity

		// Saturate fall (terminal velocity):
		if player.vy > world.terminalVelocity {
			player.vy = world.terminalVelocity
		}

		// Handle jumping mechanics.
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			player.vy = -world.jumpImpulse
			player.standing = false
		}

		// Allow air control to add a smaller effect to the horizontal velocity while in the air.
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			player.vx -= world.airControl
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.vx += world.airControl
		}

		// Allow air friction when no key is pressed.
		// if no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
		if ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.vx *= world.airDragCoeff
		}
	}

	// Update the position based on velocity.
	player.x += player.vx
	player.y += player.vy

	// Perform collision detection to prevent the rectangle from moving outside the window frame.
	if player.x < 0 { // If the player is outside the left edge of the window, move it back inside and set the velocity to zero.
		player.x = 0
		player.vx = 0
	}
	if player.x > windowWidth-player.width { // If the player is outside the right edge of the window, move it back inside and set the velocity to zero.
		player.x = windowWidth - player.width
		player.vx = 0
	}
	if player.y < 0 { // If the player is outside the top edge of the window, move it back inside and set the velocity to zero.
		player.y = 0
		player.vy = 0
	}
	if player.y > windowHeight-player.height { // If the player is outside the bottom edge of the window, move it back inside and set the velocity to zero.
		player.y = windowHeight - player.height
		player.vy = 0
		player.standing = true
	}

	// define position of the player
	opts.GeoM.Translate(float64(player.vx), float64(player.vy))

	g.cycles++

	// No errors occurred, return nil (zero-value).
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Game rendering.
	// screen.Fill(color.Black) // Clear the screen to avoid artifacts.

	screen.DrawImage(player.image, opts)

	// Print stats on the screen
	statsText := fmt.Sprintf("Stats:\nX: %d", player.x)
	// Display FPS and dropped frames on the screen.
	statsText += fmt.Sprintf("\nFPS: %0.1f\nDropped Frames: %d\nFrame=%d\ncycle=%d", ebiten.CurrentTPS(), g.droppedFrames, g.cycles, g.frameCount)
	ebitenutil.DebugPrint(screen, statsText)

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

func LoadSprites() *GameTiles {

}

func main() {
	game := NewGame()

	// Set the maximum TPS to 60
	ebiten.SetMaxTPS(10)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowResizable(true)

	// Create a new window
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Mario")

	var err error // Declare the 'err' variable to capture the error from NewImageFromFile.

	// load sprites
	// Create a Tile instance from the tiles package
	mytile := tiles.Tile{
		x: 1,
		y: 2,
	}

	player.image, _, err = ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	wallImage, _, err := ebitenutil.NewImageFromFile("./res/Tilemaps/Png Files/wood_moss_alt_tileset_2.png", ebiten.FilterDefault)

	opts.GeoM.Scale(1, 1)

	// define start position of the player
	opts.GeoM.Translate(float64(player.x), float64(player.y))

	if err != nil {
		log.Fatal(err)
	}

	// Start the game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
