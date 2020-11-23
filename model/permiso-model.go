package model

import (
	"time"
)

// Permiso model of table permisos
type Permiso struct {
	ID          string    `gorm:"primary_key;type:varchar(50)"`
	Name        string    `gorm:"type:varchar(100); not null"`
	Description string    `gorm:"type:varchar(200); not null"`
	Status      bool      `gorm:"type:boolean;not null"`
	Owner       string    `gorm:"type:varchar(100);not null;unique"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

//Permisos if array of Permiso
type Permisos []Permiso
