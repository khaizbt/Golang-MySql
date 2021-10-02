package entity

import "time"

type (
	Comment struct {
		Id        int32
		UserId    int32
		PostId    int32
		Comment   string
		CreatedAt time.Time
	}
)
