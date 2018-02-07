package demo

import (
	"fmt"
	"syscall"
)

func main(){

	s1 := make( []int, 3, 10)//default  len == cap
    fmt.Println( s1, len(s1), cap(s1) )

    a := []byte{ 1,2,3,4,5,6,7,8,9,10}
    sa := a[2:5]// len = 3 cap=10 cap是到内存块的尾部 reslice

    fmt.Println( a, sa );






 }