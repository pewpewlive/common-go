package blitz

type BlitzPlayer struct {
	Name               string  `json:"name"`
	AccountID          string  `json:"account_id"`
	DisconnectionRound int     `json:"disconnection_round,omitempty"`
	Scores             []int64 `json:"scores"`
	// The points at the end of the last round
	TotalPoints int64 `json:"points"`
	// The points before the start of the last round.
	// Is used to break ties.
	PreviousTotalPoints int64   `json:"previous_points"`
	ReplayHashes        []int64 `json:"replay_hashes"`
}

type BlitzRound struct {
	RoundName string `json:"round_name"`
}

type BlitzResults struct {
	Timestamp int64         `json:"timestamp"`
	Rounds    []BlitzRound  `json:"rounds"`
	Players   []BlitzPlayer `json:"players"`
}
