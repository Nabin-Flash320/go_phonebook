
package ModelDB


import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Role string

const (
	RoleUser Role = "user"
	RoleSuperUser Role = "superuser"
)


type UserModel struct {
	gorm.Model
	Name string 
	Email string 
	Password string 
	Role Role
}

type ModelDBUserInterface interface {
	ModelDBInterfaces
	ModelDBCreateUser(interface{}) error 
	ModelDBCreateSuperUser(interface{}) error
}

type ModelDBUserInterfaceImplementation struct {
	db *gorm.DB
}

func CreateNewUserModelInterface(db *gorm.DB) *ModelDBUserInterfaceImplementation {

	return &ModelDBUserInterfaceImplementation{db}

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBCreateUser(user_model *UserModel) error {
	
	user_model.Role = RoleUser
	return model_implementation.ModelDBCreate(user_model)

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBCreateSuperUser(user_model *UserModel) error {
	
	user_model.Role = RoleSuperUser
	return model_implementation.ModelDBCreate(user_model)

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBCreate(user_model *UserModel) error {

    return model_implementation.db.Create(user_model).Error 

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBUpdate(user_model *UserModel) error {

    return model_implementation.db.Save(user_model).Error 

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBDelete(user_model *UserModel) error {

    return model_implementation.db.Delete(user_model).Error 

}

func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBFindByID(id uint) (interface{}, error) {

    var result UserModel
    if err := model_implementation.db.First(&result, id).Error; err != nil {

        return nil, err 

    }

    return &result, nil

}


func (model_implementation *ModelDBUserInterfaceImplementation) ModelDBFindAll() ([] interface{}, error) {

    var results []UserModel
	err := model_implementation.db.Find(&results).Error
	var interfaces []interface{}
	for _, v := range results {

		interfaces = append(interfaces, v)

	}

	return interfaces, err

}

func UserModelDBMigrationHandler() bool {

    db := UserModelDBCreateConnection()
    defer UserModelDBCloseConnection(db)
    db.AutoMigrate(&UserModel{})

    return true
    
}

