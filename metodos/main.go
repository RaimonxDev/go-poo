package main

func main() {

	Go := Course{
		" Go desde 0",
		20,
		false,
		[]uint{10, 20, 30},
		map[uint]string{
			1: "Introduccion",
			2: "Primeros Pasos",
			3: "Variables",
		},
	}
	//Llamamos al metodo como cualquier otra propiedad
	Go.PrintClasses()
}
