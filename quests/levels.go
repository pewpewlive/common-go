package quests

type Level int

const (
	asteroids Level = iota
	hexagon
	eskiv
	waves
	fury
	ceasefire
	pandemonium
	partitioner
	symbiosis
	eskiv1024
	cozone
	claustrophobia
	neo_grid
	virtual_dream
	skpg
	invasion
	emerald
	madreinka
	bombardment
	bouncy_recoil
	restanvi
	linkage
)

var LevelStrings = map[Level]string{
	asteroids:      "asteroids",
	hexagon:        "hexagon",
	eskiv:          "eskiv",
	waves:          "waves",
	fury:           "fury",
	ceasefire:      "ceasefire",
	pandemonium:    "pandemonium",
	partitioner:    "partitioner",
	symbiosis:      "symbiosis",
	eskiv1024:      "8moMdFFM34sHzu1YZCMHl",
	cozone:         "OqbntT5G0qJO1kxeP1jpT",
	claustrophobia: "RwTg2aQLsJKBpt3VETRQG",
	neo_grid:       "VAQ5bWlqAsAOCjiAwQlJF",
	virtual_dream:  "Fwrhv7Ki_0cWN5qLRPBY8",
	skpg:           "hS_2SUCANyKXAZD5V453b",
	invasion:       "G11woDoCr23OFeTLSeHlU",
	emerald:        "1NfUw5OQtQqHeNS1Bmyny",
	madreinka:      "e3bQkXLwZxBJlWn0XD_v5",
	bombardment:    "XhZg3l2wPdEFKIfwUbE6Q",
	bouncy_recoil:  "iuCC8YPMdzrgESaOiqKc4",
	restanvi:       "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICAqo6EoQkM",
	linkage:        "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICAws3MmwkM",
}
