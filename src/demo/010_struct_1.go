package demo

import (
	"fmt"
)

type person struct{
	Name string
	Age int

}

func main(){

	a := person{}
	a.Name = "dcl"
	a.Age = 13

	/*
	对结构初始化的时候使用&  还可以使用.操作符对字段进行操作
	 */
	b := &person{
		Name: "ys",
		Age: 14,
	}

	b.Name = "yyy";

	var d person
	d = a //可以比较
	fmt.Println( a, b, d, a == d )

}

func AA( p person ){
	p.Age = 15;
	fmt.Println("AA", p );
}

func AAp( p *person ){
	p.Age = 15;
	fmt.Println("AA", p );
}