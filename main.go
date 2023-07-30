package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Define the variables for the rectangle position, velocity, and impulse.
var (
	playerPosX, playerPosY    float64 = 0, 0
	playerWidth, playerHeight float64 = 16, 16
	maxGroundSpeed            float64 = 2
	jumpImpulse               float64 = -5 // The upward impulse when jumping.
	gravity                   float64 = 0.2
	airControl                float64 = 0.1
	isOnGround                bool    = true
	playerVelX, playerVelY    float64 = 0, 0
)

// Get the window size.
const (
	windowWidth  = 320
	windowHeight = 240
)

var (
	// The image's dimensions
	playerImage              *ebiten.Image
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
			playerVelX = -maxGroundSpeed
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			playerVelX = maxGroundSpeed
		} else {
			// If no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
			playerVelX *= 0.95
		}
	} else {
		// Apply gravity if the rectangle is not grounded.
		playerVelY += gravity


		// Handle jumping mechanics.
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			playerVelY = jumpImpulse
			isOnGround = false
		}

		// Allow air control to add a smaller effect to the horizontal velocity while in the air.
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			playerVelX -= airControl
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			playerVelX += airControl
		}

		// Allow air friction when no key is pressed.
		// if no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
		if ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
			playerVelX *= 0.90
		}
	}

	// Update the rectangle's position based on velocity.
	playerPosX += playerVelX
	playerPosY += playerVelY

	// Perform collision detection to prevent the rectangle from moving outside the window frame.
	if playerPosX < 0 {
		playerPosX = 0
		playerVelX = 0
	}
	if playerPosX+playerWidth > float64(windowWidth) {
		playerPosX = float64(windowWidth) - playerWidth
		playerVelX = 0
	}
	if playerPosY < 0 {
		playerPosY = 0
		playerVelY = 0
	}
	if playerPosY+playerHeight > float64(windowHeight) {
		playerPosY = float64(windowHeight) - playerHeight
		playerVelY = 0
		isOnGround = true
	}

	// No errors occurred, return nil (zero-value).
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	// Game rendering.
	screen.Fill(color.Black) // Clear the screen to avoid artifacts.

	// Print stats on the screen
	statsText := fmt.Sprintf("Stats:\nX: %.1f\nY: %.1f\nVelocityX: %.1f\nVelocityY: %.1f\nisOngroud %t", playerPosX, playerPosY, playerVelX, playerVelY, isOnGround)
	ebitenutil.DebugPrint(screen, statsText)

	// Draw the rectangle at the updated position.
	// ebitenutil.DrawRect(screen, rectX, rectY, rectWidth, rectHeight, color.White)

	// define position
	opts.GeoM.Translate(playerVelX, playerVelY)
	// opts.GeoM.Translate(rectX, rectY)
	screen.DrawImage(playerImage, opts)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth*1.8, windowHeight*1.8
}

func main() {
	game := &Game{}

	// Create a new window with a width of 320 and a height of 240 pixels.
	ebiten.SetWindowSize(windowWidth*2, windowHeight*2)
	ebiten.SetWindowTitle("Mario")

	var err error // Declare the 'err' variable to capture the error from NewImageFromFile.

	playerImage, _, err = ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	opts.GeoM.Scale(2, 2)

	if err != nil {
		log.Fatal(err)
	}

	// Start the game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
