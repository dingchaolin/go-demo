package demo

import (
	"fmt"
	"sort"
)

func main(){


	m := map[int]string {1:"1", 2:"2", 3:"3", 4:"4", 5:"5"}
	s := make([]int, len(m))

	i := 0

	for key,_ := range m{
		s[i] = key
		i++
	}
	sort.Ints( s )//说明slice是个引用  ints 是对s的本身的操作 Ints 对int类型进行排序
	fmt.Println( s )


}