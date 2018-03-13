package logrecoder

import (
	"errors"
	"log"
	)

func Recoder( msg string){
	log.SetPrefix(`[TEST] `)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile )
	err := errors.New(msg)
	log.Println(err)//错误显示就是这一行的位置
}