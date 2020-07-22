package main

import (


	"github.com/veandco/go-sdl2/sdl"
)

const enemySize = 128

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

// create new enemy, requires renderer
func newEnemy(renderer *sdl.Renderer, x, y float64) (e enemy) {
	e.tex = textureFromBMP(renderer,"assets/Enemy.bmp")
	e.x = x
	e.y = y

	return e
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	x := e.x 
	y := e.y 
	renderer.Copy(e.tex,
		&sdl.Rect{X: 0, Y: 0, W: enemySize, H: enemySize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: enemySize, H: enemySize})
}
