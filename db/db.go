package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	Name    string
	TaxId   string
	Email   string
	Deleted bool
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

func InsertPerson() {
	db := Init()
	person := Person{
		Name:    "Gabriel",
		TaxId:   "87505066099",
		Email:   "gabriel@teste.com",
		Deleted: false,
	}
	if result := db.Create(&person); result.Error != nil {
		fmt.Println("Erro on person creation. Error:", result.Error)
	} else {
		fmt.Println("Person created. Result: ", result.RowsAffected)
	}
}
