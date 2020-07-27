package main

import (


	"github.com/veandco/go-sdl2/sdl"
)

const enemySize = 128

// type enemy struct {
// 	tex  *sdl.Texture
// 	x, y float64
// }

func newEnemy(renderer *sdl.Renderer, position vector) *element {
	enemy := &element{}
	enemy.position = position
	enemy.rotation = 0

	sr := newSpriteRenderer(enemy, renderer, "assets/Enemy.bmp" )

	col := circle{
		center: enemy.position,
		radius: 30,
	}
	enemy.collisions = append(enemy.collisions, col)

	dmg := newDamageable(enemy)

	enemy.addComponent(dmg) 


	enemy.addComponent(sr)

	enemy.active = true

	return enemy

}