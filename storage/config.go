package storage

import "fmt"

var (
	Username string = "root"
	Password string = "rootpw"
	Host     string = "127.0.0.1"
	Port     int    = 3306
	Database string = "softproyect"
)

//Return Route connection Database
//user:pwd@(tcp(host:port)/database)
func RouteConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", Username, Password, Host, Port, Database)
}
