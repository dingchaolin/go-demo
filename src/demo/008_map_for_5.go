package demo

import (
	"fmt"
)

func main(){


	m := map[int]string {1:"a", 2:"b", 3:"c"}
	m2 := make( map[string]int)



	for key,v := range m{
		m2[v] = key
	}

	fmt.Println( m2 )


}