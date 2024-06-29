
package PhonebookURIHandlers

import (
	"strconv"
	"github.com/gin-gonic/gin"

	"github.com/Nabin-Flash320/go_phonebook/Core/Database"
	"github.com/Nabin-Flash320/go_phonebook/Core/MessagePassing"

	"github.com/Nabin-Flash320/go_phonebook/App/Phonebook/PhonebookModels"
)


func PhonebookHomeUriGetMethodHandler(c *gin.Context) {

	db := Database.DatabaseCreateConnection()
	defer Database.DatabaseCloseConnection(db)

	phonebook_implementation := Models.CreateNewPhonebookModelInterface(db)
	records_available, message := phonebook_implementation.ModelDBFindAll()

	MessagePassing.MessagePassingPanicOnCritical(message)

	c.IndentedJSON(200, records_available)

}

func PhonebookHomeUriGetByIDMethodHandler(c *gin.Context) {

	id_str := c.Param("id")
	if id_str == "" {

		c.JSON(
			400,
			gin.H{
				"status": "ID param error",
			},
		)

	} else {

		id, err := strconv.ParseUint(id_str, 10, 64)
		if err != nil {

			c.JSON(
				400, 
				gin.H{
					"status": "ID param error",
				},
			)

		} else {

			db := Database.DatabaseCreateConnection()
			defer Database.DatabaseCloseConnection(db);

			phonebook_implementation := Models.CreateNewPhonebookModelInterface(db)
			records_available, message := phonebook_implementation.ModelDBFindByID(uint(id))

			if MessagePassing.MessagePassingDoContain(message, MessagePassing.WARNING) {

				c.JSON(
					200,
					gin.H{
						"status": message.Message,
					},
				)

			} else {

				MessagePassing.MessagePassingPanicOnCritical(message)
				c.IndentedJSON(200, records_available)

			}

		}

	}

}

func PhonebookHomeUriPostMethodHandler(c *gin.Context) {

	var phonebook Models.PhonebookModel

	if err := c.BindJSON(&phonebook); err != nil {

		c.JSON(
			400,
			gin.H{
				"status": "JSON data error",
			},
		)

	} else {

		db := Database.DatabaseCreateConnection()
		defer Database.DatabaseCloseConnection(db)

		phonebook_implementation := Models.CreateNewPhonebookModelInterface(db)
		message := phonebook_implementation.ModelDBCreate(&phonebook)
		MessagePassing.MessagePassingPanicOnCritical(message)
		c.JSON(
			400,
			gin.H{
				"status": message.Message,
			},
		)

	}

}


func PhonebookHomeUriPostDeleteRecordMethodHandler(c *gin.Context) {

	id_str := c.Param("id")
	if id_str == "" {

		c.JSON(
			400,
			gin.H{
				"status": "ID param error",
			},
		)

	} else {

		id, _ := strconv.ParseUint(id_str, 10, 64)
		db := Database.DatabaseCreateConnection()
		defer Database.DatabaseCloseConnection(db)
		phonebook_implementation := Models.CreateNewPhonebookModelInterface(db)
		result, message := phonebook_implementation.ModelDBFindByID(uint(id))

		if MessagePassing.MessagePassingDoContain(message, MessagePassing.INFO) {

			/* result is interface pointer so, conversion have to be done to phonebook model pointer */
			phonebook := result.(*Models.PhonebookModel)
			message := phonebook_implementation.ModelDBDelete(phonebook)
			if MessagePassing.MessagePassingDoContain(message, MessagePassing.INFO) {

				c.JSON(
					200,
					gin.H{
						"status": "Record deleted successfully",
					},
				)
	
			} else {

				MessagePassing.MessagePassingPanicOnCritical(message)
				c.JSON(
					200,
					gin.H{
						"status": "Something happened while deleting data",
					},
				)

			}

		} else {

			c.JSON(
				200,
				gin.H{
					"status": message.Message,
				},
			)

		}

	}

}
