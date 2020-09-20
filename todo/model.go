package todo

import (
	"time"
)

//Todo  the model
type Todo struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Title     string    `gorm:"type:varchar(40); unique_index; not null" json:"title"`
	Deadline  time.Time `json:"deadline"`
	Checked   bool      `json:"checked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
