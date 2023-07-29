package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Define the variables for the rectangle position, velocity, and impulse.
var (
	rectX, rectY          float64 = 150, 150
	rectWidth, rectHeight float64 = 20, 20
	speed                 float64 = 2
	jumpImpulse           float64 = -5 // The upward impulse when jumping.
	gravity               float64 = 0.2
	airControl            float64 = 0.1
	grounded              bool    = true
	velocityX, velocityY  float64
)

// Get the window size.
var (
	windowWidth  = 320
	windowHeight = 240
)

var (
	// The image's dimensions
	imageWidth, imageHeight float64 = 16, 16
	marioImage              *ebiten.Image
)

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
func update(screen *ebiten.Image) error {
	// Move the rectangle based on the arrow keys.
	if grounded {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			velocityX = -speed
		} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
			velocityX = speed
		} else {
			// If no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
			velocityX *= 0.90
		}
	}

	// Apply gravity if the rectangle is not grounded.
	if !grounded {
		velocityY += gravity
	}

	// Handle jumping mechanics.
	if grounded && ebiten.IsKeyPressed(ebiten.KeyUp) {
		velocityY = jumpImpulse
		grounded = false
	}

	// Allow air control to add a smaller effect to the horizontal velocity while in the air.
	if !grounded {
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			velocityX -= airControl
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			velocityX += airControl
		}
	}
	// Allow air friction when no key is pressed.
	if !grounded {
		// if no arrow keys are pressed, gradually reduce the horizontal velocity to simulate friction.
		if ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
			velocityX *= 0.90
		}
	}

	// Update the rectangle's position based on velocity.
	rectX += velocityX
	rectY += velocityY

	// Perform collision detection to prevent the rectangle from moving outside the window frame.
	if rectX < 0 {
		rectX = 0
		velocityX = 0
	}
	if rectX+rectWidth > float64(windowWidth) {
		rectX = float64(windowWidth) - rectWidth
		velocityX = 0
	}
	if rectY < 0 {
		rectY = 0
		velocityY = 0
	}
	if rectY+rectHeight > float64(windowHeight) {
		rectY = float64(windowHeight) - rectHeight
		velocityY = 0
		grounded = true
	}

	// Draw the rectangle at the updated position.
	screen.Fill(color.Black) // Clear the screen to avoid artifacts.
	// ebitenutil.DrawRect(screen, rectX, rectY, rectWidth, rectHeight, color.White)

	// define position
	// opts.GeoM.Scale(3, 3)
	opts.GeoM.Translate(rectX, rectY)
	rectX = 0
	rectY = 0
	screen.DrawImage(marioImage, opts)
	return nil
}

func main() {
	// Create a new window with a width of 320 and a height of 240 pixels.
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Move the Rectangle!")

	var err error // Declare the 'err' variable to capture the error from NewImageFromFile.

	marioImage, _, err = ebitenutil.NewImageFromFile("./res/small_mario_p0.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}


	// Start the game loop.
	if err := ebiten.Run(update, windowWidth, windowHeight, 2, "Move the Rectangle"); err != nil {
		log.Fatal(err)
	}
}
