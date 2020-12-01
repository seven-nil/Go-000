package dao

import (
	"week2/model"

	"github.com/pkg/errors"
)

// GetEntity .
func GetEntity(query int) (*model.Entity, error) {
	result := &model.Entity{}
	result, err := queryEntityFromDBByID(query)
	if err != nil {
		return nil, errors.Wrap(err, "record not found")
	}
	return result, nil
}

// queryEntityFromDBByID 模拟从数据库获取数据
func queryEntityFromDBByID(id int) (*model.Entity, error) {
	// 省略逻辑
	return &model.Entity{}, nil
}
