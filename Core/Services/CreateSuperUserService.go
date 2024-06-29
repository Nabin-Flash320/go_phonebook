

package Services

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	// "github.com/gin-gonic/gin"

	"github.com/Nabin-Flash320/go_phonebook/Core/Database"
	"github.com/Nabin-Flash320/go_phonebook/Core/CoreModels"
)


func CreateSuperUser() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the name for the super user => ")
	super_user_name, err := reader.ReadString('\n')

	if err != nil {

		panic("Something went wrong while reading super user name.")

	}

	fmt.Print("Enter the email for the super user => ")
	super_user_email, err := reader.ReadString('\n')

	if err != nil {

		panic("Something went wrong while reading super user email.")

	}

	fmt.Print("Enter the password for the super user => ")
	super_user_password, err := reader.ReadString('\n')

	if err != nil {

		panic("Something went wrong while reading super user password.")

	}

	fmt.Print("Enter the password for the super user => ")
	super_user_password_confirm, err := reader.ReadString('\n')

	if err != nil {
		panic("Something went wrong while reading super user password.")
	} else {

		if super_user_password != super_user_password_confirm {

			panic("Password did not match")

		} else {

			fmt.Println("The user name provided is:", super_user_name)
			fmt.Println("The provided password is: ", super_user_password)

			user_model := &CoreModels.CoreModelUser{
				Name: strings.TrimSpace(super_user_name),
				Password: strings.TrimSpace(super_user_password),
				Email: strings.TrimSpace(super_user_email),
			}

			db := Database.DatabaseCreateConnection()
			defer Database.DatabaseCloseConnection(db)
			user_interface_implementation := CoreModels.CreateNewCoreModelUserInterface(db)
			err := user_interface_implementation.ModelDBCreateSuperUser(user_model)
			if err != nil {

				fmt.Println(err.Message)
				panic("Failed to create super user")

			} else {

				fmt.Println("\033[32m Super user created successfully.\033[0m")

			}

		}

	}

}


