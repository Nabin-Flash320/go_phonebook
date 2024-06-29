
package CoreModels


import (
    "fmt"
    "strings"
	"github.com/jinzhu/gorm"

	"github.com/Nabin-Flash320/go_phonebook/Core/Database"
    "github.com/Nabin-Flash320/go_phonebook/Core/MessagePassing"
)


type CoreModelPermission struct {
	gorm.Model
	Permission string `gorm:"unique;not null"`
}

type CoreModelPermissionInterface interface {
    Database.ModelDBInterfaces
    ModelDBFindPermissionsFindByNames([]string) (interface{}, *MessagePassing.MessageToPass)
    ModelDBFindPermissionsFindByName(string) (interface{}, *MessagePassing.MessageToPass)
}

type CoreModelPermissionInterfaceImplementation struct {
	db *gorm.DB
}

func CreateNewPermissionModelInterface(db *gorm.DB) *CoreModelPermissionInterfaceImplementation {

	return &CoreModelPermissionInterfaceImplementation{db}

}

func (model_implementation *CoreModelPermissionInterfaceImplementation) ModelDBFindPermissionsFindByNames(permissions []string) (interface{}, *MessagePassing.MessageToPass) {

    if permissions == nil {

        return nil, &MessagePassing.MessageToPass{
            Message: "Empty permissions array passed",
            Code: MessagePassing.ERROR,
        }

    }

    var core_permissions []*CoreModelPermission
    if err := model_implementation.db.Where("Permission IN ?", permissions).Find(&core_permissions).Error; err != nil {

        return nil, &MessagePassing.MessageToPass{
            Message: "Couldn't find the permission",
            Code: MessagePassing.ERROR,
        }

    }

    return core_permissions, &MessagePassing.MessageToPass{
        Message: "All permissiosn retrieved",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *CoreModelPermissionInterfaceImplementation) ModelDBCreate(permission *CoreModelPermission) *MessagePassing.MessageToPass{

	err := model_implementation.db.Create(permission).Error 
    if err != nil {

        if strings.Contains(err.Error(), "Unique constraint failed") || strings.Contains(err.Error(), "Duplicate entry") {

            return &MessagePassing.MessageToPass{
                Message: "Record with given permission already exist",
                Code: MessagePassing.ERROR,
            }

        }

        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }

    } 
    
    fmt.Printf("No error while saving database: %v", err)
    return &MessagePassing.MessageToPass{
        Message: "Permission saved to database",
        Code: MessagePassing.INFO,
    }	

}

func (model_implementation *CoreModelPermissionInterfaceImplementation) ModelDBUpdate(permission *CoreModelPermission) *MessagePassing.MessageToPass {

    err := model_implementation.db.Save(permission).Error 
    if err != nil {

        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }

    }

    return &MessagePassing.MessageToPass{
        Message: "Permission updated successfully",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *CoreModelPermissionInterfaceImplementation) ModelDBDelete(permission *CoreModelPermission) *MessagePassing.MessageToPass {

    if err := model_implementation.db.Delete(permission).Error; err != nil {
        
        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }
        
    }

    return &MessagePassing.MessageToPass{
        Message: "Permission deleted successfully",
        Code: MessagePassing.INFO,

    }

}

func (model_implementation *CoreModelPermissionInterfaceImplementation) ModelDBFindByID(id uint) (interface{}, *MessagePassing.MessageToPass) {

    var result CoreModelPermission
    
    if err := model_implementation.db.First(&result, id).Error; err != nil {

        if err.Error() == "Permission not found" {

            return nil, &MessagePassing.MessageToPass{
                Message: "Permission not found",
                Code: MessagePassing.WARNING,
            }

        }

        return nil, &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        } 

    }

    return &result, &MessagePassing.MessageToPass{
        Message: "Permission found",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *CoreModelPermissionInterfaceImplementation) ModelDBFindByValue(db_query string) (interface{}, *MessagePassing.MessageToPass) {

    var result CoreModelPermission
    if err := model_implementation.db.First(&result, db_query).Error; err != nil {

        if err.Error() == "Permission not found" {

            return nil, &MessagePassing.MessageToPass{
                Message: "Permission query not found",
                Code: MessagePassing.WARNING,
            }

        }

        return nil, &MessagePassing.MessageToPass {
            Message: "Something went wong",
            Code: MessagePassing.CRITICAL,
        }

    }

    return &result, &MessagePassing.MessageToPass{
        Message: "Permission query found",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *CoreModelPermissionInterfaceImplementation) ModelDBFindAll() ([] interface{}, *MessagePassing.MessageToPass) {

    var results []CoreModelPermission
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
        Message: "Permission found",
        Code: MessagePassing.INFO,
    }	

}

func CoreModelPermissionDBMigrationHandler() bool {

    db := Database.DatabaseCreateConnection()
    defer Database.DatabaseCloseConnection(db)
    db.AutoMigrate(&CoreModelPermission{})

    return true
    
}

