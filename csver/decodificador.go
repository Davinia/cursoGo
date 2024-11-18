package main

import (
	"bufio"
	"os"
)

func init() {

}

func AbrirCSV(nombreCSV string) (*bufio.Scanner, error) {

	var desc *os.File
	var escáner *bufio.Scanner

	desc, err := os.Open( nombreCSV )
	if err != nil{
		return nil, err
	}else{
		escáner = bufio.NewScanner(desc)
		defer desc.Close()
	}
	return escáner, nil
}
