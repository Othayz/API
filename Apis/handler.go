package apis

import( 
    "net/http"
     "github.com/Othayz/API/db"
	 "github.com/labstack/echo/v4"
)

func (api *API) getStudents(c echo.Context) error {
	Students, err :=api.DB.GetStudentByID()
	if err != nil {
	  return c.String(http.StatusNotFound, "error getting students")
	}
	return c.JSON(http.StatusOK, Students)
  }
  
  func (api *API) createStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
	  return err
	}
	if err := api.DB.AddStudent(student); err != nil {
	  return c.String(http.StatusInternalServerError, "error adding student")
	}
  
	return c.String(http.StatusOK, "create a new student")
  }
  
  func (api *API) getStudentsByID(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
  }
  func (api *API) updateStudentsByID(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
  }
  func (api *API) deleteStudentsByID(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
  }