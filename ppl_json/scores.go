package ppl_json

import "time"

type LevelFullID struct {
	LevelUUID    string
	LevelVersion int32
}

type HofEntry struct {
	PlayerAccountIDs []string
	Level            LevelFullID
	// Either a score, or a number of frames
	Value int64
	// 0 = score. 1 = speed run.
	Type    int32
	Date    time.Time
	Country string
}

type PaginatedAllScores struct {
	NextCursor string
	Entries    []HofEntry
}
