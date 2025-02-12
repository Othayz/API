package db

import (
	"github.com/labstack/gommon/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct{
	DB *gorm.DB
}
type Student struct {
	gorm.Model
	Name string     `json:"name"`
	CPF int	        `json:"cpf"`
	Email string	`json:"email"`
	Age int	        `json:"age"`
	Active bool	    `json:"registration"`
}
func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func Init() *gorm.DB {
	db, err:= gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize SQLite: %s", err.Error())
	}
	db.AutoMigrate(&Student{})

	return db
}

func (s *StudentHandler) AddStudent(Student Student) error{
	if result := s.DB.Create(&Student); result.Error != nil {
		log.Error("Failed to add student")
		return result.Error
	}
	log.Info("Student added")
	return nil
}

func (s *StudentHandler) GetStudentByID()([]Student, error){
	students := []Student{}
	err := s.DB.Find(&students).Error
	return students, err
}

func (s *StudentHandler) GetStudent(id int) (Student, error) {
	var student Student
	err := s.DB.First(&student, id)
	return student, err.Error
  }

func (s *StudentHandler) UpdateStudent(UpdateStudent Student)(error) {
	return s.DB.Save(&UpdateStudent).Error
}