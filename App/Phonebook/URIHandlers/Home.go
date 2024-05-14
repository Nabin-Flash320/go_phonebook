
package URIHandlers

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/Nabin-Flash320/go_phonebook/Core/Database"
	"github.com/Nabin-Flash320/go_phonebook/App/Phonebook/Models"
)


func HomeUriGetMethodHandler(c *gin.Context) {

	db := Database.DatabaseCreateConnection()
	defer Database.DatabaseCloseConnection(db)

	phonebook_implementation := Models.CreateNewPhonebookModelInterface(db)
	itfs, err := phonebook_implementation.ModelDBFindAll()

	if err != nil {

		c.JSON(
			400, 
			gin.H{
				"status": "failed",
			},
		)

	} else {

		c.IndentedJSON(200, itfs)

	}

}


func HomeUriPostMethodHandler(c *gin.Context) {

	var phonebook Models.PhonebookModel

	if err := c.BindJSON(&phonebook); err != nil {

		c.JSON(
			400,
			gin.H{
				"status": "error occured",
			},
		)

	} else {

		db := Database.DatabaseCreateConnection()
		defer Database.DatabaseCloseConnection(db)

		phonebook_implementation := Models.CreateNewPhonebookModelInterface(db)
		err := phonebook_implementation.ModelDBCreate(&phonebook)
		if err != nil {

			panic("Error creating model and saving it.")

		} else {

			fmt.Println("Database saved successfully")

		}

		c.JSON(
			200, 
			gin.H{
				"status": "success",
			},
		)
	}

}

