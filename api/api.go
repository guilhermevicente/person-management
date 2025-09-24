package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/guilhermevicente/person-management/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.PersonHandler
}

func NewServer() *API {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	personDB := db.NewPersonHandler(database)

	return &API{
		Echo: e,
		DB:   personDB,
	}
}

func (api *API) Start() error {
	// Start server
	return api.Echo.Start(":8080")
}

func (api *API) ConfigRoutes() {
	// Routes
	api.Echo.GET("/persons", api.getPersons)
	api.Echo.POST("/persons", api.createPerson)
	api.Echo.GET("/persons/:id", api.getPerson)
	api.Echo.PUT("/persons/:id", api.updatePerson)
	api.Echo.DELETE("/persons/:id", api.deletePerson)
}

// Handler
func (api *API) getPersons(c echo.Context) error {
	persons, err := api.DB.GetPersons()
	if err != nil {
		return c.String(http.StatusNoContent, "Don't have person")
	}
	return c.JSON(http.StatusOK, persons)
}

func (api *API) createPerson(c echo.Context) error {
	person := db.Person{}
	if err := c.Bind(&person); err != nil {
		return err
	}
	person.Id = uuid.New()
	person.Deleted = false
	if err := api.DB.InsertPerson(person); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create person")
	}
	return c.String(http.StatusOK, "Person created")
}

func (api *API) getPerson(c echo.Context) error {
	personId := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Get person of id %s", personId))
}

func (api *API) updatePerson(c echo.Context) error {
	personId := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Update person of id %s", personId))
}

func (api *API) deletePerson(c echo.Context) error {
	personId := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Delete person of id %s", personId))
}
