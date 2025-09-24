package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	Name    string `json:"name"`
	TaxId   string `json:"tax_id"`
	Email   string `json:"email"`
	Deleted bool   `json:"deleted"`
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

func InsertPerson(person Person) error {
	db := Init()
	if result := db.Create(&person); result.Error != nil {
		fmt.Println("Erro on person creation. Error:", result.Error)
		return result.Error
	}

	fmt.Println("Person created")
	return nil
}

func GetPersons() ([]Person, error) {
	persons := []Person{}
	db := Init()
	err := db.Find(&persons).Error
	return persons, err
}
