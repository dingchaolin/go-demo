package demo

import (
	"fmt"
	"time"
)



func main(){
	t := time.Now()
	//直接使用常量或者拷贝值
	fmt.Println( t.Format(time.ANSIC) )



}


