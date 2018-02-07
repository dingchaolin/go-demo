package demo

import (
	"fmt"
)

func main(){

    FA()
    FB()
    FC()

}

func FA(){
	fmt.Println("func A")
}

func FB(){

	defer func(){
		if err:=recover(); err != nil {
			fmt.Println( "recover in b")
		}
	}()
	//如果没有defer recover FC就不能正常执行
	panic("panic in B")
}

func FC(){
	fmt.Println("func C")
}
