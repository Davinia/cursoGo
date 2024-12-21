package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	y     = "\033[93m"
	g     = "\033[92m"
	b     = "\033[33m"
	r     = "\033[91m"
	reset = "\033[0m"
)

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func decodeCharacter(encoded []int) string {
	return string(encoded[0])
}

func decodeMessage(encoded [][]int) string {
	message := ""

	for _, line := range encoded {
		currentChar := line[0]
		message += string(currentChar)
		for _, shift := range line[1:] {
			currentChar += shift
			message += string(currentChar)
		}
		message += "\n"
	}
	return message
}
func main() {

	encodedMessage := [][]int{
		{32, 0, 0, 14, 65, -55, 0, 55, -65, -14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 79, 0, 0, 0, -79, 0, 0, 79, -55, 55, -79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 79, 0, 0, 0, 0, -79, 0, 0, 0, 0, 0, 79, 0, 0, -79, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 79, -55, 55, -79, 0, 0, 0, 0, 0, 0, 0, 14, 65, -55, -24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 14, 65, -55, -24, 0, 0, 0, 0, 14, 65, -65, -14, 14, 65, -65},
		{32, 0, 0, 24, 0, 0, -24, 2, 62, -64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, -40, 0, 0, -24, 0, 0, 64, -62, 5, -7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, -40, 0, 0, 42, -52, -14, 0, 0, 0, 0, 64, -40, -17, -7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, -62, 5, -7, 0, 0, 0, 0, 0, 0, 2, 22, 0, 0, -24, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 22, 0, 0, -24, 0, 0, 0, 0, 24, 0, 0, -24, 24, 0, 0},
		{32, 0, 79, -55, 0, 0, 55, 0, -79, 0, 14, 65, 0, 0, 0, 0, -65, -14, 0, 0, 24, 0, 0, -24, 0, 79, 0, 0, 0, -79, 0, 0, 0, 79, 0, 0, 0, 0, 0, 0, 0, -79, 0, 0, 0, 0, 0, 24, -24, 64, -40, 0, 42, -52, -14, 0, 0, 0, 24, -24, 0, 0, 14, 65, 0, 0, 0, -65, -14, 0, 0, 79, 0, 0, 0, -79, 0, 0, 79, 0, 0, -79, 79, 0, 0, 0, -79, 0, 0, 14, 65, 0, 0, 0, -55, 0, 0, -24, 0, 0, 14, 65, 0, 0, 0, -65, -14, 0, 0, 0, 14, 65, 0, 0, 0, -55, 0, 0, -24, 0, 0, 0, 0, 24, 0, 0, -24, 24, 0, 0},
		{32, 0, 0, 24, 0, 0, -24, 0, 0, 68, -44, 0, -17, -7, 64, -40, 0, 42, -66, 0, 24, 0, 0, -24, 0, 64, -40, 0, 0, -24, 0, 0, 68, -61, -5, 0, 21, 45, -44, 24, -48, 0, 0, 0, 0, 0, 0, 24, -24, 0, 0, 64, -40, 0, 42, -52, -14, 0, 24, -24, 0, 64, -16, -48, 0, 9, 15, 0, 42, -66, 0, 0, 64, -40, 0, -10, -14, 14, 10, -17, -7, 0, 64, -40, 0, 0, -24, 0, 68, -44, 0, -17, -7, 64, -40, 0, 0, -24, 0, 64, -16, -48, 0, 9, 15, 0, 42, -66, 0, 68, -44, 0, -17, -7, 64, -40, 0, 0, -24, 0, 0, 0, 0, 57, -33, 24, -48, 57, -33, 24},
		{32, 0, 0, 24, 0, 0, -24, 0, 0, 24, 0, 0, 55, 0, 0, -55, 0, 0, -24, 0, 24, 0, 0, -24, 0, 0, 24, 0, 0, -24, 0, 0, 0, 0, 14, 54, -44, 24, -41, -7, 0, 0, 0, 0, 0, 0, 0, 24, -24, 0, 0, 0, 0, 64, -40, 0, 42, -52, 10, -24, 0, 0, 14, 65, -31, -46, 22, 0, 0, -24, 0, 0, 0, 64, -40, 0, -10, 0, 10, -17, -7, 0, 0, 24, 0, 0, -24, 0, 24, 0, 0, -24, 0, 0, 24, 0, 0, -24, 0, 0, 14, 65, -31, -46, 22, 0, 0, -24, 0, 24, 0, 0, -24, 0, 0, 24, 0, 0, -24, 0, 0, 0, 0, 64, -40, -17, -7, 64, -40, -17},
		{32, 0, 0, 24, 0, 0, -24, 0, 0, 24, 0, 0, -24, 0, 0, 0, 14, 65, -79, 0, 24, 0, 0, -24, 0, 0, 24, 0, 0, -24, 0, 0, 14, 54, -44, 24, -41, -7, 14, 34, -48, 0, 0, 0, 0, 0, 0, 24, -24, 0, 0, 0, 0, 0, 0, 64, -40, 0, 0, -24, 0, 68, -44, -16, -8, 0, 24, 0, 0, -24, 0, 0, 0, 0, 64, -40, 0, 0, -17, -7, 0, 0, 0, 24, 0, 0, -24, 0, 24, 0, 0, -24, 0, 0, 24, 0, 0, -24, 0, 68, -44, -16, -8, 0, 24, 0, 0, -24, 0, 24, 0, 0, -24, 0, 0, 24, 0, 0, -24, 0, 0, 0, 0, 14, 65, -65, -14, 14, 65, -65},
		{32, 0, 79, -55, 0, 0, 55, -79, 0, 64, -7, -33, 42, 13, -11, -44, 24, -41, -7, 79, -55, 0, 0, 55, -79, 79, -55, 0, 0, 55, -79, 68, -44, 0, 0, 0, 0, 0, 24, -48, 0, 0, 0, 0, 0, 0, 79, -55, 55, -79, 0, 0, 0, 0, 0, 0, 0, 64, -40, -24, 0, 64, -7, -33, 0, 0, -22, 0, 22, 55, -79, 0, 0, 0, 0, 64, -40, -17, -7, 0, 0, 0, 79, -55, 0, 0, 55, -79, 64, -7, -33, 42, 13, -11, -44, 0, 24, -46, -2, 64, -7, -33, 0, 0, -22, 0, 22, 55, -79, 64, -7, -33, 42, 13, -11, -44, 0, 24, -46, -2, 0, 0, 0, 57, -33, 24, -48, 57, -33, 24},
	}

	clearScreen()

	height := 35
	maxWidth := 2*height - 1

	baseWidth := height / 3
	if baseWidth < 1 {
		baseWidth = 1
	}

	encodedCharacter := []int{42}

	tChar := decodeCharacter(encodedCharacter)

	fmt.Println()

	spaces := height - 1 + 25
	fmt.Print(strings.Repeat(" ", spaces) + y + tChar + reset + "\n")

	for i := 1; i < height; i++ {
		lineWidth := 2*i + 1
		spaces = height - i - 1 + 25
		fmt.Print(strings.Repeat(" ", spaces) + g)
		for j := 0; j < lineWidth; j++ {
			fmt.Print(tChar)
		}
		fmt.Print(reset + "\n")
	}

	baseSpaces := ((maxWidth - baseWidth) / 2) + 25
	for i := 0; i < height/3; i++ {
		fmt.Print(strings.Repeat(" ", baseSpaces) + b + strings.Repeat("|", baseWidth) + reset + "\n")
	}

	message := decodeMessage(encodedMessage)
	fmt.Println("\n" + r + message + reset)

}