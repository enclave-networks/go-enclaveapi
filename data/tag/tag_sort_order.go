package tag

type TagQuerySortOrder int

const (
	Alphabetical      TagQuerySortOrder = 0
	RecentlyUsed      TagQuerySortOrder = 1
	ReferencedSystems TagQuerySortOrder = 2
)
