package main

import (
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "log/slog"
  "net/http"
  "errors"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/students", getStudents)
  e.GET("/students/:id", getStudentsByID)
  e.GET("/teachers", getTeachers)

  // Start server
  if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
    slog.Error("failed to start server", "error", err)
  }
}

// Handler
func getStudents(c echo.Context) error {
  return c.String(http.StatusOK, "list of all students")
}
func getStudentsByID(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, id)
}
func getTeachers(c echo.Context) error {
  return c.String(http.StatusOK, "list of all teachers")
}