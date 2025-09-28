package api

import (
	"fmt"

	"github.com/google/uuid"
)

type PersonRequest struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	TaxId   string    `json:"tax_id"`
	Email   string    `json:"email"`
	Deleted bool      `json:"deleted"`
}

func errParamRequired(param, typ string) string {
	return fmt.Sprintf("param %s of type %s is required", param, typ)
}

func (s *PersonRequest) Validate() []string {
	var erros []string
	if s.Name == "" {
		erros = append(erros, errParamRequired("Name", "string"))
	}
	if s.TaxId == "" {
		erros = append(erros, errParamRequired("TaxId", "string"))
	}
	if s.Email == "" {
		erros = append(erros, errParamRequired("Email", "string"))
	}
	return erros
}
