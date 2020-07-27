package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

var elements []*element

// Position in space
type vector struct {
	x, y float64
}

// every component requires these functions
type component interface {
	update() error
	draw(renderer *sdl.Renderer) error
	onCollision(other *element) error
}

// container of state for element, aka entity
type element struct {
	position   vector
	rotation   float64
	active     bool
	tag        string
	collisions []circle
	components []component
}

// If element doesn't already hold component, add it to the slice
// remember just a component is just a pointer
func (e *element) addComponent(c component) {
	for _, existing := range e.components {
		if reflect.TypeOf(c) == reflect.TypeOf(existing) {
			// This should be known at compile time.
			panic(fmt.Sprintf("Attempted to add new component with existing type %v", reflect.TypeOf(c)))
		}
	}
	e.components = append(e.components, c)
}

// getComponent returns component of type provided.
func (e *element) getComponent(targetType component) component {
	typ := reflect.TypeOf(targetType)

	for _, c := range e.components {
		if reflect.TypeOf(c) == typ {
			return c
		}
	}

	panic(fmt.Sprintf("Attemped to retrieve component of type %v. This component does not exist.", typ))
}

func (e *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range e.components {
		err := comp.draw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *element) update() error {
	for _, comp := range e.components {
		err := comp.update()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *element) collision(other *element) error {
	for _, comp := range e.components {
		err := comp.onCollision(other)
		if err != nil {
			fmt.Println("Error occured while checking collisions. Error: ", err)
		}
	}
	return nil
}
