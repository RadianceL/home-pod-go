package models

import "go-gin-template/src/pkg/datasource"

// Create 方法执行
func Create(value interface{}) error {
	return datasource.DB.Create(value).Error
}
