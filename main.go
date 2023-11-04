package main

import (
	"fmt"
	"table2word/db"
)

func main() {

	t := db.NewDBTable("mysql", "root", "0208", "127.0.0.1", "3306", "goadmin")
	dataTables, err := t.Connect()
	if err != nil {
		panic(err)
	}
	for tableName, tableInfo := range dataTables {
		fmt.Println(tableName)
		fmt.Println(tableInfo)
	}
}
