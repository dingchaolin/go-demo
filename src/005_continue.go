package main

import (
	"fmt"
)



// 尽量将标签放到goto语句的后面 能避免死循环
 func main(){

	 LABEL1:
		 for i := 0; i < 5; i++{
			 fmt.Println( i )
			 for {
				 continue LABEL1  //可以跳出死循环 最终会执行fmt.Println( "OK" )
			 }
			 fmt.Println( i )
		 }

		fmt.Println( "OK" )



 }