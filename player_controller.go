package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type kbInput struct {
	container *element
	speed     float64
	sr        *spriteRenderer
}

type keyBoardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKBShooter(container *element, cooldown time.Duration) *keyBoardShooter {

	return &keyBoardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (shooter *keyBoardShooter) shoot(x float64, y float64) {
	if b, ok := bulletFromPool(); ok {

		b.active = true
		b.position.x = x
		b.position.y = y
		b.rotation = 270 * (math.Pi / 180)

	}
}

func (shooter *keyBoardShooter) draw(renderer *sdl.Renderer) error {

	return nil
}
func (shooter *keyBoardShooter) update() error {

	keys := sdl.GetKeyboardState()

	pos := shooter.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.cooldown {
			shooter.shoot(pos.x, pos.y-24)
			shooter.lastShot = time.Now()
		}
		return nil
	}
	return nil
}

func newKBInput(container *element, speed float64) *kbInput {

	return &kbInput{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}

}

func (entity *kbInput) update() error {
	keys := sdl.GetKeyboardState()

	cont := entity.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if cont.position.x-(entity.sr.width/2.0) > 0 {
			cont.position.x -= entity.speed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if cont.position.x+(entity.sr.height/2.0) < windowWidth {
			cont.position.x += entity.speed
		}
	}

	return nil
}

// func (entity *kbInput) update() error {
// 	// Access keyboard inputs
// 	keys := sdl.GetKeyboardState()

// 	pos := entity.container.position

// 	if keys[sdl.SCANCODE_LEFT] == 1 {
// 		fmt.Println("PRESSED LEFT")
// 		fmt.Println(entity.speed, "SPEED")
// 		fmt.Println(pos.x, "X POSITION")
// 		if pos.x -(entity.sr.width/2) > 0 {
// 			pos.x -= entity.speed
// 		}
// 	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
// 		fmt.Println("PRESSED RIGHT")
// 		if pos.x+(entity.sr.width/2) < windowWidth {
// 			pos.x += entity.speed
// 		}
// 	}

// 	return nil
// }

func (entity *kbInput) draw(renderer *sdl.Renderer) error {
	return nil
}
