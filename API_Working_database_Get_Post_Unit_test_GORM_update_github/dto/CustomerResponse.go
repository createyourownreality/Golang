package dto

import (
	"time"
)

type CustomerResponse struct {
	Id          uint      `gorm:"primaryKey" json:"id" `
	Name        string    `json:"full_name"`
	City        string    `json:"city"`
	Zipcode     string    `json:"zipcode"`
	DateOfBirth time.Time ` gorm:"column:date_of_birth" json:"date_of_birth"`
	Status      string    `json:"status"`
}
