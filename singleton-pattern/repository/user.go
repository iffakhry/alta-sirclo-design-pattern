package repository

import (
	"fmt"
	"sirclo/designpattern/singleton/singleton"
)

func HelloUser() {
	db := singleton.GetDBInstance()
	fmt.Println("dari user", *singleton.GetDBInstance())
	db.Status = "user"
	fmt.Println(db.GetUser("sirclo"), db.Status, db.Type)
}
