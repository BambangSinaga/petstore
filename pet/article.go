package pet

import "time"

type Pet struct {
	ID        	int64   `json:"id"`
	Name     	string  `json:"name"`
	Age   		int64	`json:"age"`
	Photo     	string  `json:"photo"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
