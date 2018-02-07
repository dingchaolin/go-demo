package lib

import (
"github.com/tidwall/buntdb"
"fmt"
"encoding/json"
)

var db *buntdb.DB

func init()  {
	filePath := "./data.db"
	btdb,err := buntdb.Open(filePath)
	if err != nil{
		fmt.Println("init cache error==",err.Error())
	}
	db = btdb
}

func Read(key string) string {
	str := ""
	err := db.View(func(tx *buntdb.Tx) error {
		val,err := tx.Get(key)
		if err != nil {
			return err
		}
		str = val
		//fmt.Printf("value is %s\n", val)
		return nil
	})
	if err != nil {
		fmt.Println("Read cache error==",err.Error())
	}
	//fmt.Println("Read")
	return str
}

func Write(key string ,value  interface{})  {
	jsStr,errJson := json.Marshal(value)
	if errJson != nil{
		return
	}
	err := db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, string(jsStr), nil)
		return err
	})
	if err != nil{
		fmt.Println("Write cache error==",err.Error())
	}
}

