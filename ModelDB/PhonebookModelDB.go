
package ModelDB

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type PhonebookModel struct {
 gorm.Model
 Name  string
 Address string
 phonenumber uint64
}
 

func UserModelDBMigrationHandler() bool {

    db := UserModelDBCreateConnection()

    db.AutoMigrate(&PhonebookModel{})

    defer UserModelDBCloseConnection(db)

    return true
    
}




