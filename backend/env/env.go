package env

import "gorm.io/gorm"

type Env struct {
	db *gorm.DB
}
