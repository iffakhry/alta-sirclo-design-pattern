package main

import (
	"fmt"
	"sirclo/designpattern/singleton/repository"
	"sirclo/designpattern/singleton/singleton"
)

func main() {

	db := singleton.GetDBInstance()
	fmt.Println("dari main", *singleton.GetDBInstance())
	result := db.GetUser("alta")
	db.Status = "main"
	fmt.Println(db.Status, db.Type)
	fmt.Println(result)
	singleton.Umum = "hello"
	// singleton.instance = &db{}

	repository.HelloUser()
}
