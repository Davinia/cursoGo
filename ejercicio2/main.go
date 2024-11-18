package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func init() {
}

func main() {

	prueba := flag.String("test", "bytes", "Prueba la función RandRunes usando: bytes o runes como entrada")

	longitud := flag.Int("longitud", 10, "Longitud de la cadena a generar")

	flag.Parse()

	fmt.Println("Prueba elegida:", *prueba)

	switch *prueba {
	case "bytes":
		sliceDeBytes := []byte("こんにちは世界！")
		fmt.Printf("Se van a extraer %d elementos de este input: %s\n", *longitud, string(sliceDeBytes))
		fmt.Println(RandRunesWithBytes(*longitud, sliceDeBytes))
	case "runes":
		sliceDeRunas := []rune("こんにちは世界！")
		fmt.Printf("Se van a extraer %d elementos de este input: %s\n", *longitud, string(sliceDeRunas))
		fmt.Println(RandRunesWithRunes(*longitud, sliceDeRunas))
	}
}
func RandRunesWithBytes(totalLength int, input []byte) (outputString string) {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var output []byte = make([]byte, totalLength)

	for i := 0; i < totalLength; i++ {
		posiciónAleatoria := rand.Intn(len(input) - 1)
		output[i] = input[posiciónAleatoria]
	}
	outputString = string(output)
	return
}

func RandRunesWithRunes(totalLength int, input []rune) (outputString string) {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var output []rune = make([]rune, 0, totalLength)

	for i := 0; i < totalLength; i++ {
		posiciónAleatoria := rand.Intn(len(input) - 1)
		output = append(output, input[posiciónAleatoria])
	}
	outputString = string(output)
	return
}
