/**
 * Created by chaolinding on 2018/4/2.
 */

package main

import (
	"reflect"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"`
}

func main() {
	user := User{Name: "dcl", Age: 14, Sex: true,}
	vt := reflect.TypeOf(user)
	vv := reflect.ValueOf(user)
	userMap := map[string]interface{}{}

	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		key := f.Tag.Get("json")
		value := vv.FieldByName(f.Name).Interface()
		fmt.Printf("%v => %v \n", key, value)
		userMap[key] = value
	}

	fmt.Println("map====",userMap)

}
