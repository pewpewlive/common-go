package ppl_json

import "time"

type LevelFullID struct {
	LevelUUID    string
	LevelVersion int32
}

type HofEntry struct {
	PlayerAccountIDs []string
	Level            LevelFullID
	Value            int64
	Date             time.Time
	Country          string
}

type PaginatedAllScores struct {
	NextCursor string
	Entries    []HofEntry
}
