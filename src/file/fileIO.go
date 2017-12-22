package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main(){
	inputFile, inputError := os.Open("./data.txt")

	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n" )

		fmt.Println(inputError )
		fmt.Println( filepath.Abs("data.txt") )

		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader( inputFile )

	for{
		inputString, readError := inputReader.ReadString('\n')
		fmt.Printf("the input was: %s", inputString)
		if readError == io.EOF {
			return
		}
	}
}
