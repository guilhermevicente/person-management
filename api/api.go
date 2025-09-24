package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/guilhermevicente/person-management/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{
		Echo: e,
		DB:   db,
	}
}

func (api *API) Start() error {
	// Start server
	return api.Echo.Start(":8080")
}

func (api *API) ConfigRoutes() {
	// Routes
	api.Echo.GET("/persons", getPersons)
	api.Echo.POST("/persons", createPerson)
	api.Echo.GET("/persons/:id", getPerson)
	api.Echo.PUT("/persons/:id", updatePerson)
	api.Echo.DELETE("/persons/:id", deletePerson)
}

// Handler
func getPersons(c echo.Context) error {
	persons, err := db.GetPersons()
	if err != nil {
		return c.String(http.StatusNoContent, "Don't have person")
	}
	return c.JSON(http.StatusOK, persons)
}

func createPerson(c echo.Context) error {
	person := db.Person{}
	if err := c.Bind(&person); err != nil {
		return err
	}
	person.Id = uuid.New()
	person.Deleted = false
	if err := db.InsertPerson(person); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create person")
	}
	return c.String(http.StatusOK, "Person created")
}

func getPerson(c echo.Context) error {
	personId := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Get person of id %s", personId))
}

func updatePerson(c echo.Context) error {
	personId := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Update person of id %s", personId))
}

func deletePerson(c echo.Context) error {
	personId := c.Param("id")
	return c.String(http.StatusOK, fmt.Sprintf("Delete person of id %s", personId))
}
