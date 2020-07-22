package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const enemySize = 128

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

// create new enemy, requires renderer
func newEnemy(renderer *sdl.Renderer, x, y float64) (e enemy, err error) {
	img, err := sdl.LoadBMP("assets/Enemy.bmp")
	if err != nil {
		return enemy{}, fmt.Errorf("Loading sprite failed: %v", err)
	}
	// after image is in texture, free it
	defer img.Free()
	// enemy texture
	e.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return enemy{}, fmt.Errorf("Texture wasn't able to be mapped. Error: %v", err)
	}
	return e, nil
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	x := e.x 
	y := e.y 
	renderer.Copy(e.tex,
		&sdl.Rect{X: 0, Y: 0, W: enemySize, H: enemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: enemySize, H: enemySize})
}
