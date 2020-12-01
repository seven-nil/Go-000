package service

import (
	"week2/dao"
	"week2/model"
)

// GetEntity .
func GetEntity(query int) (*model.Entity, error) {
	return dao.GetEntity(query)
}
