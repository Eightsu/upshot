package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  float64       = 1.5
	playerSize = 128
	shotCooldown time.Duration = time.Millisecond * 200
)
// create new user, requires renderer
func newPlayer(renderer *sdl.Renderer) *element {

	p := &element{}

	p.position = vector{
		x: windowWidth / 2.0,
		y: windowHeight - spriteSize/2.0,
	}

	p.active = true

	sr := newSpriteRenderer(p, renderer, "assets/Player.bmp")
	// Add sprite functionality
	p.addComponent(sr)

	kb := newKBInput(p, playerSpeed)
	// Add controller functionality
	p.addComponent(kb)

	shooter := newKBShooter(p, shotCooldown)
	p.addComponent(shooter)

	// p.tex = textureFromBMP(renderer, "assets/Player.bmp")
	return p
}
