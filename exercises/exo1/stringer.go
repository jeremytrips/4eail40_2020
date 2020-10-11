package exo1

import (
	"fmt"
)

type Shape interface {
	String() string
	fmt.Stringer // Permet d'implémenter l'interface String sur les type shape
}

type Rectangle struct {
	width, length int
}

type Circle struct {
	radius int
}

func (i Circle) String() string {
	return fmt.Sprintf("Circle of radius %d", i.radius)
}

func (i Rectangle) String() string {
	return fmt.Sprintf("Square of width %d and length %d", i.width, i.length)
}

func main() {
	r := &Rectangle{2, 3}
	c := &Circle{5}

	shapes := []Shape{r, c}

	for _, s := range shapes {
		fmt.Println(s)
	}

}
