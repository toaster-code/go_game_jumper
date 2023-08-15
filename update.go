// The Update function for the game.
package main

import 	(
	"github.com/hajimehoshi/ebiten"
)

// Update function is called every frame.
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
