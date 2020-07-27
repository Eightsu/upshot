package main

import "github.com/veandco/go-sdl2/sdl"

type damageable struct {
	container *element
}

func newDamageable(container *element) *damageable {

	return &damageable{
		container: container,
	}
}

func (damageTarget *damageable) draw(renderer *sdl.Renderer) error {

	return nil

}

func (damageTarget *damageable) update() error {
	return nil
}

func (damageTarget *damageable) onCollision(other *element) error {
	damageTarget.container.active = false
	return nil

}
