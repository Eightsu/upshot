package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	// container is the element associated with said sprite
	container     *element
	tex           *sdl.Texture
	width, height float64
}

// textureFromBMP does exactly that. Returns a pointer to sdl.Texture from a BMP
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

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {

	tex := textureFromBMP(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("Querying %v caused an error: ", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       tex,
		width:     float64(width),
		height:    float64(height),
	}
}
// TODO fill out the method.
func (sr *spriteRenderer) onCollision(e *element)error {
	return nil
}

func (sr *spriteRenderer) draw(renderer *sdl.Renderer) error {

	// convert coords to top left of sprite
	x := sr.container.position.x - sr.width/2.0
	y := sr.container.position.y - sr.height/2.0

	// convert x,y from floats to ints
	// main character needs no rotation. default is zero.
	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(sr.width), H: int32(sr.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(sr.width), H: int32(sr.height)},
		sr.container.rotation,
		&sdl.Point{X: int32(sr.width) / 2, Y: int32(sr.height) / 2},
		sdl.FLIP_NONE,
	)
	return nil
}

// TODO Does this function require anything? Or is it just satisfying the contract?
func (sr *spriteRenderer) update() error {
	return nil
}
