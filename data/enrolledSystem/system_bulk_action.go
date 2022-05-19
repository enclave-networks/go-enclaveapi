package enrolledSystem

// The result of a bulk system revocation
type EnrolledSystemBulkRevokedResult struct {
	SystemsRevoked int
}

// A System bulk action struct
type EnrolledSystemBulkAction struct {
	SystemIds []SystemId
}
