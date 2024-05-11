

package ModelDB

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "fmt"

    "github.com/Nabin-Flash320/go_phonebook/Settings"
)


func UserModelDBCreateConnection() *gorm.DB {

	var database_name = Settings.DBSettingsGetName()
    var database_user = Settings.DBSettingsGetUser()
    var database_pass = Settings.DBSettingsGetPass()
    var database_host = Settings.DBSettingsGetHost()
    var database_port = Settings.DBSettingsGetPort()

    var connection_params string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", database_user, database_pass, database_host, database_port, database_name)

    db, err := gorm.Open("mysql", connection_params)

    if err != nil {

        panic("Failed to connect to database")

    }

	return db
}

func UserModelDBCloseConnection(db *gorm.DB) {
	
	db.Close()

}






