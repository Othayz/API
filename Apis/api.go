package apis

import ("gorm.io/gorm"
	"net/http"

	"github.com/Othayz/API/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


type API struct {
	Echo *echo.Echo
	DB *gorm.DB
}

func NewServer()*API {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
  
	db:= db.Init()

	return &API{
		Echo: e,
		DB: db,
  }
}
func (api *API) Start() error{
	return api.Echo.Start(":8080")
}

  func (api *API) ConfigRoutes(){
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/students", createStudents)
	api.Echo.GET("/students/:id", getStudentsByID)
	api.Echo.PUT("/students/:id", updateStudentsByID)
	api.Echo.DELETE("/students/:id", deleteStudentsByID)
}

func getStudents(c echo.Context) error {
	Students, err := db.GetStudentByID()
	if err != nil {
	  return c.String(http.StatusNotFound, "error getting students")
	}
	return c.JSON(http.StatusOK, Students)
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