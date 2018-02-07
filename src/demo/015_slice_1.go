package demo

import (
	"fmt"
)


func _append( s []int )[]int{
	s = append( s, 3)
	return s
}
func main(){
	s := make([]int,0)
	fmt.Println( s )
	//append 一定要对s重新赋值
	s = _append( s )
	fmt.Println( s )



}


