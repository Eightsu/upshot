package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	// container is the element associated with said sprite
	container *element
	tex *sdl.Texture
}

// HELPERS
func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("Loading %v: %v", filename, err))
	}

	defer img.Free()

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Creating texture from %v: %v", filename, err))
	}
	return tex
}



func newSpriteRenderer(container *element ,renderer *sdl.Renderer,filename string) *spriteRenderer {
	return &spriteRenderer{
		container: container,
		tex : textureFromBMP(renderer,filename),

	}


}

func (sr *spriteRenderer) Draw(renderer *sdl.Renderer) error {

	_, _,width,height, err := sr.tex.Query()
	if err != nil {
		return fmt.Errorf("Querying %v caused an error: ", err)
	}

	x := sr.container.position.x - float64(width)/2.0
	y := sr.container.position.y -float64( height) /2.0
	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, width, height},
		&sdl.Rect{X: int32(x), Y: int32(y), W: width, H: height },
		sr.container.rotation,
		&sdl.Point{X: width / 2, Y: height / 2},
		sdl.FLIP_NONE
	)

	return nil
}

// To do - An update function isn't necessary, for anything besides the contract to be a component
func (sr *spriteRenderer) Update(renderer *sdl.Renderer) error {
	return nil
}



