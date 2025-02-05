package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Othayz/API/db"
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
  e.GET("/students", getStudents)
  e.POST("/students", createStudents)
  e.GET("/students/:id", getStudentsByID)
  e.PUT("/students/:id", updateStudentsByID)
  e.DELETE("/students/:id", deleteStudentsByID)

  // Start server
  if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
    slog.Error("failed to start server", "error", err)
  }
}

// Handler
func getStudents(c echo.Context) error {
  return c.String(http.StatusOK, "list of all students")
}
func createStudents(c echo.Context) error {
  student := db.Student{}
  if err := c.Bind(&student); err != nil {
    return err
  }
  if err := db.AddStudent(student); err != nil {
    return c.String(http.StatusInternalServerError, "error adding student")
  }

  return c.String(http.StatusOK, "create a new student")
}

func getStudentsByID(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, id)
}
func updateStudentsByID(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, id)
}
func deleteStudentsByID(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, id)
}