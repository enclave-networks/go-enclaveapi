package policy

// The sort order when making a Policy Request.
type PolicySortOrder string

const (
	Description     PolicySortOrder = "Description"
	RecentlyCreated PolicySortOrder = "RecentlyCreated"
)
