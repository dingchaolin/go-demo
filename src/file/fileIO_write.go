package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	inputFile := "data.txt"
	outputFile := "out.txt"

	buf, err := ioutil.ReadFile( inputFile )

	if err != nil {
		fmt.Fprintf( os.Stderr, "file error： %s\n", err )
	}

	fmt.Printf("%s\n", string(buf) )

	err = ioutil.WriteFile( outputFile, buf, 0777 )

	if err != nil {
		panic(err.Error())
	}
}
