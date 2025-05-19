package main

import (
	"fmt"
	"math"
)

/* interface contract */
type geometry interface {
    area() float64
    perim() float64
}

/* struct (reciever) that will be implementing the interface*/
type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}


/* interface implementation */

// area and perim methods for rect
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}
// area and perim methods for circle
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

/* function that takes the interface as a parameter */
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

/* type assertion to check if the interface is of a specific type */
func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius)
    }
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)

    detectCircle(r)
    detectCircle(c)
}