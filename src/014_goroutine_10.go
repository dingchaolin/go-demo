package main

import (
	"fmt"
	"time"
)

func mainTest(  ){
	fmt.Println( " 主 test 启动啦")
	go testA()
	go testB()
	fmt.Println( " 主 test 借宿啦")
	return

}

func testA(  ){
	fmt.Println( " Test A 启动啦")
	time.Sleep( time.Second * 3 )
	fmt.Println( " Test A 结束啦")
	return
}

func testB(  ){
	fmt.Println( " Test B 启动啦")
	time.Sleep( time.Second * 5 )
	fmt.Println( " Test B 结束啦")
	return
}

func main(){
	go mainTest( )
	time.Sleep( 10 * time.Second )
}


