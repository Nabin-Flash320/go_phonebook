
package CoreModels


import (
    "fmt"
    "encoding/json"
	"strings"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Nabin-Flash320/go_phonebook/Core/Database"
    "github.com/Nabin-Flash320/go_phonebook/Core/Mailer"
	"github.com/Nabin-Flash320/go_phonebook/Core/MessagePassing"
)

type Role string

const (
	RoleUser Role = "user"
	RoleSuperUser Role = "superuser"
)



type CoreModelUser struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Role Role `json:"role" gorm:"not null"`
    Permissions []CoreModelPermission `json:"permissions" gorm:"many2many:user_permissions"`
}

type CoreModelUserInterface interface {
	Database.ModelDBInterfaces
	ModelDBCreateUser(interface{}) *MessagePassing.MessageToPass 
	ModelDBCreateSuperUser(interface{}) *MessagePassing.MessageToPass
}

type CoreModelUserInterfaceImplementation struct {
	db *gorm.DB
}

func CreateNewCoreModelUserInterface(db *gorm.DB) *CoreModelUserInterfaceImplementation {

	return &CoreModelUserInterfaceImplementation{db}

}

func (model_implementation *CoreModelUserInterfaceImplementation) ModelDBCreateUser(user_model *CoreModelUser) *MessagePassing.MessageToPass {
	
	user_model.Role = RoleUser
    
    default_user_permission := "read"

    db := Database.DatabaseCreateConnection()
    defer Database.DatabaseCloseConnection(db)
    permission_model_implementation := CreateNewPermissionModelInterface(db)
    permissions_retrned, _ := permission_model_implementation.ModelDBFindAll()
    var is_available bool = false
    for _, saved_permission := range permissions_retrned {

        saved_permission_model := (saved_permission).(CoreModelPermission)
        if saved_permission_model.Permission == default_user_permission {

            is_available = true
            user_model.Permissions = append(user_model.Permissions, saved_permission_model)
            break

        }

    }

    if false == is_available {

        return &MessagePassing.MessageToPass{
            Message: "Default read permission is not available, please migrate permissions to the database first.",
            Code: MessagePassing.INFO,
        }

    }

	return model_implementation.ModelDBCreate(user_model)

}

func (model_implementation *CoreModelUserInterfaceImplementation) ModelDBCreateSuperUser(user_model *CoreModelUser) *MessagePassing.MessageToPass {

    if(!CoreMailer.IsValidEmail(user_model.Email)) {

        return &MessagePassing.MessageToPass {
            Message: fmt.Sprintf("Invalid Email `%s`", user_model.Email),
            Code: MessagePassing.ERROR,
        }

    }
	
	user_model.Role = RoleSuperUser

    default_superuser_permissions := []string{"read", "create", "modify", "delete"}

    db := Database.DatabaseCreateConnection()
    defer Database.DatabaseCloseConnection(db)
    permission_model_implementation := CreateNewPermissionModelInterface(db)
    permissions_retrned, _ := permission_model_implementation.ModelDBFindAll()
    for _, permission := range default_superuser_permissions {
        
        var is_available bool = false;
        for _, saved_permission := range permissions_retrned {
            
            saved_permission_model := (saved_permission).(CoreModelPermission)
            if saved_permission_model.Permission == permission {

                is_available = true
                user_model.Permissions = append(user_model.Permissions, saved_permission_model)
                break

            }

        }

        if(false == is_available) {

            return &MessagePassing.MessageToPass {
                Message: fmt.Sprintf("Default permission `%s` not found", permission),
                Code: MessagePassing.ERROR,
            }

        }

    }
    
    user_model_json, err := json.MarshalIndent(user_model, "", "  ")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(user_model_json))
	return model_implementation.ModelDBCreate(user_model)

}

func (model_implementation *CoreModelUserInterfaceImplementation) ModelDBCreate(user_model *CoreModelUser) *MessagePassing.MessageToPass {

    err := model_implementation.db.FirstOrCreate(user_model).Error 
	if err != nil {

        if strings.Contains(err.Error(), "Unique constraint failed") {

            return &MessagePassing.MessageToPass{
                Message: "User with given email already exist",
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
            Message: "User saved to database",
            Code: MessagePassing.INFO,
        }

    }

}

func (model_implementation *CoreModelUserInterfaceImplementation) ModelDBUpdate(user_model *CoreModelUser) *MessagePassing.MessageToPass {

    err := model_implementation.db.Save(user_model).Error 
	if err != nil {

        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }

    } else {

        return &MessagePassing.MessageToPass{
            Message: "User updated successfully",
            Code: MessagePassing.INFO,
        }

    }

}

func (model_implementation *CoreModelUserInterfaceImplementation) ModelDBDelete(user_model *CoreModelUser) *MessagePassing.MessageToPass {

    err := model_implementation.db.Delete(user_model).Error 
	if err != nil {
        
        return &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        }
        
    } else {

        return &MessagePassing.MessageToPass{
            Message: "User deleted successfully",
            Code: MessagePassing.INFO,
        }

    }

}

func (model_implementation *CoreModelUserInterfaceImplementation) ModelDBFindByID(id uint) (interface{}, *MessagePassing.MessageToPass) {

    var result CoreModelUser
    if err := model_implementation.db.First(&result, id).Error; err != nil {

        return nil, &MessagePassing.MessageToPass{
            Message: "Something went wrong",
            Code: MessagePassing.CRITICAL,
        } 

    }

    return &result, &MessagePassing.MessageToPass{
        Message: "User found",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *CoreModelUserInterfaceImplementation) ModelDBFindByValue(db_query string) (interface{}, *MessagePassing.MessageToPass) {

    var result CoreModelUser
    if err := model_implementation.db.First(&result, db_query).Error; err != nil {
        return nil, &MessagePassing.MessageToPass {
            Message: "Something went wong",
            Code: MessagePassing.CRITICAL,
        }

    }

    return &result, &MessagePassing.MessageToPass{
        Message: "Query found",
        Code: MessagePassing.INFO,
    }

}

func (model_implementation *CoreModelUserInterfaceImplementation) ModelDBFindAll() ([] interface{}, *MessagePassing.MessageToPass) {

    var results []CoreModelUser
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
            Message: "Users found",
            Code: MessagePassing.INFO,
        }

	}

}

func CoreModelUserDBMigrationHandler() bool {

    db := Database.DatabaseCreateConnection()
    defer Database.DatabaseCloseConnection(db)
    db.AutoMigrate(&CoreModelUser{})

    return true
    
}

