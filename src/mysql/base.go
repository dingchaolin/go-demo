package mysql

import (
	"database/sql"
	"fmt"
)

func Insert(tableName string, rowMap map[string]interface{}) (int64, error) {
	sql := "insert into `" + tableName + "` ( "
	var argArray []interface{}
	fieldStr := ""
	valueStr := ""
	indexNum := 0
	/*
	key-value 必须在一次循环中取出
	 */
	for key, value := range rowMap {
		if indexNum == 0 {
			fieldStr = fieldStr + "`" + key + "`"
			valueStr = valueStr + "?"
		} else {
			fieldStr = fieldStr + ",`" + key + "`"
			valueStr = valueStr + ",?"
		}
		argArray = append(argArray, value)
		indexNum++
	}
	sql = sql + fieldStr + ") values (" + valueStr + ")"
	result, err := mysqlConn.Exec(sql, argArray...)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertId, nil
}

func FetchList(tableName string, fields string, whereMap map[string]interface{}) (*sql.Rows, error) {
	if fields == "" {
		fields = "*"
	}
	indexNum := 0
	var argArray []interface{}
	whereStr := ""
	for key, value := range whereMap {
		if indexNum == 0 {
			whereStr = whereStr + " where `" + key + "` = ? "
		} else {
			whereStr = whereStr + "And `" + key + "` = ? "
		}
		argArray = append(argArray, value)
		indexNum++
	}
	sql := "select " + fields + " from " + tableName + whereStr
	rows, err := mysqlConn.Query(sql, argArray...)
	fmt.Println( fmt.Sprintf("%v,%v", rows, err))
	if err != nil {
		return nil, err
	}
	return rows, err

}

func Update(tableName string, fieldsMap map[string]interface{}, whereMap map[string]interface{}) (int64, error) {

	indexNum := 0
	var argArray []interface{}
	setStr := ""

	for key, value := range fieldsMap {
		if indexNum == 0 {
			setStr = setStr + " `" + key + "` = ? "
		} else {
			setStr = setStr + ", `" + key + "` = ? "
		}
		argArray = append(argArray, value)
		indexNum++
	}
	indexNum = 0
	whereStr := ""
	for key, value := range whereMap {
		if indexNum == 0 {
			whereStr = whereStr + " where `" + key + "` = ? "
		} else {
			whereStr = whereStr + "And `" + key + "` = ? "
		}
		argArray = append(argArray, value)
		indexNum++
	}

	sql := "update " + tableName + " set " + setStr + whereStr
	_, err := mysqlConn.Exec(sql, argArray)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func Count(tableName string, whereMap map[string]interface{})(int64,error){
	rows, err := FetchList( tableName, "count(1) as num", whereMap)
	if err != nil{
		return 0, err
	}
	num := int64(0)
	for rows.Next() {
		rows.Scan(&num)
		break
	}
	return num, nil
}