package factory

import "github.com/DanoFP/dp-golang-training/pd/01-factory/connection"

func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &connection.MySql{}
	case 2:
		return &connection.Postgres{}
	default:
		return nil
	}
}
