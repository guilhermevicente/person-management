package api

import (
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
	api.Echo.PATCH("/persons/:id", api.updatePartOfPerson)
	api.Echo.DELETE("/persons/:id", api.deletePerson)
}
