package main

import "fmt"

type query struct{
	sql chan string
	result chan string
}

func execQuery( q query ){
	go func(){
		sql := <-q.sql
		/*
		数据库查询
		 */
		 q.result <- "get [ " + sql + " ]"
	}()
}


func main(){
	q := query{
		make(chan string),
		make(chan string),
	}

	execQuery(q)

	q.sql <- "select * from tb"

	fmt.Println( <-q.result )
}