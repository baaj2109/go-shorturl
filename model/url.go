package model

import "time"

type UrlCode struct {
	Id        int `gorm:"primary_key"`
	MD5       string
	Code      string
	Url       string
	Click     int
	UserId    int
	CreatedAt time.Time
}

func (u *UrlCode) TableName() string {
	return "url_code"
}
