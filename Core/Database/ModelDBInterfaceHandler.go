

package Database


import (

	"github.com/Nabin-Flash320/go_phonebook/Core/MessagePassing"
)


type ModelDBInterfaces interface {

	ModelDBCreate(interface{}) *MessagePassing.MessageToPass
	ModelDBUpdate(interface{}) *MessagePassing.MessageToPass
	ModelDBDelete(interface{}) *MessagePassing.MessageToPass
	ModelDBFindByID(uint) (interface{}, *MessagePassing.MessageToPass)
	ModelDBFindAll() ([]interface{}, *MessagePassing.MessageToPass)

}




