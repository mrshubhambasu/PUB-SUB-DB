package datastore

import (
	"assignment/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

// Push function takes the data struct of JSON and insert it into database
func Push(dataModel model.Data, db *gorm.DB) error {
	// load data to struct

	//Inserting Data to Hotels Table
	hotel := dataModel.Offers[0].Hotel // need to iterate for multiple objects
	result := db.Create(hotel)         // pass pointer of data to Create
	if result.Error != nil {
		return result.Error
	}

	//Inserting Data to RatePlans Table
	ratePlan := dataModel.Offers[0].RatePlan // need to iterate for multiple objects
	result = db.Create(ratePlan)             // pass pointer of data to Create
	if result.Error != nil {
		return result.Error
	}

	//Inserting Data to Rooms Table
	room := dataModel.Offers[0].Room // need to iterate for multiple objects
	result = db.Create(room)         // pass pointer of data to Create
	if result.Error != nil {
		return result.Error
	}
	return nil
}
