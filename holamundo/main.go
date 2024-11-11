package main

import (
	"flag"

	"github.com/Davinia/cursoGo/holamundo/packageTranslator"
)

func init() {
}

func main() {
	//fmt.Println("¡Hola mundo!")
	//fmt.Println("¿Quieres que te saludemos en otro idioma?..Dinos cuál:")
	//fmt.Println("- Español")
	//fmt.Println("- Inglés")
	//fmt.Println("- Francés")
	//fmt.Println("- Japonés")
	//fmt.Println("- Árabe")

	idioma := flag.String("idioma", "Español", "Idioma en el que quieras que te saludemos (Español, Inglés, Francés, Japonés, Árabe)")

	flag.Parse()
	packageTranslator.SaludarEn(idioma)
}
