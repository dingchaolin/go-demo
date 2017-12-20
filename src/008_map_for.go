package main

import (
	"fmt"
)

func main(){

	// v是一个拷贝 slice遍历
	//for i,v := range slice{
	//	slice[i] = 1
	//}

	//for key,v := range map{
	//	map[key] = "1"
	//}
	// v 是拷贝  对其操作不影响原值
	sm := make( []map[int]string, 5)
	for i,v:= range sm{
		sm[i] = make( map[int]string, 1)
		sm[i][1] = "ok"
		fmt.Println(sm[i], v)
	}


}