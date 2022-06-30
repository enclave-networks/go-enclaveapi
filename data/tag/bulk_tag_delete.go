package tag

type BulkTagAction struct {
	Tags []string
}

type BulkTagDeleteResult struct {
	TagsDeleted int
}
