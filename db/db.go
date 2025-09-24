package db

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersonHandler struct {
	DB *gorm.DB
}

type Person struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	TaxId   string    `json:"tax_id"`
	Email   string    `json:"email"`
	Deleted bool      `json:"deleted"`
}

func (Person) TableName() string {
	return "person_management.person"
}

func Init() *gorm.DB {
	dsn := "host=localhost user=person-management password=person-management dbname=person-management port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If this error happens, kill app
		log.Fatalln(err)
	}
	db.AutoMigrate(&Person{})
	return db
}

func NewPersonHandler(db *gorm.DB) *PersonHandler {
	return &PersonHandler{DB: db}
}

func (p *PersonHandler) InsertPerson(person Person) error {
	if result := p.DB.Create(&person); result.Error != nil {
		fmt.Println("Erro on person creation. Error:", result.Error)
		return result.Error
	}

	fmt.Println("Person created")
	return nil
}

func (p *PersonHandler) GetPersons() ([]Person, error) {
	persons := []Person{}
	err := p.DB.Find(&persons).Error
	return persons, err
}
