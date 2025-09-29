package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/guilhermevicente/person-management/schemas"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// Handler
func (api *API) getPersons(c echo.Context) error {
	deleted := c.QueryParam("deleted")

	del, err := strconv.ParseBool(deleted)
	if err != nil {
		log.Warn().Msg("error to parse deleted bool param, set false by default")
		del = false
	}

	persons, err := api.DB.GetPersons(del)
	if err != nil {
		return c.String(http.StatusNoContent, "Don't have person")
	}

	listOfPersons := map[string][]schemas.PersonResponse{"persons": schemas.NewResponse(persons)}
	return c.JSON(http.StatusOK, listOfPersons)
}

func (api *API) createPerson(c echo.Context) error {
	personReq := PersonRequest{}
	if err := c.Bind(&personReq); err != nil {
		return err
	}
	erros := personReq.Validate()
	if len(erros) > 0 {
		log.Error().Msg("error validating person")
		return c.JSON(http.StatusBadRequest, erros)
	}
	person := schemas.Person{
		Name:  personReq.Name,
		Email: personReq.Email,
		TaxId: personReq.TaxId,
	}
	person.Id = uuid.New()
	person.Deleted = false
	if err := api.DB.InsertPerson(person); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create person")
	}
	return c.JSON(http.StatusOK, person.Id)
}

func (api *API) getPerson(c echo.Context) error {
	personId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id for person search")
	}

	person, err := api.DB.GetPerson(personId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNoContent, "Person not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Person")
	}

	return c.JSON(http.StatusOK, person)
}

func (api *API) updatePerson(c echo.Context) error {
	personId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id for person update")
	}
	person := schemas.Person{}
	if err := c.Bind(&person); err != nil {
		return err
	}
	person.Id = personId
	if err := api.DB.UpdatePerson(person); err != nil {
		return c.String(http.StatusInternalServerError, "Error to update person")
	}
	return c.String(http.StatusOK, "Person updated")
}

func (api *API) updatePartOfPerson(c echo.Context) error {
	personId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id for person search")
	}

	patchPerson := schemas.Person{}
	if err := c.Bind(&patchPerson); err != nil {
		return err
	}

	person, err := api.DB.GetPerson(personId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNoContent, "Person not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Person")
	}

	mergedPerson := mergePerson(patchPerson, person)

	if err := api.DB.UpdatePerson(mergedPerson); err != nil {
		return c.String(http.StatusInternalServerError, "Error to update person")
	}
	return c.String(http.StatusOK, "Person updated")
}

func mergePerson(patchPerson, person schemas.Person) schemas.Person {
	if patchPerson.Email != "" {
		person.Email = patchPerson.Email
	}
	if patchPerson.Name != "" {
		person.Name = patchPerson.Name
	}
	if patchPerson.TaxId != "" {
		person.TaxId = patchPerson.TaxId
	}
	return person
}

func (api *API) deletePerson(c echo.Context) error {
	personId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id for person search")
	}

	person, err := api.DB.GetPerson(personId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNoContent, "Person not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Person")
	}

	person.Deleted = true
	if err := api.DB.UpdatePerson(person); err != nil {
		return c.String(http.StatusInternalServerError, "Error to delete person")
	}
	return c.String(http.StatusOK, "Person deleted")
}
