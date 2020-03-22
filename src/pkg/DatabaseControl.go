package models

import "app-test/src/pkg/datasource"

// Create
func Create(value interface{}) error{
	return datasource.DB.Create(value).Error
}
