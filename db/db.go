package db

import (
	"github.com/Othayz/API/schemas"
	"github.com/labstack/gommon/log"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct{
	DB *gorm.DB
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func Init() *gorm.DB {
	db, err:= gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize SQLite: %s", err.Error())
	}
	db.AutoMigrate(&schemas.Student{})

	return db
}

func (s *StudentHandler) AddStudent(Student schemas.Student) error{
	if result := s.DB.Create(&Student); result.Error != nil {
		log.Error("Failed to add student")
		return result.Error
	}
	log.Info("Student added")
	return nil
}

func (s *StudentHandler) GetStudentByID()([]schemas.Student, error){
	students := []schemas.Student{}
	err := s.DB.Find(&students).Error
	return students, err
}

func (s *StudentHandler) GetStudent(id int) (schemas.Student, error) {
	var student schemas.Student
	err := s.DB.First(&student, id)
	return student, err.Error
  }

func (s *StudentHandler) UpdateStudent(UpdateStudent schemas.Student)(error) {
	return s.DB.Save(&UpdateStudent).Error
}
func (s *StudentHandler) DeleteStudent(Student schemas.Student) error {
	return s.DB.Delete(&Student).Error
}