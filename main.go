
package main



import (
	"flag"
	"fmt"
	"bufio"
	"os"
	"github.com/gin-gonic/gin"

	"github.com/Nabin-Flash320/go_phonebook/Core/Middlewares"
	"github.com/Nabin-Flash320/go_phonebook/Core/Database"
	"github.com/Nabin-Flash320/go_phonebook/Core/Settings"
	"github.com/Nabin-Flash320/go_phonebook/Core/Services"

	"github.com/Nabin-Flash320/go_phonebook/App/Phonebook/PhonebookModels"
)

func serverInit() {

	host_addr := Settings.HostSettingsGetValidHostAddress()
	
	router := gin.Default()

	router.Use(Middlewares.LoggerwareMiddleware())

	Services.ServerRegisterUIR(router)
	
	fmt.Printf("\033[32m Server started at: http://%s:8000 \n\033[0m", host_addr)

	router.Run(host_addr + ":8000")

}

func read_super_user_cred() {

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

			user_model := &Models.UserModel{
				Name: super_user_name,
				Password: super_user_password,
				Email: super_user_email,
			}

			db := Database.DatabaseCreateConnection()
			defer Database.DatabaseCloseConnection(db)
			user_interface_implementation := Models.CreateNewUserModelInterface(db)
			err := user_interface_implementation.ModelDBCreateSuperUser(user_model)
			if err != nil {

				panic("Failed to create super user")

			} else {

				fmt.Println("\033[32m Super user created successfully.\033[0m")

			}

		}

	}

}


func main() {

	var serve *bool = flag.Bool("serve", false, "Run the server")
	var migrate *bool = flag.Bool("migrate", false, "Migrate models to the database")
	var model *string = flag.String("model", "all", "Migrate specific model to the database")
	var createsuperuser *bool = flag.Bool("createsuperuser", true, "Create user with role of super user")

	flag.Parse()

	if *serve {

		serverInit()

	}else if *migrate {

		Services.ModelDBMakeMigrations(*model)

	} else if *createsuperuser {

		read_super_user_cred()

	}


}


