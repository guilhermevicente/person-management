package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/guilhermevicente/person-management/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/persons", getPersons)
	e.POST("/persons", createPerson)
	e.GET("/persons/:id", getPerson)
	e.PUT("/persons/:id", updatePerson)
	e.DELETE("/persons/:id", deletePerson)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
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
