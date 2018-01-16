package main

import "fmt"

const N = 10
/*
async.WaitGroup 也能实现类似的效果
 */
func doSometing(index int, val int) {
	fmt.Println(index, val)
}

/*
不能确定哪个协程先执行结束
 */
func main() {
	sem := make(chan int, N)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, val := range arr {
		go func(index int, val int) {
			doSometing(index, val)
			sem <- 0
		}(index, val)
	}

	/*
	等待循环结束
	 */
	for i := 0; i < N; i ++ {
		<- sem
	}


}
