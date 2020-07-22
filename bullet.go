package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 3.0
)

type bullet struct {
	tex    *sdl.Texture
	x, y   float64
	angle  float64
	active bool
}

func newBullet(renderer *sdl.Renderer) (b bullet) {
	b.tex = textureFromBMP(renderer, "assets/Bullet.bmp")
	return b
}

func (b *bullet) draw(renderer *sdl.Renderer) {

	if !b.active {
		return
	}

	x := b.x
	y := b.y

	renderer.Copy(b.tex,
		&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize})
}

func (b *bullet) update() {
	b.x += bulletSpeed * math.Cos(b.angle)
	b.y += bulletSpeed * math.Sin(b.angle)

	if b.y < 0-(spriteSize/2) {
		b.active = false
	}
}

var bulletPool []*bullet

func initBPool(renderer *sdl.Renderer) {
	for i := 0; i < 20; i++ {
		b := newBullet(renderer)
		bulletPool = append(bulletPool, &b)
	}
}

func bulletFromPool() (*bullet, bool) {
	for _, b := range bulletPool {
		if !b.active {
			return b, true
		}
	}

	return nil, false
}
