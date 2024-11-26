package models

import "time"

type Submission struct {
	Code         string    `json:"code"`
	ProblemTitle string    `json:"problemTitle"`
	Status       string    `json:"status"`
	Time         time.Time `json:"time"`
	RunTime      string    `json:"runTime"`
	Language     string    `json:"language"`
	Error        string    `json:"error"`
	Attempts     int       `json:"attempts"`
}
