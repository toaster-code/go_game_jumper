// File that stores the game logic and physics.
package main

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


// Game implements ebiten.Game interface.
type Game struct {
	cycles        int
	frameCount    int
	droppedFrames int
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

// Layout method takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

