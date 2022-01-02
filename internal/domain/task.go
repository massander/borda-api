package domain

import (
	"time"
)

type Tasks struct {
	TaskID     int    `json:"taskid"`
	Title      string `json:"title"`
	Decription string `json:"description"`
	Category   string `json:"category"`
	Complexity string `json:"complexity"`
	Pionts     int    `json:"pints"`
	Hint       string `json:"hint"`
	Flag       string `json:"flag"`
	IsActive   bool   `json:"isactive"`
	IsDisable  bool   `json:"isdisable"`
	Authors    []Authors
}

type Authors struct {
	AuthorID int    `json:"authorid"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
}

type SolvedTasks struct {
	TaskID    int       `json:"taskid"`
	TeamID    int       `json:"teamid"`
	Timestamp time.Time `json:"timestemp"`
}

type TasksSubmissions struct {
	TaskID        int       `json:"taskid"`
	TeamID        int       `json:"teamid"`
	SubmitionerID int       `json:"submitionerid"`
	Submission    string    `json:"submission"`
	IsCorrect     bool      `json:"iscorrect"`
	Timestemp     time.Time `json:"timestemp"`
}
