package packageTranslator

import "fmt"

func init() {
}

func SaludarEn(idioma string) {
	var saludo string
	switch idioma {
	case "Japonés":
		saludo = "こんにちは世界！"
	case "Árabe":
		saludo = "مرحبا بالعالم!"
	case "Español":
		saludo = "¡Hola mundo!"
	case "Inglés":
		saludo = "Hello world!"
	case "Francés":
		saludo = "Bonjour le monde!"
	default:
		saludo = "¡Hola mundo!"
	}
	saludoEnRunas := []rune(saludo)
	fmt.Println(saludoEnRunas)
	fmt.Println("¡Es broma! Aquí tienes tu saludo:", string(saludoEnRunas))
}
