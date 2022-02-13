package models

import (
	"github.com/google/uuid"
	"time"
)

type Movement struct {
	UUID   uuid.UUID `json:"uuid" gorm:"primary_key"`
	Name   string    `json:"name"`
	Amount int64     `json:"amount"`
	Date   time.Time `json:"date"`
}
