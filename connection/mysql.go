package connection

import (
	"database/sql"
	"fmt"
	"time"
)

type MySql struct {
	db *sql.DB
}

//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

func (m *MySql) Connect() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&allowNativePassword=true&parseTime=true",
		"root",
		"1234",
		"127.0.0.1",
		"3306",
		"mysql",
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	m.db = db
	return nil
}

func (m *MySql) GetNow() (*time.Time, error) {
	t := &time.Time{}
	err := m.db.QueryRow("select CURDATE()").Scan(t)
	if err != nil {
		fmt.Printf("error to read time of server: %v", err)
		return nil, err
	}
	return t, nil
}
