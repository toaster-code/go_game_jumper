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
	x, y, vx, vy, width, height int
	xmax, ymax                  int
}

// the player
var player Player = Player{
	x:      0,
	y:      50,
	vx:     0,
	vy:     0,
	width:  16,
	height: 16,
	xmax:   0,
	ymax:   0,
}

// Get the window size.
const (
	subPixelSize      int     = 4                      // The size in bits.
	floatNumSubPixels float64 = 2 * 2 * 2 * 2          // The number of sub-pixels per pixel.
	intNumSubPixels   int     = int(floatNumSubPixels) // The number of sub-pixels per pixel.
	windowWidth       int     = 320 * intNumSubPixels  // The width of the window in pixels.
	windowHeight      int     = 240 * intNumSubPixels  // The height of the window in pixels.
)

// Define the variables for the player position, velocity, and impulse.
// Note that precision using integers is considering 1 decimal place.
// This is to simulate sub-pixel precision using integers.
var (
	subPixelPosX, subPixelPosY int  = int(0 * intNumSubPixels), int(50 * intNumSubPixels) // The sub-pixel position of the target.
	playerWidth, playerHeight  int  = int(1 * intNumSubPixels), int(1 * intNumSubPixels)  // The width and height of the target.
	maxGroundSpeed             int  = int(1.25 * floatNumSubPixels)                       // The maximum horizontal speed when on the ground.
	jumpImpulse                int  = int(-3.125 * floatNumSubPixels)                     // The upward impulse when jumping.
	gravity                    int  = int(0.625 * floatNumSubPixels)                      // The downward acceleration due to gravity.
	airControl                 int  = int(0.625 * floatNumSubPixels)                      // The horizontal deceleration when in the air.
	groundFriction             int  = int(0.0625 * floatNumSubPixels)                     // The horizontal deceleration when on the ground.
	isOnGround                 bool = true
)

var (
	// The image's dimensions
	playerImage *ebiten.Image
)

// Game implements ebiten.Game interface.
type Game struct{}

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
	if isOnGround {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			player.vx = -maxGroundSpeed
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.vx = maxGroundSpeed
		} else {
			// If no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
			if player.vx > 0 {
				player.vx -= groundFriction
			} else if player.vx < 0 {
				player.vx += groundFriction
			}
		}
	} else {
		// Apply gravity if the rectangle is not grounded.
		player.vy += player.vx * gravity

		// Handle jumping mechanics.
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			player.vy = jumpImpulse
			isOnGround = false
		}

		// Allow air control to add a smaller effect to the horizontal velocity while in the air.
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			player.vx -= airControl
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.vx += airControl
		}

		// Allow air friction when no key is pressed.
		// if no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
		if ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
			player.vx *= 9
		}
	}

	// Update the sub-pixel position based on velocity.
	subPixelPosX += player.vx
	subPixelPosY += player.vy
	player.x = subPixelPosX
	player.y = subPixelPosY

	// Perform collision detection to prevent the rectangle from moving outside the window frame.
	if player.x < 0 { // If the player is outside the left edge of the window, move it back inside and set the velocity to zero.
		subPixelPosX = 0
		player.vx = 0
	}
	if player.x+playerWidth > windowWidth { // If the player is outside the right edge of the window, move it back inside and set the velocity to zero.
		subPixelPosX = windowWidth - playerWidth
		player.vx = 0
	}
	if player.y < 0 { // If the player is outside the top edge of the window, move it back inside and set the velocity to zero.
		subPixelPosY = 0
		player.vy = 0
	}
	if player.y+playerHeight > windowHeight { // If the player is outside the bottom edge of the window, move it back inside and set the velocity to zero.
		subPixelPosY = windowHeight - playerHeight
		player.vy = 0
		isOnGround = true
	}

	player.xmax = Max(player.xmax, player.x)
	player.ymax = Max(player.ymax, player.y)

	// No errors occurred, return nil (zero-value).
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Game rendering.
	// screen.Fill(color.Black) // Clear the screen to avoid artifacts.

	// define position of the player
	// opts.GeoM.Translate(float64(player.vx)/floatNumSubPixels, float64(player.vy)/floatNumSubPixels)
	opts.GeoM.Translate(float64(player.vx/intNumSubPixels), float64(player.vy/intNumSubPixels))
	screen.DrawImage(playerImage, opts)

	// Print stats on the screen
	// statsText := fmt.Sprintf("Stats:\nX: %.1f\nY: %.1f\nVelocityX: %.1f\nVelocityY: %.1f\nisOngroud %t", player.x, player.y, player.vx, player.vy, isOnGround)
	// statsText := fmt.Sprintf("Stats:\nX: %1.2f(%1.2f)\nY: %d\nVelocityX: %d\nVelocityY: %d\nisOngroud %t", player.x, player.x*intSubPixelQty , player.y, player.vx, player.vy, isOnGround)
	statsText := fmt.Sprintf("Stats:\nX: %d(%d)\nmax: %d", player.x, player.x/intNumSubPixels, player.xmax)
	ebitenutil.DebugPrint(screen, statsText)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth/intNumSubPixels, windowHeight/intNumSubPixels
}

func main() {
	game := &Game{}

	// Create a new window
	ebiten.SetWindowSize(windowWidth/intNumSubPixels, windowHeight/intNumSubPixels)
	ebiten.SetWindowTitle("Mario")

	var err error // Declare the 'err' variable to capture the error from NewImageFromFile.

	playerImage, _, err = ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	opts.GeoM.Scale(1, 1)

	if err != nil {
		log.Fatal(err)
	}

	// Start the game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
