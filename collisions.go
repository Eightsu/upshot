package main

import (
	"math"
)

type circle struct {
	center vector
	radius float64
}

// d=√((x_2-x_1)²+(y_2-y_1)²)
func isCollision(c1 circle, c2 circle) bool {
	distance := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) + math.Pow(c2.center.y-c1.center.y, 2))

	return distance <= c1.radius+c2.radius
}

// FIXME there must be a better way?
func checkCollisions() error {
	for i := 0; i < len(elements)-1; i++ {
		for j := i + 1; j < len(elements); j++ {
			for _, c1 := range elements[i].collisions {
				for _, c2 := range elements[j].collisions {
					if isCollision(c1, c2) && elements[i].active && elements[j].active {
						err := elements[i].collision(elements[j])
						if err != nil {
							return err
						}
						err = elements[j].collision(elements[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}

	}
	return nil
}
