package domain

import (
	"time"
)

type Task struct {
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
	Authors    []Author
}

type Author struct {
	AuthorID int    `json:"authorid"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
}

type SolvedTasks struct {
	TaskID    int       `json:"taskid"`
	TeamID    int       `json:"teamid"`
	Timestamp time.Time `json:"timestemp"`
}

type TaskSubmissions struct {
	TaskID        int       `json:"taskid"`
	TeamID        int       `json:"teamid"`
	SubmitionerID int       `json:"submitionerid"`
	Submission    string    `json:"submission"`
	IsCorrect     bool      `json:"iscorrect"`
	Timrstemp     time.Time `json:"timestemp"`
}
