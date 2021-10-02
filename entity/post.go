package entity

import "time"

type (
	Post struct {
		Id        int32
		Title     string
		Body      string
		CreatedBy int32
		CreatedAt time.Time
		UserName  string
	}
)
