package main

import (
	"fmt"
	"github.com/Davinia/cursoGo/module1package1"
)

func init() {
	fmt.Println("You are in init from package main")
}

func main() {
	fmt.Println("You are in main entry point from package main")
	fmt.Println("Welcome")
}
