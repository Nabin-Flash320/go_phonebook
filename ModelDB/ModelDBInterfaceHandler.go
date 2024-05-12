

package ModelDB


type ModelDBInterfaces interface {

	ModelDBCreate(interface{}) error
	ModelDBUpdate(interface{}) error
	ModelDBDelete(interface{}) error
	ModelDBFindByID(uint) (interface{}, error)
	ModelDBFindAll() ([]interface{}, error)

}


type ModelDBManagerInterface interface {

	ModelDBManagerCreateUser(interface{}) error 
	ModelDBManagerCreateSuperUser(interface{}) error

}



