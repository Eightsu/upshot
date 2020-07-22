package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 512
	windowHeight = 768
	spriteSize   = 128
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("MainFrame", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// surface, err := window.GetSurface()
	// if err != nil {
	// 	panic(err)
	// }
	// surface.FillRect(nil, 0)

	// rect := sdl.Rect{X: 0, Y: 0, W: 20,H: 20}
	// surface.FillRect(&rect, 0xffff0000)
	// window.UpdateSurface()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		panic(err)
	}

	defer renderer.Destroy()

	p, err := newPlayer(renderer)

	if err != nil {
		panic(err)
	}

	var enemies []enemy

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			x := (float64(i) / 5) * windowWidth - 16
			y := float64(j) * enemySize / 2.0

			e, err := newEnemy(renderer, x, y)
			if err != nil {
				panic(err)
			}
			

			// remember to assign positions
			e.x = x
			e.y = y
			fmt.Println(e.x,e.y)
			enemies = append(enemies, e)
		}
	}

	fmt.Println(len(enemies))


	// Loop so window doesn't close
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit", event.GetTimestamp())
				running = false
				break
			}
		}

		// Bg Color
		renderer.SetDrawColor(250, 240, 230, 255)

		// Updates screen?
		renderer.Clear()

		// Put images on screen
		// Second parameter chooses the size of texture based on img file. Could be used to grab different parts of a spritesheet
		// third parameter could be used in relation to scaling(?)

		p.draw(renderer)
		p.update(renderer)

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}
		renderer.Present()
	}

}
