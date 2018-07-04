package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

func (s *Saiyan) Super() {
	// Add Suoer() method to Saiyan structure
	s.Power += 10000
}

func main() {

	goku := Saiyan{"Goku", 9000}
	Super(goku)
	fmt.Println(goku.Power) // 9000

	goku_pointer := &Saiyan{"Goku Pointer", 9000}
	SuperPointer(goku_pointer)
	fmt.Println(goku_pointer.Power) // 19000

	goku_pointer.Super()
	fmt.Println(goku_pointer.Power) // 19000

}

func Super(s Saiyan) {
	s.Power += 10000
}

func SuperPointer(s *Saiyan) {
	s.Power += 10000
}
