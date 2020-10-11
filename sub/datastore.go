package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// This function takes the byte array of JSON and insert it into database
func pushToDatabase(bytes []byte, db *gorm.DB) {
	// load data to struct
	dataModel := Data{}
	dataModel, err := byteToStruct(bytes, dataModel)
	if err != nil {
		log.Error(err)
	}

	//Inserting Data to Hotels Table
	hotel := dataModel.Offers[0].Hotel // need to iterate for multiple objects
	result := db.Create(hotel)         // pass pointer of data to Create

	if result.Error != nil {
		log.Error(result.Error)
	}

	//Inserting Data to RatePlans Table
	ratePlan := dataModel.Offers[0].RatePlan // need to iterate for multiple objects
	result = db.Create(ratePlan)             // pass pointer of data to Create

	if result.Error != nil {
		log.Error(result.Error)
	}

	//Inserting Data to Rooms Table
	room := dataModel.Offers[0].Room // need to iterate for multiple objects
	result = db.Create(room)         // pass pointer of data to Create

	if result.Error != nil {
		log.Error(result.Error)
	}
}

// This function populates the struct with data from byte array
func byteToStruct(bytes []byte, dataModel Data) (Data, error) {
	err := json.Unmarshal(bytes, &dataModel)
	if err != nil {
		return Data{}, err
	}
	return dataModel, nil
}
