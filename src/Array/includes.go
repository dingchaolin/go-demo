package main

func main(){
	arr := []string{"aa","bb","bb","dd"}
	println( indexOf( arr, "bb"))

}

func indexOf( arr [] string, findStr string ) int {
	index := -1
	for i := 0; i < len(arr); i++{
		if arr[i] == findStr {
			index = i
			break
		}
	}
	return index
}