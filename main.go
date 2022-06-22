package main

import "fmt"

//Abstraccion

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserID  []uint
	Classes map[uint]string
}

func main() {

	/*
		Go := Course{
			Name:   "go desde 0",
			Price:  12.34,
			IsFree: false,
			UserID: []uint{12, 56, 89},
			Classes: map[uint]string{
				1: "introduccion",
				2: "Variables"},
		}
	*/
	Go := Course{
		//Tambien podemos enviarlos en el mismo orden sin especificar el nombre
		"go desde 0",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "introduccion",
			2: "Variables"},
	}

	fmt.Println(Go.Name)
	Css := Course{Name: "css desde 0"}
	fmt.Println(Css)

}
