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

func (Person) TableName() string {
	return "person_management.person"
}
