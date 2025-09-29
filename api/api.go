package api

import (
	"github.com/guilhermevicente/person-management/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/guilhermevicente/person-management/docs"
)

type API struct {
	Echo *echo.Echo
	DB   *db.PersonHandler
}

// @title Person Management API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost
// @BasePath /persons
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
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
