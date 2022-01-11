package customerservice

import (
	"sirclo/designpattern/facade/hrd"
	"sirclo/designpattern/facade/marketing"
)

type Ticket struct {
	User    string
	Type    string
	Message string
}

type Response struct {
	Status string
}

//facade
func HandleTicket(t Ticket) Response {
	switch t.Type {
	case "marketing":
		response := marketing.ProcessTicket(t.Message)
		return Response{Status: response}
	case "hrd":
		response := hrd.ProcessTicket(t.Message)
		return Response{Status: response}

	}
	return Response{Status: "customer service is busy"}
}
