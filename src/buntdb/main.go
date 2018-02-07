package main

import (
	"buntdb/lib"
	"fmt"
)

func main(){
	//lib.Write( "name", "ys")
	//会读取最后一次设置的值
	name := lib.Read( "name")
	fmt.Println( name )
}
