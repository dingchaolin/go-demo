package demo

import (
	"fmt"
)




 func main(){

	 //a := 1
	 //switch a {
		// case 0:
		//	fmt.Println( 0 )
		// case 1:
		//	fmt.Println( 1 )
		// default:
		//	 fmt.Println( "NONE")
	 //}

	 //switch {
	 //case a >= 0:
		// fmt.Println( 0 )
		// fallthrough //可继续执行下面的case
	 //case a >= 1:
		// fmt.Println( 1 )
	 //default:
		// fmt.Println( "NONE")
	 //}

	 switch b := 1; {//b 在switch中有效 需要; 否在报错
	 case b >= 0:
		 fmt.Println( 0 )
		 fallthrough //可继续执行下面的case
	 case b >= 1:
		 fmt.Println( 1 )
	 default:
		 fmt.Println( "NONE")
	 }



 }