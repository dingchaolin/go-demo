package main

import (
	"fmt"
	"import/a"
	"import/b"
)

func main(){
	fmt.Println("a.ID=====", a.ID )
	a.SetId()
	b.PrintId()
}
