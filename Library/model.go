package main

import "time"

type Books struct {
	Id          string `gorm:"type:varchar;primary_key;not null"`
	Member      Member `gorm:"association_foreignkey:book_id;foreignkey:id"`
	AddedAt     time.Time
	IssueAt     time.Time
	ReturnedAt  time.Time
	Name        string `gorm:"type:varchar;not null"`
	Publication string `gorm:"type:varchar;not null"`
}
type Member struct {
	MemberId string `gorm:"type:varchar;not null;primary_key"`
	Name     string `gorm:"type:varchar;not null"`
	Email    string `gorm:"type:varchar;not null"`
}
