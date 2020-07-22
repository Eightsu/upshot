package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  float64       = 1.5
	shotCooldown time.Duration = time.Millisecond * 200
)

type player struct {
	tex      *sdl.Texture
	x, y     float64
	lastShot time.Time
}


// create new user, requires renderer
func newPlayer(renderer *sdl.Renderer) (*element) {

	p := &element{}

	p.position = vector{
		x: windowWidth / 2.0,
		y: windowHeight - spriteSize/2.0,
	}

	p.active = true


	// p.tex = textureFromBMP(renderer, "assets/Player.bmp")

	// return p
}

func (p *player) draw(renderer *sdl.Renderer) {


}

func (p *player) shoot() {
	if b, ok := bulletFromPool(); ok {
		b.active = true
		b.x = p.x - 16
		b.y = p.y - (bulletSize * 1.5)
		b.angle = 270 * (math.Pi / 180)

	}
}

func (p *player) update(renderer *sdl.Renderer) {

	// Access keyboard inputs
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {

		if p.x <= 32 {
			return
		}
		p.x -= playerSpeed
		// subtract from x position

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {

		if p.x >= windowWidth-32 {
			return
		}
		p.x += playerSpeed
		// add to x position
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShot) >= shotCooldown {
			p.shoot()
			p.lastShot = time.Now()
		}
	}
}
