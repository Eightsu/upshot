package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSize  = 32
	bulletSpeed = 3.0
)

func newBullet(renderer *sdl.Renderer) *element {

	bullet := &element{}
	sr := newSpriteRenderer(bullet, renderer, "assets/Bullet.bmp")

	bullet.addComponent(sr)

	mv := newBulletMover(bullet, bulletSpeed)

	bullet.addComponent(mv)


	col := circle{
		center: bullet.position,
		radius: 1,
	}
	bullet.collisions = append(bullet.collisions, col)
	bullet.active = false

	return bullet
}

var bulletPool []*element

// renderer.Copy(b.tex,
// 	&sdl.Rect{X: 0, Y: 0, W: bulletSize, H: bulletSize},
// 	&sdl.Rect{X: int32(x), Y: int32(y), W: bulletSize, H: bulletSize})

// func (b *bullet) update() {
// 	b.x += bulletSpeed * math.Cos(b.angle)
// 	b.y += bulletSpeed * math.Sin(b.angle)

// 	if b.y < 0-(spriteSize/2) {
// 		b.active = false
// 	}
// }

func initBPool(renderer *sdl.Renderer) {
	for i := 0; i < 20; i++ {
		b := newBullet(renderer)
		elements = append(elements, b)
		bulletPool = append(bulletPool, b)
	}
}

func bulletFromPool() (*element, bool) {
	for _, b := range bulletPool {
		if !b.active {
			return b, true
		}
	}
	return nil, false
}
