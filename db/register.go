package db

import "fmt"

type DBDriver interface {
	Query(d *DBTable) (map[string][]DataTable, error)
}

var (
	drivers = make(map[string]DBDriver)
)

func init() {
	drivers["mysql"] = NewMySqlTable()
}

type DataTable struct {
	Table   string // 表名
	Field   string // 字段
	Type    string // 类型
	Null    string // 是否Null
	Comment string // 描述
}

type DBTable struct {
	DbType   string // 数据库类型
	Username string
	Password string
	Port     string
	Address  string
	Db       string
}

func NewDBTable(dbType, username, password, address, port, db string) *DBTable {
	return &DBTable{
		DbType:   dbType,
		Username: username,
		Password: password,
		Port:     port,
		Address:  address,
		Db:       db,
	}
}

func (d *DBTable) Connect() (map[string][]DataTable, error) {
	if driver, ok := drivers[d.DbType]; ok {
		dataTables, err := driver.Query(d)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return dataTables, nil
	}

	return nil, nil
}
