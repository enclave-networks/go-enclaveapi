package tag

type TagQuerySortOrder string

const (
	Alphabetical      TagQuerySortOrder = "Alphabetical"
	RecentlyUsed      TagQuerySortOrder = "RecentlyUsed"
	ReferencedSystems TagQuerySortOrder = "ReferencedSystems"
)
