package enrolmentkey

// The request for an Enrolment Key Action
type EnrolmentKeyBulkAction struct {
	KeyIds []EnrolmentKeyId
}

// The result of a bulk enrolment key operation.
type EnrolmentKeyBulkActionResult struct {
	KeysModified int
}
