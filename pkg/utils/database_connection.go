package utils

import (
	"fmt"

	"github.com/kasfil/offspace/pkg/schemas"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnect connect and migrate database schema
func DBConnect() *gorm.DB {
	host := viper.GetString("database_host")
	port := viper.GetInt("database_port")
	user := viper.GetString("database_user")
	pass := viper.GetString("database_pass")
	name := viper.GetString("database_name")
	charset := viper.GetString("database_charset")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", user, pass, host, port, name, charset)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

// RunMigration run migration
func RunMigration(db *gorm.DB) {
	// running all database model schema database
	db.Set("gorm.table_options", "ENGINE=InnoDB").AutoMigrate(&schemas.User{})
}
