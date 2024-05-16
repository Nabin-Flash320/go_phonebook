
package Models


import (
    "strings"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "github.com/Nabin-Flash320/go_phonebook/Core/Database"
    "github.com/Nabin-Flash320/go_phonebook/Core/MessagePassing"
)


type PhonebookModel struct {
    gorm.Model
    Name  string `json:"name" gorm:"not null"`
    Address string `json:"address" gorm:"not null"`
    PhoneNumber uint64  `json:"phonenumber" gorm:"unique;not null"`
    // database constraint enforcement is used to describe the additional information about model column
}

type ModelDBPhonebookInterfaceImplementation struct {
	db *gorm.DB
}

func CreateNewPhonebookModelInterface(db *gorm.DB) *ModelDBPhonebookInterfaceImplementation {

    return &ModelDBPhonebookInterfaceImplementation{db}

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBCreate(phonebook *PhonebookModel) *MessagePassing.MessageToPass {

    err := model_implementation.db.Create(phonebook).Error 
    if err != nil {

        if strings.Contains(err.Error(), "Unique constraint failed") {

            return &MessagePassing.MessageToPass{
                Message: "Record with given phonenumber already exist",
                Code: MessagePassing.ERROR,
            }

        }

        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }

    } 

    return &MessagePassing.MessageToPass{
        Message: "Record saved to database",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBUpdate(phonebook *PhonebookModel) *MessagePassing.MessageToPass {

    err := model_implementation.db.Save(phonebook).Error 
    if err != nil {

        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }

    }

    return &MessagePassing.MessageToPass{
        Message: "Record updated successfully",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBDelete(phonebook *PhonebookModel) *MessagePassing.MessageToPass {

    if err := model_implementation.db.Delete(phonebook).Error; err != nil {
        
        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }
        
    }

    return &MessagePassing.MessageToPass{
        Message: "Record deleted successfully",
        Code: MessagePassing.INFO,

    }

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBFindByID(id uint) (interface{}, *MessagePassing.MessageToPass) {

    var result PhonebookModel
    
    if err := model_implementation.db.First(&result, id).Error; err != nil {

        if err.Error() == "record not found" {

            return nil, &MessagePassing.MessageToPass{
                Message: "Record not found",
                Code: MessagePassing.WARNING,
            }

        }

        return nil, &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        } 

    }

    return &result, &MessagePassing.MessageToPass{
        Message: "Record found",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBFindAll() ([] interface{}, *MessagePassing.MessageToPass) {

    var results []PhonebookModel
	err := model_implementation.db.Find(&results).Error
    if err != nil {

	    return nil, &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        } 

    }

    var interfaces []interface{}
    for _, v := range results {

        interfaces = append(interfaces, v)

    }

    return interfaces, &MessagePassing.MessageToPass{
        Message: "Records found",
        Code: MessagePassing.INFO,
    }	

}

func PhonebookModelDBMigrationHandler() bool {

    db := Database.DatabaseCreateConnection()
    defer Database.DatabaseCloseConnection(db)
    db.AutoMigrate(&PhonebookModel{})

    return true
    
}




