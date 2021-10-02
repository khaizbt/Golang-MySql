package entity

import "time"

type (
	User struct {
		Id        int32
		Name      string
		Email     string
		CreatedAt time.Time
	}
)
