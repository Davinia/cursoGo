package packageTranslator

import "fmt"

func init() {
}

func SaludarEn(idioma string) {
	var saludo rune
	switch idioma {
	case "Japonés":
		saludo := 'こんにちは世界！'
	case "Árabe":
		saludo := 'مرحبا بالعالم!'
	case "Español":
		saludo := '¡Hola mundo!'
	case "Inglés":
		saludo := 'Hello world!'
	case "Francés":
		saludo := 'Bonjour le monde!'
	default:
		saludo := '¡Hola mundo!'
	}
	fmt.Println(saludo)
}
