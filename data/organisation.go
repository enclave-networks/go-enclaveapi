package data

import "time"

type Organisation struct {
	Id                string
	Created           time.Time
	Name              string
	Plan              int
	Website           string
	Phone             string
	MaxSystems        int
	EnrolledSystems   int64
	UnapprovedSystems int64
}
