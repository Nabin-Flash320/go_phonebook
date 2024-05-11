
package ModelDB

import (
	"fmt"
)

type migrationHandlerFunction func() bool

type ModelMigrationManager struct {
	name string
	migration_handler migrationHandlerFunction
}

var migration_manager = [...] ModelMigrationManager {
	{
		"phonebook",
		UserModelDBMigrationHandler,
	},
}

func modelDBMakeMigrationsAll() {

	for _, value := range migration_manager {

		fmt.Printf("\033[33m Migrating %s\n\033[0m", value.name)
		if value.migration_handler() {

			fmt.Println("\033[32m Successfully migrated\033[0m")

		}
		
	}

}

func modelDBMakeMigrationsSelective(model_to_migrate string) {

	var migrated bool = false

	for _, value := range migration_manager {

		if value.name == model_to_migrate {

			fmt.Printf("\033[33m Migrating %s\n\033[0m", value.name)

			if value.migration_handler() {

				fmt.Println("\033[32m Successfully migrated\033[0m")
				migrated = true
				break

			}

		}

	}
	
	if !migrated {
		
		error_message := fmt.Sprintf("\033[31m Model name error. Model named `%s` does not exist.\033[0m", model_to_migrate)
		panic(error_message)

	}

}

func ModelDBMakeMigrations(model_to_migrate string) {

	if model_to_migrate == "all" {

		modelDBMakeMigrationsAll()

	} else {

		modelDBMakeMigrationsSelective(model_to_migrate)

	}

}







