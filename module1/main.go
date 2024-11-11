package main

import (
	"fmt"

	_ "github.com/Davinia/cursoGo/module1/package1"
)

func init() {
	fmt.Println("You are in init from package main")
}

func main() {
	fmt.Println("You are in main entry point from package main")
	fmt.Println("Welcome")
	a := 10
	b := 20
	c := 30
	d := 40
	e := 50
	f := 60
	fmt.Println(a + b + c + d + e + f)
}
