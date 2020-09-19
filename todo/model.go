package todo

import "time"

//Todo  the model
type Todo struct {
	Title    string    `json:"title"`
	Deadline time.Time `json:"deadline"`
	Checked  bool      `json:"checked"`
}
