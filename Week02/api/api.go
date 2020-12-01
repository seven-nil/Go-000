package api

import (
	"log"
	"week2/model"
	"week2/service"
)

// GetEntity .
func GetEntity() (*model.Entity, error) {
	testID := 1
	entity, err := service.GetEntity(testID)
	if err != nil {
		log.Fatal(err)
	}
	return entity, nil
}
