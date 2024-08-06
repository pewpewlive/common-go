package ppl_types

// Level holds the pair of level UUID and version.
type Level struct {
	LevelUUID    string
	LevelVersion int32
}

// RankThresholds holds the medals data in a manifest.
type RankThresholds struct {
	Bronze *int64 `json:"bronze"`
	Silver *int64 `json:"silver"`
	Gold   *int64 `json:"gold"`
}

// LevelManifestJSON is the JSON for the manifest of a level.
type LevelManifestJSON struct {
	Name                *string         `json:"name"`
	Description         *[]string       `json:"descriptions"`
	Information         *string         `json:"information"`
	EntryPoint          *string         `json:"entry_point"`
	RankThresholds1P    *RankThresholds `json:"rank_thresholds_1p"`
	HasScoreLeaderboard *bool           `json:"has_score_leaderboard"`
}
