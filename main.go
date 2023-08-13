package main

import (
	"fmt"
	// "image/color"
	"log"

	// "math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Define the Player struct
type Player struct {
	image         *ebiten.Image
	x, y          int
	vx, vy        int
	width, height int
	xmax, ymax    int
	isOnGround    bool
}

// the player
var player Player = Player{
	image:      nil,
	x:          50,
	y:          50,
	vx:         0,
	vy:         0,
	width:      16,
	height:     16,
	isOnGround: true,
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
	airDragCoeff		   int  // The coefficient of air drag.
	terminalVelocity       int  // The maximum downward speed.
	isOnGround             bool // Whether the player is on the ground or not.
}

const blockSize = float64(16) // The resolution of a voxel.

var world Physics = Physics{
	xBlockSize:       int(4),
	yBlockSize:       int(4),
	maxGroundSpeed:   int(1.25 * blockSize),
	jumpImpulse:      int(3.125 * blockSize),
	gravity:          int(0.625 * blockSize),
	airControl:       int(0.625 * blockSize),
	groundFriction:   int(0.0625 * blockSize),
	airDragCoeff:    int(9),
	terminalVelocity: int(4 * blockSize),
	isOnGround:       true,
}

// Game implements ebiten.Game interface.
type Game struct {
	cycles        int
	frameCount    int
	droppedFrames int
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
	if player.isOnGround {
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
			player.isOnGround = false
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
	if player.x > windowWidth - player.width{ // If the player is outside the right edge of the window, move it back inside and set the velocity to zero.
		player.x = windowWidth - player.width
		player.vx = 0
	}
	if player.y < 0 { // If the player is outside the top edge of the window, move it back inside and set the velocity to zero.
		player.y = 0
		player.vy = 0
	}
	if player.y > windowHeight - player.height { // If the player is outside the bottom edge of the window, move it back inside and set the velocity to zero.
		player.y = windowHeight - player.height
		player.vy = 0
		player.isOnGround = true
	}

	g.cycles++
	// No errors occurred, return nil (zero-value).
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Game rendering.
	// screen.Fill(color.Black) // Clear the screen to avoid artifacts.


	// define position of the player
	opts.GeoM.Translate(float64(player.vx), float64(player.vy))
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
	return windowWidth , windowHeight
}

func main() {
	game := &Game{
		frameCount:    0,
		cycles:        0,
		droppedFrames: 0,
	}

	// Set the maximum TPS to 60
	ebiten.SetMaxTPS(60)
	ebiten.SetVsyncEnabled(true)

	// Create a new window
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Mario")

	var err error // Declare the 'err' variable to capture the error from NewImageFromFile.

	player.image, _, err = ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	opts.GeoM.Scale(1, 1)

	if err != nil {
		log.Fatal(err)
	}

	// Start the game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
