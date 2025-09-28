package schemas

import (
	"github.com/google/uuid"
)

type Person struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	TaxId   string    `json:"tax_id"`
	Email   string    `json:"email"`
	Deleted bool      `json:"deleted"`
}

type PersonResponse struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	TaxId string    `json:"tax_id"`
	Email string    `json:"email"`
}

func NewResponse(persons []Person) []PersonResponse {
	personsResponse := []PersonResponse{}
	for _, person := range persons {
		personResponse := PersonResponse{
			Id:    person.Id,
			Name:  person.Name,
			TaxId: person.TaxId,
			Email: person.Email,
		}
		personsResponse = append(personsResponse, personResponse)
	}
	return personsResponse
}

func (Person) TableName() string {
	return "person_management.person"
}
