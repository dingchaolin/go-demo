package logrecoder

import "log"

var LOG log

func init(){
	log.SetPrefix(`[TEST] `)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile )
	LOG = log
}

