package policy

// Struct representing a Bulk action request.
type PolicyBulkAction struct {
	PolicyIds []PolicyId
}

// Struct for the result of a policy bulk delete operation.
type PolicyBulkDeleteResult struct {
	PoliciesDeleted int
}

// Defines the result of a bulk policy update.
type PolicyBulkUpdateResult struct {
	PoliciesUpdated int
}
