
package MessagePassing


type ErrorCode uint8

const (
	NONE 		=	0
	ERROR 		=	1
	WARNING 	=	2
	INFO 		=	3
	CRITICAL	= 	4
)


type MessageToPass struct {
	Message string `json:"message"`
	Code ErrorCode `json:"code"`
}


func MessagePassingDoContain(message_to_check *MessageToPass, code_to_check ErrorCode) bool {

	if message_to_check != nil {

		if message_to_check.Code == code_to_check {

			return true
				
		}

	}

	return false 

}

func MessagePassingPanicOnCritical(message_to_check *MessageToPass) {

	if message_to_check != nil {

		if MessagePassingDoContain(message_to_check, CRITICAL) {

			panic("Message level critical encountered.")

		}

	}

}

