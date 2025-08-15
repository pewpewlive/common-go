package blitz

type BlitzPlayer struct {
	Name               string  `json:"name"`
	AccountID          string  `json:"account_id"`
	DisconnectionRound int     `json:"disconnection_round,omitempty"`
	Scores             []int64 `json:"scores"`
}

type BlitzRound struct {
	RoundName string `json:"round_name"`
}

type BlitzResults struct {
	Timestamp int64         `json:"timestamp"`
	Rounds    []BlitzRound  `json:"rounds"`
	Players   []BlitzPlayer `json:"players"`
}
