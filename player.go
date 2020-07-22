package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const playerSpeed float64 = 1.5

type player struct {
	tex  *sdl.Texture
	x, y float64
}

// create new user, requires renderer
func newPlayer(renderer *sdl.Renderer) (p player) {

	p.tex = textureFromBMP(renderer, "assets/Player.bmp")
	p.x = windowWidth / 2.0
	p.y = windowHeight - spriteSize/2.0
	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - spriteSize/2.0
	y := windowHeight - spriteSize
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: spriteSize, H: spriteSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: spriteSize, H: spriteSize})
}

func (p *player) update(renderer *sdl.Renderer) {

	// Access keyboard inputs
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
		// subtract from x position

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// add to x position
		p.x += playerSpeed
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if b, ok := bulletFromPool(); ok {
			b.active = true
			b.x = p.x - spriteSize
			b.y = p.y - bulletSize
		}
	}

}
