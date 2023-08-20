//main2.go

package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	mainGameWindow *ebiten.Image
	secondWindow   *ebiten.Image
	spriteSheetImage *ebiten.Image
)

func init() {
	var err error
	mainGameWindow, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	secondWindow, _ = ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)

	spriteSheetImage, _, err = ebitenutil.NewImageFromFile("spritesheet.png")
	if err != nil {
		fmt.Println("Error loading image:", err)
	}
}

func updateMainGame() error {
	// Put your main game logic here
	return nil
}

func drawMainGame(screen *ebiten.Image) {
	screen.DrawImage(mainGameWindow, nil)
}

func updateSecondWindow() error {
	// Put your logic for the second window here
	return nil
}

func drawSecondWindow(screen *ebiten.Image) {
	screen.DrawImage(secondWindow, nil)
	secondWindow.DrawImage(spriteSheetImage, nil) // Draw the sprite sheet image on the second window
}

func main2() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Main Game Window")

	// Initialize the main game loop for the main game window
	if err := ebiten.RunGame(updateMainGame, drawMainGame); err != nil {
		fmt.Println("Error in main game loop:", err)
	}

	// Initialize the game loop for the second window
	if err := ebiten.RunGame(updateSecondWindow, drawSecondWindow); err != nil {
		fmt.Println("Error in second window loop:", err)
	}
}
