package quests

type QuestKind int

const (
	destroyEnemies QuestKind = iota
	reachScore
	surviveDuration
	playLevelsForDuration
	playMultiplayerForDuration
)

type LevelIDWithName struct {
	ID      string `json:"uuid"`
	Version int    `json:"version"`
	Name    string `json:"name"`
}

type DestroyEnemiesQuest struct {
	Kind   QuestKind         `json:"kind"`
	Goal   int64             `json:"goal"`
	Enemy  string            `json:"enemy"`
	Weapon string            `json:"weapon,omitempty"`
	Levels []LevelIDWithName `json:"levels,omitempty"`
	XP     int64             `json:"xp"`
}

type ReachScoreQuest struct {
	Kind   QuestKind         `json:"kind"`
	Goal   int64             `json:"goal"`
	Levels []LevelIDWithName `json:"levels,omitempty"`
	XP     int64             `json:"xp"`
}

type SurviveDurationQuest struct {
	Kind QuestKind `json:"kind"`
	// In ticks
	Goal   int64             `json:"goal"`
	Levels []LevelIDWithName `json:"levels"`
	XP     int64             `json:"xp"`
}

type PlayLevelsForDurationQuest struct {
	Kind   QuestKind         `json:"kind"`
	Goal   int64             `json:"goal"`
	Levels []LevelIDWithName `json:"levels"`
	XP     int64             `json:"xp"`
}

type PlayMultiplayerForDurationQuest struct {
	Kind   QuestKind         `json:"kind"`
	Goal   int64             `json:"goal"`
	Levels []LevelIDWithName `json:"levels,omitempty"`
	XP     int64             `json:"xp"`
}

type DailyQuests struct {
	Version    int64 `json:"version"`
	Expiration int64 `json:"expiration"`
	QuestsID   int   `json:"quests_id"`
	Quests     []any `json:"quests"`
}
