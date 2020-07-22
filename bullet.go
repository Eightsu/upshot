package main

import "github.com/veandco/go-sdl2/sdl"

const bulletSize = 32

type bullet struct {
	tex  *sdl.Texture
	x, y float64
}

func newBullet(renderer *sdl.Renderer) (b bullet) {
	b.tex = textureFromBMP(renderer, "assets/Bullet.bmp")
	return b
}

func (b *bullet) draw(renderer *sdl.Renderer) {

	x := b.x
	y := b.y

	renderer.Copy(b.tex,
		&sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 32, H: 32})
}

var bulletPool []bullet

func initBPool(renderer *sdl.Renderer) {
	for i := 0; i < 20; i++ {
		bulletPool = append(bulletPool, newBullet(renderer))
	}
}
