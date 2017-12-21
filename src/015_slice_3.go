package main

import (
	"fmt"
)



func main(){
	s := []string{"a", "b", "c"}
	//闭包中得到的是地址
	// 传参数 就能得到想要的值
	for _, v := range s{
		go func( v string ){
			fmt.Println( v )
		}( v )
	}

	select {

	}



}


