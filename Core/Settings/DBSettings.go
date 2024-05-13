
package Settings


const database_name string = "phonebook"
const database_user string = "root"
const database_pass string = "Nabin&dangi320"
const database_host string = "localhost"
const database_port uint16 = 3306


func DBSettingsGetName() string {
	return database_name
}	

func DBSettingsGetUser() string {
	return database_user
}

func DBSettingsGetPass() string {
	return database_pass
}

func DBSettingsGetHost() string {
	return database_host
}

func DBSettingsGetPort() uint16 {
	return database_port
}


