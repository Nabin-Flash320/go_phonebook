
package Services


import (
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/Nabin-Flash320/go_phonebook/Core/Database"
	"github.com/Nabin-Flash320/go_phonebook/Core/CoreModels"
    "github.com/Nabin-Flash320/go_phonebook/Core/MessagePassing"
)

const (
	PermissionRead string = "read"
	PermissionCreate string = "create"
	PermissionModify string = "modify"
	PermissionDelete string = "delete"
)

var default_permissions = [4] string{
	PermissionRead,
	PermissionCreate,
	PermissionModify,
	PermissionDelete,
}

func checkDefaultPermissions(db *gorm.DB) []string {

	var default_permissions_to_add []string

	permission_implementation := CoreModels.CreateNewPermissionModelInterface(db)
	all_permissions, _ := permission_implementation.ModelDBFindAll()
	for _, permission := range default_permissions {

		var save_to_db bool = true;
		for _, saved_permission := range all_permissions {

			saved_permissions_model := saved_permission.(CoreModels.CoreModelPermission)
			if saved_permissions_model.Permission == permission {

				save_to_db = false
				fmt.Printf("Permission %s already added.\n", permission)

			}
	
		}

		if true == save_to_db {

			default_permissions_to_add = append(default_permissions_to_add, permission)

		}

	}

	return default_permissions_to_add

}

func populateDefaultPermissions() {

	fmt.Println("Populating the default permissions.")
	db := Database.DatabaseCreateConnection()
	defer Database.DatabaseCloseConnection(db)
	permissions_to_add := checkDefaultPermissions(db)
	if permissions_to_add != nil {

		permission_implementation := CoreModels.CreateNewPermissionModelInterface(db)
		for _, permission_to_add := range permissions_to_add {

			permission_model := &CoreModels.CoreModelPermission{
				Permission: permission_to_add,
			}
			message := permission_implementation.ModelDBCreate(permission_model)
			MessagePassing.MessagePassingPanicOnCritical(message)
			if message.Message == "Record with given permission already exist" {

				fmt.Println("This permission %s record already exist in database\n", permission_to_add)

			} else if message.Message == "Permission saved to database" {

				fmt.Printf("Default permission %s saved\n", permission_to_add)

			}

		}

	}

	fmt.Println("Completed")

}

func PopulateDefault() {

	populateDefaultPermissions()

}

