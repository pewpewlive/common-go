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

	// Featured
	eskiv1024
	untouchable
	cozone
	summoner
	claustrophobia
	neo_grid
	virtual_dream
	skpg
	invasion
	emerald
	madreinka
	bombardment
	bouncy_recoil
	heavy_plummet
	restanvi
	linkage

	// Non featured
	deadzone
	challenge1
	waves_red
	teampewnejos
	dodge_bafs
	marching_cubes_v2
	ufomania
	polygons
	interdimensional_problems
	zale
	synthesis
	the_generator
	cubes_and_squares
	cubes_and_squares_plus
	crush_the_casuals
	inertiacs_spawner
	roulette
	xcrush
	cube_prison
	purple_world
	hexagon_challenge1
	just_pong
	hellskiv
	cube
	wrath
	pew_pong
	dont_touch_grass
	timebomb
	eskivghost
	field_of_wormholes
	_0800
	waves_pro
	inertiac_world
	numfall
	pew_man
	bolls
	pacifism_live
	challenge3
	challenge4
	simon_says
	pinctagon
	laser
	intro
	hexaskiv
	mutoxs_sandbox
	blindness
	bee_hive
	the_depths_of_agony
	duel
	waves_blue
	challenge2
	rolling_cubes_and_spinning_hexagons
	crasbaf
	angry_kouglof
	warmup
	inertiacs_rage
	khornes_box
	exvexance
	cage_theory
	stripped_hell
	felp
	baf_rain
	mothership_mix
	midnight
	triskiv
	no_we_rlkira
	deadline
)

var LevelStrings = map[Level]string{
	asteroids:   "asteroids",
	hexagon:     "hexagon",
	eskiv:       "eskiv",
	waves:       "waves",
	fury:        "fury",
	ceasefire:   "ceasefire",
	pandemonium: "pandemonium",
	partitioner: "partitioner",
	symbiosis:   "symbiosis",

	// Featured
	eskiv1024:      "8moMdFFM34sHzu1YZCMHl",
	untouchable:    "PJEieMdh5kBQIRw7ltAD2",
	cozone:         "OqbntT5G0qJO1kxeP1jpT",
	summoner:       "4SxY4CpWSQKLAEWVA63Zn",
	claustrophobia: "RwTg2aQLsJKBpt3VETRQG",
	neo_grid:       "VAQ5bWlqAsAOCjiAwQlJF",
	virtual_dream:  "Fwrhv7Ki_0cWN5qLRPBY8",
	skpg:           "hS_2SUCANyKXAZD5V453b",
	invasion:       "G11woDoCr23OFeTLSeHlU",
	emerald:        "1NfUw5OQtQqHeNS1Bmyny",
	madreinka:      "e3bQkXLwZxBJlWn0XD_v5",
	bombardment:    "XhZg3l2wPdEFKIfwUbE6Q",
	bouncy_recoil:  "iuCC8YPMdzrgESaOiqKc4",
	heavy_plummet:  "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICAuv6RrwoM",
	restanvi:       "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICAqo6EoQkM",
	linkage:        "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICAws3MmwkM",

	// Non featured
	deadzone:                            "mzaIpL6hj2OVImwwPNIwF",
	challenge1:                          "fMl0qqG4DSegLafEQFvvq",
	waves_red:                           "E7aMX6rxjacUwv8wH82Wt",
	teampewnejos:                        "15jaTz8PYGK3G7i47Jhvj",
	dodge_bafs:                          "EhHs75INFybVH_JiYKuWU",
	marching_cubes_v2:                   "tmnIEV3mrTNcrRzoZEB77",
	ufomania:                            "6cyZW8LspfvQhTNuMhGeN",
	polygons:                            "HlGj39ziiJUkGJTFKrT4H",
	interdimensional_problems:           "i3lCbEwS0VtMjtSqivd9c",
	zale:                                "RHl42tQqV4o2D0UiC5m40",
	synthesis:                           "r0wvVoXcZ1XauROJdts5L",
	the_generator:                       "KTN7DudFR2ja4uYl4i39_",
	cubes_and_squares:                   "pnoksVSPYlAcGMb62KoVh",
	cubes_and_squares_plus:              "Yzv5XQfWLo0n1jU4i8edH",
	crush_the_casuals:                   "vJqshN0BpN7wRu4nijkmQ",
	inertiacs_spawner:                   "ZlxFwr1dXdXsZslFNyzL_",
	roulette:                            "C2RHXNcDabo5KWzrw0cop",
	xcrush:                              "OMwgNUlxFJehgAuxylIOA",
	cube_prison:                         "Zo2LEt6kJAUByB37JIeK_",
	purple_world:                        "YLnrN_6dpWCHoFXiOMS2W",
	hexagon_challenge1:                  "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICAsqqUsggM",
	just_pong:                           "r6tB425YS638C1Ij_DQEx",
	hellskiv:                            "lgUZ7OQZlVc5HUHFoGgLd",
	cube:                                "bccv8gaex2JwCXYEj1yJS",
	wrath:                               "RU5LAgKefz5BZY19KLoCP",
	pew_pong:                            "xo8lychHUy7lass08Cyjb",
	dont_touch_grass:                    "7fa9_mUjTldAgdgDmrB4q",
	timebomb:                            "8Q4kCBw3ihiwZJsCp7R6b",
	eskivghost:                          "4HtPcf4TvR2A6W8lzVcRr",
	field_of_wormholes:                  "1rqj5Eiqn2gqqFFGmfSvg",
	_0800:                               "DKDEYSsEgdjElLWg7Nw0R",
	waves_pro:                           "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICAwpWehgkM",
	inertiac_world:                      "81TpajqWIlSXYmZrfjSjw",
	numfall:                             "0ApjN0EmmtQgStfklWJx2",
	pew_man:                             "Om8xrkBMgDfX4fuyJYVk1",
	bolls:                               "Qu7D3vwIMQJsKB97m2naw",
	pacifism_live:                       "jSm8oA3VvXxn4TmISdMEh",
	challenge3:                          "kB5t864Jh_eoJOng21VCq",
	challenge4:                          "DnzLdegDTaF1u0OdOnxZx",
	simon_says:                          "HpfC42TNwH5AEmvWkkQof",
	pinctagon:                           "adCPFI5GMhVayYVGSsST1",
	laser:                               "TizfulCYU06IAPPkelVnB",
	intro:                               "MjglrGsOEzNyerfjWnE9S",
	hexaskiv:                            "GQ1L23gb5ksVXKMaI2Wat",
	mutoxs_sandbox:                      "mf1ll1qtK0TC0wJbHGNgc",
	blindness:                           "ccAHd_9A1WhMf458BrRXB",
	bee_hive:                            "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICAyv2e2woM",
	the_depths_of_agony:                 "E4bXiWnXmyqdYHnkeGqUJ",
	duel:                                "bpprOLBAjF6hFL1wKvfMJ",
	waves_blue:                          "B6QvndkC2LW8O6Jp1HsjB",
	challenge2:                          "Yf55GvAzKh6aUTKMDBIpe",
	rolling_cubes_and_spinning_hexagons: "ZgllNqbQ1i5RfnvyyDwhY",
	crasbaf:                             "D7RfAt7ZhAcCe7XKnJxt5",
	angry_kouglof:                       "ag9zfnBld3Bldy1vbmxpbmVyFwsSCkxldmVsRW50cnkYgICA-rvmjQsM",
	warmup:                              "ITYMk7vLVuOfNduT9XmFP",
	inertiacs_rage:                      "9871aKvEAJ7nUY6TI6NNQ",
	khornes_box:                         "1wYWYJoSae2r3RvjvDwp3",
	exvexance:                           "MMDUFnJSiDwTl20JGDr8C",
	cage_theory:                         "N2rTjBapnHoHk4J96wBhf",
	stripped_hell:                       "slywgaLb97yKLnysnzacj",
	felp:                                "mTRQoEASAjJoYSFTCN0wP",
	baf_rain:                            "BJ2KMIkQBKCdzI0xq0ABc",
	mothership_mix:                      "JslegBRVETOgbjSkohjBg",
	midnight:                            "Eeqlk0cYSvpRhEYbJnEGq",
	triskiv:                             "APYMJWtMO1sIHhhGjp6iB",
	no_we_rlkira:                        "uAeBBrmuzKIqLNQKwM1HG",
	deadline:                            "9lLsHWhsO4SxfAOErBHOe",
}
