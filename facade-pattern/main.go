package main

import (
	"fmt"
	"sirclo/designpattern/facade/customerservice"
)

func main() {
	response := customerservice.HandleTicket(customerservice.Ticket{
		User:    "alta",
		Type:    "marketing",
		Message: "pesan untuk marketing",
	})
	fmt.Println(response.Status)
	response1 := customerservice.HandleTicket(customerservice.Ticket{
		User:    "alta",
		Type:    "hrd",
		Message: "pesan untuk hrd",
	})
	fmt.Println(response1.Status)
}
