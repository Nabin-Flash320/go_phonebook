
package Models


import (
	"strings"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Nabin-Flash320/go_phonebook/Core/Database"
	"github.com/Nabin-Flash320/go_phonebook/Core/MessagePassing"
)

type Role string

const (
	RoleUser Role = "user"
	RoleSuperUser Role = "superuser"
)


type UserModel struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Role Role `json:"role" gorm:"not null"`
}

type ModelDBUserInterface interface {
	Database.ModelDBInterfaces
	ModelDBCreateUser(interface{}) *MessagePassing.MessageToPass 
	ModelDBCreateSuperUser(interface{}) *MessagePassing.MessageToPass
}

type ModelDBUserInterfaceImplementation struct {
	db *gorm.DB
}

func CreateNewUserModelInterface(db *gorm.DB) *ModelDBUserInterfaceImplementation {

	return &ModelDBUserInterfaceImplementation{db}

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBCreateUser(user_model *UserModel) *MessagePassing.MessageToPass {
	
	user_model.Role = RoleUser
	return model_implementation.ModelDBCreate(user_model)

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBCreateSuperUser(user_model *UserModel) *MessagePassing.MessageToPass {
	
	user_model.Role = RoleSuperUser
	return model_implementation.ModelDBCreate(user_model)

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBCreate(user_model *UserModel) *MessagePassing.MessageToPass {

    err := model_implementation.db.Create(user_model).Error 
	if err != nil {

        if strings.Contains(err.Error(), "Unique constraint failed") {

            return &MessagePassing.MessageToPass{
                Message: "Record with given email already exist",
                Code: MessagePassing.ERROR,
            }

        } else {

            return &MessagePassing.MessageToPass{
                Message: "Something went wrong",
                Code: MessagePassing.CRITICAL,
            }

        }

    } else {

        return &MessagePassing.MessageToPass{
            Message: "Record saved to database",
            Code: MessagePassing.INFO,
        }

    }

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBUpdate(user_model *UserModel) *MessagePassing.MessageToPass {

    err := model_implementation.db.Save(user_model).Error 
	if err != nil {

        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }

    } else {

        return &MessagePassing.MessageToPass{
            Message: "Record updated successfully",
            Code: MessagePassing.INFO,
        }

    }

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBDelete(user_model *UserModel) *MessagePassing.MessageToPass {

    err := model_implementation.db.Delete(user_model).Error 
	if err != nil {
        
        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }
        
    } else {

        return &MessagePassing.MessageToPass{
            Message: "Record deleted successfully",
            Code: MessagePassing.INFO,
        }

    }

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBFindByID(id uint) (interface{}, *MessagePassing.MessageToPass) {

    var result UserModel
    if err := model_implementation.db.First(&result, id).Error; err != nil {

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


func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBFindAll() ([] interface{}, *MessagePassing.MessageToPass) {

    var results []UserModel
	err := model_implementation.db.Find(&results).Error
	if err != nil {

		return nil, &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        } 

	} else {

		var interfaces []interface{}
		for _, v := range results {
	
			interfaces = append(interfaces, v)
	
		}
	
		return interfaces, &MessagePassing.MessageToPass{
            Message: "Records found",
            Code: MessagePassing.INFO,
        }

	}

}

func UserModelDBMigrationHandler() bool {

    db := Database.DatabaseCreateConnection()
    defer Database.DatabaseCloseConnection(db)
    db.AutoMigrate(&UserModel{})

    return true
    
}

