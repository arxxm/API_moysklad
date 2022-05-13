package main

import (
	"github.com/arxxm/API_moysklad.git/commands"
)

const (
	URLAuth = "https://online.moysklad.ru/api/remap/1.2/security/token"
)

func main() {

	accessToken, err := commands.Authorization(URLAuth, "POST")
	if err != nil {
		panic(err)
	}

	//Get list employees
	// err = commands.GetListEmplyees(accessToken.Access_token)
	// if err != nil {
	// 	panic(err)
	// }

	//Create new employee
	// err = commands.CreateNewEmployee(accessToken.Access_token)
	// if err != nil {
	// 	panic(err)
	// }

	//Change employee
	// err = commands.ChangeEmployee(accessToken.Access_token)
	// if err != nil {
	// 	panic(err)
	// }

	//Delete employee
	err = commands.DeleteEmployee(accessToken.Access_token)
	if err != nil {
		panic(err)
	}

}
