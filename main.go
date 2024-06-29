
package main



import (
	"flag"

	"github.com/Nabin-Flash320/go_phonebook/Core/Services"	
)



func main() {
	
	var serve *bool = flag.Bool("serve", false, "Run the server")
	var migrate *bool = flag.Bool("migrate", false, "Migrate models to the database")
	var model *string = flag.String("model", "all", "Migrate specific model to the database")
	var createsuperuser *bool = flag.Bool("createsuperuser", false, "Create user with role of super user")
	var populatepermissions *bool = flag.Bool("populatepermissions", false, "Populate database with default permissions namely read, create, modify and delete")
	
	flag.Parse()
	if *serve {

		Services.ServerInit()

	}else if *migrate {

		Services.ModelDBMakeMigrations(*model)

	} else if *createsuperuser {

		Services.CreateSuperUser()

	} else if *populatepermissions {

		Services.PopulateDefault()

	}

}


