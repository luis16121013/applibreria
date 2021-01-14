package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlReponsitory struct {
	Repo *sql.DB
}

//return new MysqlReponsitory
func NewMyslRepository() (*MysqlReponsitory, error) {
	db, err := sql.Open("mysql", RouteConnection())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &MysqlReponsitory{db}, nil
}

//return rows database
func (r *MysqlReponsitory) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := r.Repo.Query(query, args...)
	if err != nil {
		log.Println(err)
	}
	return rows, err
}

//execute consult database
func (r *MysqlReponsitory) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := r.Repo.Exec(query, args...)
	if err != nil {
		log.Println(err)
	}
	return result, err
}
