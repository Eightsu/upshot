package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *element
	speed     float64
}

func newBulletMover(container *element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}
}

func (b *bulletMover) draw(renderer *sdl.Renderer) error {
return nil
}

func (b *bulletMover) update() error {

	entity := b.container

	entity.position.x += bulletSpeed * math.Cos(entity.rotation)
	entity.position.y += bulletSpeed * math.Sin(entity.rotation)

	if b.container.position.y > windowHeight || b.container.position.y < 0 {
		b.container.active = false
	}
	return nil
}
