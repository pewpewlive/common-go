package quests

import (
	"math/rand"
)

func getRandomElement[T any](arr []T) T {
	index := rand.Intn(len(arr))
	return arr[index]
}

func areLevelsEqual(a, b []LevelIDWithName) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].ID != b[i].ID {
			return false
		}
	}
	return true
}

// A function type used internally to get back a LevelIDWithName constructed by the server
type LevelIDWithNameFunc func(levelUUID string) (*LevelIDWithName, error)

func GenerateRandomDailyQuests(expiration int64, questID int, levelIdFunc LevelIDWithNameFunc) (*DailyQuests, error) {
	scoreQuest, errScore := RandomReachScoreQuest(levelIdFunc)
	destroyEnemiesQuest, errDestroy := RandomDestroyEnemiesQuest(levelIdFunc)
	surviveQuest, errSurvive := RandomSurviveDurationQuest(levelIdFunc)

	if errScore != nil {
		return nil, errScore
	}
	if errDestroy != nil {
		return nil, errDestroy
	}
	if errSurvive != nil {
		return nil, errSurvive
	}

	for areLevelsEqual(surviveQuest.Levels, scoreQuest.Levels) {
		scoreQuest, errScore = RandomReachScoreQuest(levelIdFunc)
		if errScore != nil {
			return nil, errScore
		}
		surviveQuest, errSurvive = RandomSurviveDurationQuest(levelIdFunc)
		if errSurvive != nil {
			return nil, errSurvive
		}
	}

	return &DailyQuests{
		Version:    0,
		Expiration: expiration,
		QuestsID:   questID,
		Quests: []any{
			scoreQuest,
			destroyEnemiesQuest,
			surviveQuest,
		},
	}, nil
}

func RandomDestroyEnemiesQuest(levelIdFunc LevelIDWithNameFunc) (*DestroyEnemiesQuest, error) {
	type LevelEnemyCountFor100XP struct {
		Levels []Level
		Enemy  EnemyType
		Count  int
	}

	level_kill_requirements := []LevelEnemyCountFor100XP{
		// Applies to any official level
		{[]Level{}, baf, 1000},
		{[]Level{}, crowder, 1000},
		{[]Level{}, inertiac, 50},
		{[]Level{}, mothership, 100},
		{[]Level{}, rollingCube, 100},
		{[]Level{}, wary, 100},
		// Applies to official levels
		{[]Level{asteroids}, asteroid, 300},
		{[]Level{asteroids}, ufo, 10},
		{[]Level{asteroids}, rollingCube, 200},
		{[]Level{hexagon}, mothership, 200},
		{[]Level{hexagon}, wary, 50},
		{[]Level{hexagon}, inertiac, 35},
		{[]Level{eskiv}, mothership, 30},
		{[]Level{waves}, baf, 2000},
		{[]Level{waves}, bafBlue, 500},
		{[]Level{waves}, bafRed, 200},
		{[]Level{waves}, mothership, 30},
		{[]Level{waves}, inertiac, 30},
		{[]Level{fury}, mothership, 100},
		{[]Level{ceasefire}, crowder, 2000},
		{[]Level{pandemonium}, spiny, 300},
		{[]Level{pandemonium}, mothership, 100},
		{[]Level{pandemonium}, mothershipSuper, 50},
		{[]Level{partitioner}, wary, 50},
		{[]Level{partitioner}, rollingCube, 40},
		{[]Level{partitioner}, rollingSphere, 100},
		{[]Level{partitioner}, inertiac, 50},
		{[]Level{partitioner}, crowder, 1000},
		{[]Level{symbiosis}, crowder, 500},
		{[]Level{symbiosis}, rollingCube, 300},
		{[]Level{symbiosis}, inertiac, 100},
		{[]Level{symbiosis}, kamikaze, 100},

		// Featured community levels
		{[]Level{cozone}, asteroid, 1000},
		{[]Level{cozone}, mothership, 10},
		{[]Level{claustrophobia}, mothership, 50},
		{[]Level{claustrophobia}, baf, 200},
		{[]Level{neo_grid}, wary, 30},
		{[]Level{neo_grid}, rollingCube, 150},
		{[]Level{virtual_dream}, asteroid, 100},
		{[]Level{virtual_dream}, baf, 400},
		{[]Level{virtual_dream}, mothership, 60},
		{[]Level{skpg}, mothership, 30},
		{[]Level{skpg}, inertiac, 15},
		{[]Level{invasion}, wary, 30},
		{[]Level{invasion}, mothership, 30},
		{[]Level{invasion}, inertiac, 30},
		{[]Level{emerald}, bafBlue, 100},
		{[]Level{emerald}, wary, 20},
		{[]Level{emerald}, mothership, 30},
		{[]Level{madreinka}, baf, 150},
		{[]Level{madreinka}, mothership, 50},
		{[]Level{bombardment}, rollingCube, 100},
		{[]Level{bombardment}, mothership, 100},
		{[]Level{bouncy_recoil}, wary, 60},
		{[]Level{bouncy_recoil}, mothership, 60},
		{[]Level{restanvi}, mothership, 1000},
		{[]Level{restanvi}, rollingCube, 100},
		{[]Level{restanvi}, wary, 50},
		// Requirements to be made by someone with access to tools
		//   which estimate amount of enemies killed in a round
		// {[]Level{linkage}, asteroid, 10000},
		// {[]Level{linkage}, rollingCube, 2000},
	}

	template := getRandomElement(level_kill_requirements)

	levels := make([]LevelIDWithName, 0)
	for _, l := range template.Levels {
		level_info, err := levelIdFunc(LevelStrings[l])
		if err != nil {
			return nil, err
		}
		levels = append(levels, *level_info)
	}
	if len(levels) == 0 {
		levels = nil
	}
	return &DestroyEnemiesQuest{Kind: destroyEnemies, Goal: int64(template.Count), Enemy: EnemyTypeStrings[template.Enemy], XP: int64(100), Levels: levels}, nil
}

func RandomReachScoreQuest(levelIdFunc LevelIDWithNameFunc) (*ReachScoreQuest, error) {
	type LevelReachScore struct {
		Level Level
		Score int
		XP    int
	}
	reach_score_requirements := []LevelReachScore{
		// Official levels
		{asteroids, 30000, 100},
		{asteroids, 60000, 200},
		{hexagon, 30000, 100},
		{hexagon, 50000, 200},
		{eskiv, 2000, 100},
		{eskiv, 4000, 200},
		{waves, 10000, 100},
		{waves, 20000, 200},
		{fury, 4000, 100},
		{fury, 7000, 200},
		{ceasefire, 30000, 100},
		{ceasefire, 50000, 200},
		{pandemonium, 10000, 100},
		{pandemonium, 20000, 200},
		{partitioner, 30000, 100},
		{partitioner, 45000, 200},
		{symbiosis, 50000, 100},
		{symbiosis, 80000, 200},

		// Featured levels
		{eskiv1024, 150, 100},
		{eskiv1024, 250, 200},
		{cozone, 20000, 100},
		{cozone, 50000, 200},
		{claustrophobia, 5000, 100},
		{claustrophobia, 9000, 200},
		{neo_grid, 8000, 100},
		{neo_grid, 14000, 200},
		{virtual_dream, 9000, 100},
		{virtual_dream, 18000, 200},
		{invasion, 7000, 100},
		{emerald, 8000, 100},
		{emerald, 14000, 200},
		{madreinka, 20000, 100},
		{madreinka, 40000, 200},
		{bombardment, 10000, 100},
		{bombardment, 25000, 200},
		{bouncy_recoil, 8000, 100},
		{bouncy_recoil, 14000, 200},
		// {restanvi, 100000, 100},
		// {restanvi, 200000, 200},
		{linkage, 15000, 100},
		{linkage, 35000, 200},
	}
	template := getRandomElement(reach_score_requirements)
	level, err := levelIdFunc(LevelStrings[template.Level])
	if err != nil {
		return nil, err
	}
	return &ReachScoreQuest{Kind: reachScore, Goal: int64(template.Score), XP: int64(template.XP), Levels: []LevelIDWithName{*level}}, nil
}

func RandomSurviveDurationQuest(levelIdFunc LevelIDWithNameFunc) (*SurviveDurationQuest, error) {
	type SurviveDuration struct {
		Level   Level
		Seconds int
		XP      int
	}
	// Don't include levels where it's possible to survive easy while not doing much, like eskiv
	survive_requirements := []SurviveDuration{
		{waves, 60, 100},
		{waves, 90, 200},
		// {waves, 120, 300},
		{fury, 24, 100},
		// {fury, 32, 300},
		{ceasefire, 60, 100},
		{ceasefire, 90, 200},
		// {ceasefire, 120, 300},
		{pandemonium, 60, 100},
		{pandemonium, 90, 200},
		// {pandemonium, 120, 300},
		// {pandemonium, 150, 400},
		// {pandemonium, 180, 500},
		{partitioner, 60, 100},
		{partitioner, 80, 200},
		// {partitioner, 100, 300},

		// Featured community levels
		{cozone, 90, 100},
		{cozone, 150, 200},
		{claustrophobia, 50, 100},
		{neo_grid, 60, 100},
		{virtual_dream, 90, 100},
		{virtual_dream, 120, 200},
		{emerald, 60, 100},
		{emerald, 90, 200},
		{madreinka, 60, 100},
		{madreinka, 120, 200},
		{bombardment, 90, 100},
		{bombardment, 120, 200},
		{bouncy_recoil, 120, 100},
		{bouncy_recoil, 180, 200},
		{restanvi, 30, 100},
		{restanvi, 60, 200},
		{linkage, 60, 100},
		{linkage, 120, 200},
	}
	template := getRandomElement(survive_requirements)
	level, err := levelIdFunc(LevelStrings[template.Level])
	if err != nil {
		return nil, err
	}
	return &SurviveDurationQuest{Kind: surviveDuration, Goal: int64(template.Seconds * 30), XP: int64(template.XP), Levels: []LevelIDWithName{*level}}, nil
}
