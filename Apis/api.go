package apis

import (
	"github.com/Othayz/API/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type API struct {
	Echo *echo.Echo
	DB *db.StudentHandler
}

func NewServer()*API {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  
	dataBase:= db.Init()
	storageDB := db.NewStudentHandler(dataBase)

	return &API{
		Echo: e,
		DB: storageDB,
  }
}
func (api *API) Start() error{
	return api.Echo.Start(":8080")
}

  func (api *API) ConfigRoutes(){
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudents)
	api.Echo.GET("/students/:id", api.getStudentsByID)
	api.Echo.PUT("/students/:id", api.updateStudentsByID)
	api.Echo.DELETE("/students/:id", api.deleteStudentsByID)
}

