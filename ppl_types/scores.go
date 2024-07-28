package ppl_types

type HofEntry struct {
	// Player account IDs separated by a | (bar, pipe). Example: "qEY3iKDcv8JQvNKhNdCXb|KnogIETlchLFxgxtLD72o"
	PlayerAccountIDs string `csv:"account_ids"`
	LevelUUID        string `csv:"level_uuid"`
	LevelVersion     int32  `csv:"level_version"`
	// Either a score, or a number of frames
	Value int64 `csv:"value"`
	// 0 = score. 1 = speed run.
	ValueType int32 `csv:"value_type"`
	// UNIX Timestamp
	Date    int64  `csv:"date"`
	Country string `csv:"country"`
}

type PaginatedScores struct {
	NextCursor string
	Entries    []HofEntry
}
