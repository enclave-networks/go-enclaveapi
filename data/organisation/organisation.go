package organisation

import "time"

type OrganisationId string

// Organisation properties model.
type Organisation struct {
	Id                OrganisationId
	Created           time.Time
	Name              string
	Plan              int
	Website           string
	Phone             string
	MaxSystems        int
	EnrolledSystems   int64
	UnapprovedSystems int64
}
