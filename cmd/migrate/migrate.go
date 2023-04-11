package migrate

import "gorm.io/gorm"

var MysqlDB *gorm.DB

// SetMigrateDB sets the mysqlDB instance for the cmd package
func SetMigrateDB(mysql *gorm.DB) {
	MysqlDB = mysql
}
