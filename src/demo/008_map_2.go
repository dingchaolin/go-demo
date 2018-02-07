package demo

import (
	"fmt"
)

func main(){

	var m map[int]map[int]string

	m = make(map[int]map[int]string)



	m[1] = make( map[int]string)

	m[1][1] = "ok"
	a := m[1][1]
	a, ok := m[1][1] //ok表示是否存在 true false
	if ok {
		m[1] = make( map[int]string)
		m[1][1] = "good"
	}
	//判断一个map是不是被初始化

	fmt.Println( m, a, ok );

}