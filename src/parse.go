package main

import (
	"fmt"
	"parse"
)

func main(){

	var examples = []string{
		"1 2 3 4 5",
		"100 2.3 50 5.7 99",
		"2 + 1 = 3",
		"1st class",
		"",
	}

	for _, ex := range examples{
		fmt.Printf("Parseing %q:\v", ex );
		nums, err := parse.Parse( ex )
		if err != nil{
			fmt.Println( err )
			continue
		}
		fmt.Println( nums )
	}

}


