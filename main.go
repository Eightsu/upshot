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
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	p := newPlayer(renderer)
	elements = append(elements, p)

	// TODO There's something weird happening when we draw enemies to the screen. For now, we'll leave the magic numbers.
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			x := (float64(i)/5)*windowWidth + enemySize/2 - 16
			y := float64(j) * enemySize / 2
			e := newEnemy(renderer, vector{x, y + (enemySize / 2)})
			elements = append(elements, e)
		}
	}

	initBPool(renderer)
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
		renderer.SetDrawColor(250, 240, 230, 255)
		renderer.Clear()
		for _, elem := range elements {
			if elem.active {
				err := elem.update()
				if err != nil {
					fmt.Println("An error occurred while updating elements. Error: ", err)
					return
				}

				err = elem.draw(renderer)
				if err != nil {
					fmt.Println("An error occured while updating the elements. Error: ", err)
					return
				}
			}
		}
		renderer.Present()
	}

}
