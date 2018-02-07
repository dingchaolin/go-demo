package demo

import (
	"fmt"
)




 func main(){

	 LABEL1:
	 	for{
	 		for i := 0; i < 5; i++{
	 			if i > 3 {
	 				break LABEL1 //会调到LABEL1 同级处 执行 fmt.Println( "OK" )代码
	 				// goto LABEL1 //将程序调整到LABEL1处  尽量将将标签放到goto之后
	 				//continue LABEL1  死循环
				}
			}
		}

		fmt.Println( "OK" )



 }