package main

import (
"net/http"
"io/ioutil"
"time"
"fmt"
"net"
"encoding/json"
"bytes"
)

var HttpTimeOut = 30 * time.Second
type Tag_HttpResponse struct {
	StatusCode int `json:"statusCode"`
	Body string	 `json:"body"`
	During int  `json:"during"`
}

type USER struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Sex int `json:"sex"`
}


func Get(url string) Tag_HttpResponse {
	start := time.Now()
	var netResponse = Tag_HttpResponse{200,"",0}
	resp,err := http.Get(url)
	if err != nil{
		netResponse.StatusCode = 400
		netResponse.Body = err.Error()
		return netResponse
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	elapsed := int(time.Since(start).Seconds()*1000)
	if err != nil {
		netResponse.StatusCode = 400
		netResponse.Body = err.Error()
		netResponse.During = elapsed
		return netResponse
	}
	netResponse.Body = string(body)
	netResponse.During = elapsed
	return netResponse
}




func HttpGet(url string )Tag_HttpResponse{

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(network, addr, HttpTimeOut)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(HttpTimeOut))
				return conn, nil
			},
			ResponseHeaderTimeout: HttpTimeOut,
		},
	}

	start := time.Now()
	response,err := client.Get(url)
	elapsed := int(time.Since(start).Seconds()*1000)


	if err != nil {
		errMsg := fmt.Sprintf("%v", err )
		return Tag_HttpResponse{
			StatusCode: 400,
			Body : string(errMsg),
			During : elapsed,
		}
	}

	body,err := ioutil.ReadAll(response.Body)
	if err != nil {
		errMsg := fmt.Sprintf("%v", err )
		return Tag_HttpResponse{
			StatusCode: 400,
			Body : string(errMsg),
			During : elapsed,
		}
	}

	return Tag_HttpResponse{
		StatusCode: response.StatusCode,
		Body : string(body),
		During : elapsed,
	}
}

func HttpPost(url string, form interface{} )Tag_HttpResponse{
	start := time.Now()
	v, _ := json.Marshal(form)
	request, err := http.NewRequest("POST", url, bytes.NewReader(v))
	if err != nil {
		elapsed := int(time.Since(start).Seconds()*1000)
		errMsg := fmt.Sprintf("%v", err )
		return Tag_HttpResponse{
			StatusCode: 400,
			Body : string(errMsg),
			During : elapsed,
		}
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	request.Header.Set("Connection", "close")

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(network, addr, HttpTimeOut)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(HttpTimeOut))
				return conn, nil
			},
			ResponseHeaderTimeout: HttpTimeOut,
		},
	}

	response,err := client.Do(request)

	if err != nil {
		elapsed := int(time.Since(start).Seconds()*1000)
		errMsg := fmt.Sprintf("%v", err )
		return Tag_HttpResponse{
			StatusCode: 400,
			Body : string(errMsg),
			During : elapsed,
		}
	}
	defer response.Body.Close()
	body,err := ioutil.ReadAll(response.Body)
	if err != nil {
		elapsed := int(time.Since(start).Seconds()*1000)
		errMsg := fmt.Sprintf("%v", err )
		return Tag_HttpResponse{
			StatusCode: 400,
			Body : string(errMsg),
			During : elapsed,
		}
	}
	elapsed := int(time.Since(start).Seconds()*1000)
	return Tag_HttpResponse{
		StatusCode: response.StatusCode,
		Body : string(body),
		During : elapsed,
	}
}



//func main(){
//	//url := "http://127.0.0.1:3333/200"
//	//url := "http://localhost:3333/200"
//	url := "http://www.dcl.com:3333/200"
//	//url := "http://www.baidu.com"
//	//form := USER{
//	//	Name:"dingchaolin",
//	//	Age: 23,
//	//	Sex: 1,
//	//}
//	//json, _ := json.Marshal( form )
//	//fmt.Println( string(json) )
//	//ret := HttpPost( url, form )
//	user := USER{
//		Name:"",
//		Age:0,
//		Sex:0,
//	}
//	ret := HttpGet( url )
//	json.Unmarshal( []byte(ret.Body), &user)
//	fmt.Println( "ret ========", ret , user, user.Name)
//}

 func Post( url string, form interface{}, ch chan Tag_HttpResponse ){
	 ch <- HttpPost( url, form )
 }
func main(){
	url := "http://127.0.0.1:3333/200"

	form := USER{
		Name:"dingchaolin",
		Age: 23,
		Sex: 1,
	}

	ch := make( chan Tag_HttpResponse  )

	go Post( url, form, ch )

	res := <- ch

	fmt.Println( "return====", res)


}