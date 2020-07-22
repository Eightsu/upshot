package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const playerSpeed float64 = 1.5

type player struct {
	tex  *sdl.Texture
	x, y float64
}

// create new user, requires renderer
func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	img, err := sdl.LoadBMP("assets/Player.bmp")
	if err != nil {
		return player{}, fmt.Errorf("Loading sprite failed: %v", err)
	}
	// after image is in texture, free it
	defer img.Free()
	// player texture
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("Texture wasn't able to be mapped. Error: %v", err)
	}
	p.x = windowWidth / 2.0
	p.y = windowHeight - spriteSize/2.0
	return p, nil
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

}
