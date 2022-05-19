package policy

// The sort order when making a Policy Request.
type PolicySortOrder int

const (
	Description     PolicySortOrder = 0
	RecentlyCreated PolicySortOrder = 1
)
