package ppl_types

type AccountInfo struct {
	AccountID string `csv:"account_id"`
	Username  string `csv:"username"`
}

type PaginatedAccountInfo struct {
	NextCursor string
	Entries    []AccountInfo
}
