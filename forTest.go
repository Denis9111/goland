package main

import (
	"fmt"
)

type Sum struct {
	a int
	b int
}

func (s Sum) ToSum() {
	fmt.Println(s.a + s.b)
}

func main() {
	sum := Sum{1, 2}
	sum.ToSum()

}
