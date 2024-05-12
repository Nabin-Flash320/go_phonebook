
package ModelDB


import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)


type PhonebookModel struct {
    gorm.Model
    Name  string
    Address string
    PhoneNumber uint64 `gorm:"unique"` 
    // database constraint enforcement is used to describe the additional information about model column
}

type ModelDBPhonebookInterfaceImplementation struct {
	db *gorm.DB
}

func CreateNewPhonebookModelInterface(db *gorm.DB) *ModelDBPhonebookInterfaceImplementation {

    return &ModelDBPhonebookInterfaceImplementation{db}

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBCreate(phonebook *PhonebookModel) error {

    return model_implementation.db.Create(phonebook).Error 

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBUpdate(phonebook *PhonebookModel) error {

    return model_implementation.db.Save(phonebook).Error 

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBDelete(phonebook *PhonebookModel) error {

    return model_implementation.db.Delete(phonebook).Error 

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBFindByID(id uint) (interface{}, error) {

    var result PhonebookModel
    if err := model_implementation.db.First(&result, id).Error; err != nil {

        return nil, err 

    }

    return &result, nil

}

func (model_implementation *ModelDBPhonebookInterfaceImplementation) ModelDBFindAll() ([] interface{}, error) {

    var results []PhonebookModel
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
    db.AutoMigrate(&PhonebookModel{})

    return true
    
}




