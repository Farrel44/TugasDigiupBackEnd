package main
import ("fmt")

type person struct{
	name string
}

func (p person) printName() {
	fmt.Println("Nama =", p.name)
}

func main() {
	Person := person{name: "John Doe"}
	Person.printName()
}