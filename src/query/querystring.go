/**
 * Created by chaolinding on 2018/3/30.
 */

package main

import (
	"net/url"
	"fmt"
)

func main(){
	URL := "http://www.baidu.com?name=dcl&age=13&sex=1"
	u, _ := url.Parse(URL)
	fmt.Println( "u=======", u )
	// Parse get query
	query, _ := url.ParseQuery(u.RawQuery)
	fmt.Println( "query======", query,"name===",query["name"][0])

}