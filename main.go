package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Define the variables for the rectangle position.
var (
	rectX, rectY          float64 = 160, 120
	rectWidth, rectHeight float64 = 100, 50
	speed                 float64 = 2
)

// This function is called every frame.
func update(screen *ebiten.Image) error {
	// Draw a white rectangle in the center of the screen.
	ebitenutil.DrawRect(screen, 160, 120, 100, 50, color.White)
	return nil
}

func main() {
	// Create a new window with a width of 320 and a height of 240 pixels.
	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Hello, World!")

	// Start the game loop.
	if err := ebiten.Run(update, 320, 240, 2, "hello World"); err != nil {
		panic(err)
	}
}
