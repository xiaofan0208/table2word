package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlTable struct {
	Field      sql.NullString // 字段
	Type       sql.NullString // 类型
	Collation  sql.NullString // 排序字符集
	Null       sql.NullString // 是否Null
	Key        sql.NullString
	Default    sql.NullString // 默认值
	Extra      sql.NullString
	Privileges sql.NullString
	Comment    sql.NullString // 描述
}

func NewMySqlTable() *MySqlTable {
	return &MySqlTable{}
}

func (s *MySqlTable) Query(d *DBTable) (map[string][]DataTable, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", d.Username, d.Password, d.Address, d.Port, d.Db)
	db, err := sql.Open("mysql", source)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		fmt.Println("连接失败")
		return nil, err
	}

	// 查询所有表
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tablesName string
		err = rows.Scan(&tablesName)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		tables = append(tables, tablesName)
	}

	var dataTables = make(map[string][]DataTable, 0)
	for _, table := range tables {
		res, err := db.Query("SHOW FULL COLUMNS FROM " + table + ";")
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		for res.Next() {
			err = res.Scan(&s.Field, &s.Type, &s.Collation, &s.Null,
				&s.Key, &s.Default, &s.Extra, &s.Privileges, &s.Comment)
			if err != nil {
				fmt.Println(err.Error())
				return nil, err
			}

			dataTables[table] = append(dataTables[table], DataTable{
				Table:   table,
				Field:   s.Field.String,
				Type:    s.Type.String,
				Null:    s.Null.String,
				Comment: s.Comment.String,
			})

		}
	}

	return dataTables, nil
}
