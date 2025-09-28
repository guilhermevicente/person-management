package db

import (
	"github.com/google/uuid"
	"github.com/guilhermevicente/person-management/schemas"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersonHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	dsn := "host=localhost user=person-management password=person-management dbname=person-management port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize DB connection")
	}
	db.AutoMigrate(&schemas.Person{})
	return db
}

func NewPersonHandler(db *gorm.DB) *PersonHandler {
	return &PersonHandler{DB: db}
}

func (p *PersonHandler) InsertPerson(person schemas.Person) error {
	if result := p.DB.Create(&person); result.Error != nil {
		if len(person.TaxId) > 5 {
			log.Fatal().Err(result.Error).Msgf("Failed to create person: %s", person.TaxId[:len(person.TaxId)-6])
		} else {
			log.Fatal().Err(result.Error).Msgf("Failed to create person")
		}
		return result.Error
	}

	log.Info().Msg("Person created")
	return nil
}

func (p *PersonHandler) GetPersons() ([]schemas.Person, error) {
	persons := []schemas.Person{}
	err := p.DB.Where("deleted = ?", false).Find(&persons).Where("deleted = ?", false).Error
	return persons, err
}

func (p *PersonHandler) GetPerson(uuid uuid.UUID) (schemas.Person, error) {
	var person schemas.Person
	err := p.DB.Where("deleted = ?", false).First(&person, uuid)
	return person, err.Error
}

func (p *PersonHandler) UpdatePerson(person schemas.Person) error {
	if result := p.DB.Save(&person); result.Error != nil {
		if len(person.TaxId) > 5 {
			log.Fatal().Err(result.Error).Msgf("Failed to update person: %s", person.TaxId[:len(person.TaxId)-6])
		} else {
			log.Fatal().Err(result.Error).Msgf("Failed to update person")
		}
		return result.Error
	}

	log.Info().Msg("Person updated")
	return nil
}
