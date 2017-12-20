package main

import (
	"fmt"
)

func main(){
 	//slice 基本可以理解为变长数组  C++中 vector
    var s1 []int
    fmt.Println( s1 );//[]

    a := [10]int{ 0:1, 5:10}

    s2 := a[9]
    s3 := a[5:10]// 包头不包尾
	s4 := a[5:len(a)]// 包头不包尾
	s5 := a[5:]// 包头不包尾
	s6 := a[:5]// 包头不包尾

    fmt.Println( s2, s3, s4, s5, s6 )




 }