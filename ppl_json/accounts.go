package ppl_json

type AccountInfo struct {
	AccountID string
	Username  string
}

type PaginatedAllAccountInfo struct {
	NextCursor string
	Entries    []AccountInfo
}
