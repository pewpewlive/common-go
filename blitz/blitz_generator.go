package blitz

import (
	"math/rand"
)

type Environment int

const (
	Env_Box             Environment = 0
	Env_Disc            Environment = 1
	Env_Donut           Environment = 2
	Env_Lattice         Environment = 3
	Env_LeftEqualDanger Environment = 4
)

type Weapon int

const (
	Weapon_None Weapon = iota
	Weapon_Single
	Weapon_TicToc
	Weapon_Double
	Weapon_Triple
	Weapon_FourDirections
	Weapon_DoubleSwipe
	Weapon_Hemisphere
	Weapon_Laser
	Weapon_Shotgun
)

type SandboxKey int

const (
	key_Asteroids SandboxKey = iota
	key_Atomizebombs
	key_BAFs
	key_BlueBAFs
	key_Brownians
	key_Crowders
	key_CyanMotherships
	key_Dodgers
	key_Environment
	key_Environmentsize
	key_Exploders
	key_Freezebombs
	key_GreenMotherships
	key_Inertiacs
	key_Kamikaze
	key_Lassolength
	key_Maceboxfrequency
	key_Macecount
	key_OrangeMotherships
	key_PinkMotherships
	key_Playerspeed
	key_Playertype
	key_RedBAFs
	key_RedMotherships
	key_Repulsivebombs
	key_RollingCubes
	key_RollingSpheres
	key_Scoreboxfrequency
	key_Shieldboxfrequency
	key_Shieldcount
	key_Smallatomizebombs
	key_Smallfreezebombs
	key_Speedboxfrequency
	key_Speedboxvalue
	key_Spinies
	key_Splitters
	key_SuperCyanMotherships
	key_SuperGreenMotherships
	key_SuperOrangeMotherships
	key_SuperPinkMotherships
	key_SuperRedMotherships
	key_UFOs
	key_Waries
	key_Weaponboxfrequency
	key_Weaponboxshootingfrequency
	key_Weaponboxtype
	key_Weaponfrequency
	key_Weapontype
)

func GetSandboxKeyMappings() map[SandboxKey]string {
	return map[SandboxKey]string{
		key_Asteroids:                  "Asteroids",
		key_Atomizebombs:               "Atomize bombs",
		key_BAFs:                       "BAFs",
		key_BlueBAFs:                   "Blue BAFs",
		key_Brownians:                  "Brownians",
		key_Crowders:                   "Crowders",
		key_CyanMotherships:            "Cyan Motherships",
		key_Dodgers:                    "Dodgers",
		key_Environment:                "Environment",
		key_Environmentsize:            "Environment size",
		key_Exploders:                  "Exploders",
		key_Freezebombs:                "Freeze bombs",
		key_GreenMotherships:           "Green Motherships",
		key_Inertiacs:                  "Inertiacs",
		key_Kamikaze:                   "Kamikaze",
		key_Lassolength:                "Lasso length",
		key_Maceboxfrequency:           "Mace box frequency",
		key_Macecount:                  "Mace count",
		key_OrangeMotherships:          "Orange Motherships",
		key_PinkMotherships:            "Pink Motherships",
		key_Playerspeed:                "Player speed",
		key_Playertype:                 "Player type",
		key_RedBAFs:                    "Red BAFs",
		key_RedMotherships:             "Red Motherships",
		key_Repulsivebombs:             "Repulsive bombs",
		key_RollingCubes:               "Rolling Cubes",
		key_RollingSpheres:             "Rolling Spheres",
		key_Scoreboxfrequency:          "Score box frequency",
		key_Shieldboxfrequency:         "Shield box frequency",
		key_Shieldcount:                "Shield count",
		key_Smallatomizebombs:          "Small atomize bombs",
		key_Smallfreezebombs:           "Small freeze bombs",
		key_Speedboxfrequency:          "Speed box frequency",
		key_Speedboxvalue:              "Speed box value",
		key_Spinies:                    "Spinies",
		key_Splitters:                  "Splitters",
		key_SuperCyanMotherships:       "Super Cyan Motherships",
		key_SuperGreenMotherships:      "Super Green Motherships",
		key_SuperOrangeMotherships:     "Super Orange Motherships",
		key_SuperPinkMotherships:       "Super Pink Motherships",
		key_SuperRedMotherships:        "Super Red Motherships",
		key_UFOs:                       "UFOs",
		key_Waries:                     "Waries",
		key_Weaponboxfrequency:         "Weapon box frequency",
		key_Weaponboxshootingfrequency: "Weapon box shooting frequency",
		key_Weaponboxtype:              "Weapon box type",
		key_Weaponfrequency:            "Weapon frequency",
		key_Weapontype:                 "Weapon type"}
}

func constant(value int) func() int {
	// The returned function is a closure, capturing the 'value'
	// from the outer function's scope.
	return func() int {
		return value
	}
}

// Returns a function that returns one randomly chosen value from `values`
func pickFromSet(values []int) func() int {
	return func() int {
		index := rand.Intn(len(values))
		value := values[index]
		return value
	}
}

func pickFromInclusiveRange(min int, max int) func() int {
	return func() int {
		return rand.Intn(1+max-min) + min
	}
}

func weapon(weapon Weapon) func() int {
	return func() int {
		return int(weapon)
	}
}

func randomEnvironement() func() int {
	return randomEnvironementExluding([]Environment{})
}

func randomWeaponInSet(weapons []Weapon) func() int {
	return func() int {
		index := rand.Intn(len(weapons))
		value := weapons[index]
		return int(value)
	}
}

func randomEnvironementExluding(blacklist []Environment) func() int {
	return func() int {
		allEnvs := []Environment{Env_Box, Env_Disc, Env_Donut, Env_Lattice, Env_LeftEqualDanger}
		selectedEnvs := make([]Environment, 0, len(allEnvs)-len(blacklist))
		for _, env := range allEnvs {
			found := false
			for _, blacklistedEnv := range blacklist {
				if env == blacklistedEnv {
					found = true
					break
				}
			}
			if !found {
				selectedEnvs = append(selectedEnvs, env)
			}
		}
		index := rand.Intn(len(selectedEnvs))
		value := selectedEnvs[index]
		return int(value)
	}
}

type SandboxConfigTemplate struct {
	name string
	keys map[SandboxKey]func() int
}

var templates = []SandboxConfigTemplate{
	{
		// Mothership, bullet
		name: "Mothership vs Bullets",
		keys: map[SandboxKey]func() int{
			key_Environment:                randomEnvironement(),
			key_Environmentsize:            constant(2),
			key_Playerspeed:                constant(2),
			key_RedMotherships:             constant(3),
			key_Shieldcount:                constant(2),
			key_Weaponboxshootingfrequency: constant(4),
			key_Weaponboxtype:              randomWeaponInSet([]Weapon{Weapon_None, Weapon_DoubleSwipe, Weapon_Shotgun, Weapon_Laser}),
			key_Weaponfrequency:            constant(6),
			key_Weapontype:                 randomWeaponInSet([]Weapon{Weapon_Double, Weapon_Triple}),
		},
	},
	{
		// 15hz Triple with asteroids, inertiacs, UFOs, red BAFs
		name: "Asteroids vs Bullets",
		keys: map[SandboxKey]func() int{
			key_Asteroids:          constant(2),
			key_Environment:        pickFromInclusiveRange(1, 3),
			key_Environmentsize:    constant(4),
			key_Freezebombs:        constant(1),
			key_Inertiacs:          constant(1),
			key_RedBAFs:            constant(1),
			key_Shieldboxfrequency: constant(1),
			key_Shieldcount:        constant(3),
			key_UFOs:               constant(1),
			key_Weaponfrequency:    constant(7),
			key_Weapontype:         weapon(Weapon_Triple),
		},
	},
	{
		// 10hz Double with rolling cubes, waries, and small atomize bombs
		name: "Rolling Cubes & Waries vs Bullets",
		keys: map[SandboxKey]func() int{
			key_Environment:                constant(0),
			key_Environmentsize:            constant(0),
			key_RollingCubes:               constant(2),
			key_Shieldboxfrequency:         constant(1),
			key_Smallatomizebombs:          constant(2),
			key_Waries:                     constant(2),
			key_Weaponboxfrequency:         constant(1),
			key_Weaponboxshootingfrequency: constant(6),
			key_Weaponboxtype:              weapon(Weapon_Triple),
			key_Weaponfrequency:            constant(6),
			key_Weapontype:                 weapon(Weapon_Double),
		},
	},
	{
		// 10hz Double with crowders and brownians, no shield powerups
		name: "Crowders vs Bullets",
		keys: map[SandboxKey]func() int{
			key_Brownians:                  constant(1),
			key_Crowders:                   constant(2),
			key_Environment:                pickFromInclusiveRange(0, 2),
			key_Shieldcount:                constant(4),
			key_Weaponboxfrequency:         constant(1),
			key_Weaponboxshootingfrequency: constant(7),
			key_Weaponboxtype:              weapon(Weapon_Triple),
			key_Weaponfrequency:            constant(6),
			key_Weapontype:                 weapon(Weapon_Double),
		},
	},
	{
		// 1hz Hemisphere with yellow and red BAFs, and optional blue BAFs
		name: "BAFs vs 1Hz Hemisphere",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironementExluding([]Environment{Env_Lattice, Env_LeftEqualDanger}),
			key_BAFs:               constant(4),
			key_Environmentsize:    constant(2),
			key_Playerspeed:        constant(2),
			key_RedBAFs:            constant(3),
			key_BlueBAFs:           pickFromSet([]int{0, 3}),
			key_Shieldboxfrequency: constant(1),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_Hemisphere),
		},
	},
	{
		// 1Hz Hemisphere with blue BAFs and inertiac, and optional Exploder
		name: "Blue BAFs and inertiac vs 1Hz Hemisphere",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironementExluding([]Environment{Env_Lattice, Env_LeftEqualDanger}),
			key_BlueBAFs:           constant(4),
			key_Environmentsize:    constant(2),
			key_Inertiacs:          constant(1),
			key_Exploders:          pickFromInclusiveRange(0, 1),
			key_Playerspeed:        constant(2),
			key_Shieldboxfrequency: constant(1),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_Hemisphere),
		},
	},
	{
		// 2Hz Hemisphere with lots of kamikaze, a few exploders
		name: "Kamikaze vs 2Hz Hemisphere",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironementExluding([]Environment{Env_LeftEqualDanger}),
			key_Environmentsize:    constant(2),
			key_Exploders:          constant(1),
			key_Kamikaze:           constant(4),
			key_Playerspeed:        constant(2),
			key_Shieldboxfrequency: constant(2),
			key_Weaponfrequency:    constant(1),
			key_Weapontype:         weapon(Weapon_Hemisphere),
		},
	},
	{
		// 3Hz Hemisphere with lots of UFOs
		name: "UFOs vs 3Hz Hemisphere",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    pickFromInclusiveRange(2, 3),
			key_Playerspeed:        constant(2),
			key_Shieldboxfrequency: constant(2),
			key_UFOs:               constant(3),
			key_Weaponfrequency:    constant(2),
			key_Weapontype:         weapon(Weapon_Hemisphere),
		},
	},
	{
		// Multiple enemies, weapon shotgun, large area
		name: "Multiple enemies vs Shotgun",
		keys: map[SandboxKey]func() int{
			key_Environment:                randomEnvironement(),
			key_Asteroids:                  constant(1),
			key_BAFs:                       constant(1),
			key_Environmentsize:            constant(4),
			key_Inertiacs:                  constant(1),
			key_Kamikaze:                   constant(1),
			key_Playerspeed:                constant(2),
			key_RedMotherships:             constant(1),
			key_Shieldboxfrequency:         constant(1),
			key_Shieldcount:                constant(4),
			key_Speedboxfrequency:          constant(1),
			key_Spinies:                    constant(2),
			key_SuperCyanMotherships:       constant(1),
			key_SuperOrangeMotherships:     constant(1),
			key_UFOs:                       constant(2),
			key_Weaponboxshootingfrequency: constant(3),
			key_Weaponboxtype:              weapon(Weapon_Triple),
			key_Weaponfrequency:            constant(5),
			key_Weapontype:                 weapon(Weapon_Shotgun),
		},
	},
	{
		// 1hz Shotgun with pink, orange, green motherships, and score boxes
		name: "Motherships vs 1Hz Shotgun",
		keys: map[SandboxKey]func() int{
			key_Environment:        pickFromInclusiveRange(0, 2),
			key_Environmentsize:    constant(1),
			key_GreenMotherships:   constant(1),
			key_OrangeMotherships:  constant(1),
			key_PinkMotherships:    constant(2),
			key_Repulsivebombs:     constant(2),
			key_Scoreboxfrequency:  constant(2),
			key_Shieldboxfrequency: constant(1),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_Shotgun),
		},
	},
	{
		// Mothership, laser
		name: "Mothership vs Laser",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    constant(2),
			key_Playerspeed:        constant(2),
			key_RedMotherships:     constant(3),
			key_Shieldboxfrequency: constant(2),
			key_Shieldcount:        constant(2),
			key_Weaponfrequency:    constant(2),
			key_Weapontype:         weapon(Weapon_Laser),
		},
	},
	{
		// Multiple super motherships, bombs, large area
		name: "Supermotherships vs Bombs",
		keys: map[SandboxKey]func() int{
			key_Environment:           randomEnvironement(),
			key_Environmentsize:       constant(3),
			key_Freezebombs:           constant(2),
			key_Playerspeed:           constant(3),
			key_GreenMotherships:      constant(2),
			key_SuperGreenMotherships: constant(1),
			key_SuperPinkMotherships:  constant(3),
			key_Shieldboxfrequency:    constant(3),
			key_Shieldcount:           constant(5),
			key_Scoreboxfrequency:     constant(3),
			key_Weapontype:            weapon(Weapon_None),
		},
	},
	{
		// Lots cyan super motherships, bombs
		name: "Cyan super motherships vs Bombs",
		keys: map[SandboxKey]func() int{
			key_Environment:          randomEnvironement(),
			key_Asteroids:            constant(1),
			key_BlueBAFs:             constant(1),
			key_Environmentsize:      pickFromInclusiveRange(2, 3),
			key_Freezebombs:          constant(2),
			key_Playerspeed:          constant(3),
			key_Shieldcount:          constant(3),
			key_Smallfreezebombs:     constant(3),
			key_SuperCyanMotherships: constant(4),
			key_Weaponfrequency:      constant(4),
			key_Weapontype:           weapon(Weapon_Shotgun),
		},
	},
	{
		// Lots of red BAFs, player can't shoot, must use bombs
		name: "Red BAFs vs Bombs",
		keys: map[SandboxKey]func() int{
			key_Environment:       randomEnvironement(),
			key_Environmentsize:   constant(2),
			key_Playerspeed:       constant(2),
			key_RedBAFs:           constant(4),
			key_Shieldcount:       constant(2),
			key_Smallatomizebombs: constant(2),
			key_Weapontype:        weapon(Weapon_None),
		},
	},
	{
		// Lasso with lots of rolling spheres
		name: "Rolling Spheres vs Lasso",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    pickFromInclusiveRange(2, 3),
			key_Lassolength:        constant(6),
			key_Playerspeed:        constant(2),
			key_RollingSpheres:     constant(4),
			key_Shieldboxfrequency: constant(2),
			key_Shieldcount:        constant(2),
			key_Speedboxfrequency:  constant(2),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_None),
		},
	},
	{
		// Lasso with inertiacs and exploders
		name: "Inertiacs and exploders vs Lasso",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    pickFromInclusiveRange(2, 3),
			key_Exploders:          constant(1),
			key_Inertiacs:          constant(3),
			key_Lassolength:        constant(6),
			key_Playerspeed:        constant(2),
			key_Shieldboxfrequency: constant(2),
			key_Shieldcount:        constant(2),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_None),
		},
	},
	{
		// Lasso with crowders, and optional brownians
		name: "Crowders vs Lasso",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    pickFromInclusiveRange(2, 3),
			key_Brownians:          pickFromInclusiveRange(0, 1),
			key_Crowders:           constant(3),
			key_Lassolength:        constant(6),
			key_Playerspeed:        constant(2),
			key_Shieldboxfrequency: constant(2),
			key_Shieldcount:        constant(2),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_None),
		},
	},
	{
		// Lasso with kamikaze, and optional rolling cubes
		name: "Kamikaze vs Lasso",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    pickFromInclusiveRange(2, 3),
			key_Kamikaze:           constant(3),
			key_Lassolength:        constant(6),
			key_Playerspeed:        constant(2),
			key_RollingCubes:       pickFromInclusiveRange(0, 2),
			key_Shieldboxfrequency: constant(2),
			key_Shieldcount:        constant(2),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_None),
		},
	},

	{
		// Lasso with red motherships, and optional rolling cubes
		name: "Red motherships vs Lasso",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    pickFromInclusiveRange(2, 3),
			key_Lassolength:        constant(6),
			key_Playerspeed:        constant(2),
			key_RedMotherships:     constant(3),
			key_RollingCubes:       pickFromInclusiveRange(0, 2),
			key_Shieldboxfrequency: constant(2),
			key_Shieldcount:        constant(2),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_None),
		},
	},
	{
		// Lasso with Crowders, Inertiacs, Kamikaze
		name: "Crowders, Inertiacs, Kamikaze vs Lasso",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    pickFromInclusiveRange(2, 3),
			key_Crowders:           constant(2),
			key_Inertiacs:          constant(2),
			key_Kamikaze:           constant(2),
			key_Lassolength:        constant(6),
			key_Playerspeed:        constant(2),
			key_Shieldboxfrequency: constant(2),
			key_Shieldcount:        constant(2),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_None),
		},
	},
	{
		name: "Motherships vs Maces",
		keys: map[SandboxKey]func() int{
			key_Environment:       randomEnvironement(),
			key_Environmentsize:   constant(3),
			key_CyanMotherships:   constant(2),
			key_GreenMotherships:  constant(2),
			key_Macecount:         pickFromSet([]int{2, 3, 4, 5}),
			key_OrangeMotherships: constant(2),
			key_PinkMotherships:   constant(2),
			key_Playerspeed:       constant(2),
			key_RedMotherships:    constant(2),
			key_Shieldcount:       constant(2),
			key_Speedboxfrequency: constant(1),
			key_Speedboxvalue:     constant(1),
			key_Weaponfrequency:   constant(0),
			key_Weapontype:        weapon(Weapon_None),
		},
	},
	{
		// Maces with Kamikaze, Exploders
		name: "Kamikaze vs Maces",
		keys: map[SandboxKey]func() int{
			key_Environment:       randomEnvironement(),
			key_Environmentsize:   constant(3),
			key_Exploders:         constant(1),
			key_Kamikaze:          constant(4),
			key_Macecount:         pickFromSet([]int{2, 3, 4, 5}),
			key_Playerspeed:       constant(2),
			key_Shieldcount:       constant(2),
			key_Speedboxfrequency: constant(2),
			key_Speedboxvalue:     constant(1),
			key_Weaponfrequency:   constant(0),
			key_Weapontype:        weapon(Weapon_None),
		},
	},
	{
		name: "Asteroids vs Maces",
		keys: map[SandboxKey]func() int{
			key_Environment:       randomEnvironement(),
			key_Asteroids:         constant(4),
			key_Environmentsize:   constant(3),
			key_Macecount:         pickFromSet([]int{2, 3, 4, 5}),
			key_Playerspeed:       constant(2),
			key_Shieldcount:       constant(2),
			key_Speedboxfrequency: constant(2),
			key_Speedboxvalue:     constant(1),
			key_Weaponfrequency:   constant(0),
			key_Weapontype:        weapon(Weapon_None),
		},
	},
	{
		name: "Super Motherships + Asteroids vs Maces",
		keys: map[SandboxKey]func() int{
			key_Environment:          randomEnvironement(),
			key_Asteroids:            constant(2),
			key_Environmentsize:      constant(3),
			key_Macecount:            pickFromSet([]int{2, 3, 4, 5}),
			key_Playerspeed:          constant(2),
			key_Shieldcount:          constant(2),
			key_Speedboxfrequency:    constant(2),
			key_Speedboxvalue:        constant(1),
			key_SuperCyanMotherships: constant(3),
			key_Weaponfrequency:      constant(0),
			key_Weapontype:           weapon(Weapon_None),
		},
	},
	{
		name: "Crowders vs Maces",
		keys: map[SandboxKey]func() int{
			key_Environment:       randomEnvironement(),
			key_Crowders:          constant(4),
			key_Environmentsize:   constant(3),
			key_Macecount:         pickFromSet([]int{2, 3, 4, 5}),
			key_Playerspeed:       constant(2),
			key_Shieldcount:       constant(2),
			key_Speedboxfrequency: constant(2),
			key_Speedboxvalue:     constant(1),
			key_Weaponfrequency:   constant(0),
			key_Weapontype:        weapon(Weapon_None),
		},
	},
	{
		name: "Exploders + Inertiacs vs Maces",
		keys: map[SandboxKey]func() int{
			key_Environment:       randomEnvironement(),
			key_Environmentsize:   constant(3),
			key_Exploders:         constant(2),
			key_Inertiacs:         constant(2),
			key_Macecount:         constant(3),
			key_Playerspeed:       constant(2),
			key_Shieldcount:       constant(2),
			key_Speedboxfrequency: constant(2),
			key_Speedboxvalue:     constant(1),
			key_Weaponfrequency:   constant(0),
			key_Weapontype:        weapon(Weapon_None),
		},
	},
	{
		// Lots of yellow BAFs, player can't shoot, have to get score boxes
		name: "Yellow BAFs and score boxes",
		keys: map[SandboxKey]func() int{
			key_Environment:                randomEnvironement(),
			key_BAFs:                       constant(2),
			key_Environmentsize:            constant(2),
			key_Playerspeed:                constant(2),
			key_Scoreboxfrequency:          constant(3),
			key_Shieldboxfrequency:         constant(2),
			key_Shieldcount:                constant(2),
			key_Weaponboxshootingfrequency: constant(5),
			key_Weaponboxtype:              weapon(Weapon_None),
		},
	},
	{
		// Lots of rolling spheres, player can't shoot, have to get score boxes
		name: "Rolling spheres and score boxes",
		keys: map[SandboxKey]func() int{
			key_Environment:        randomEnvironement(),
			key_Environmentsize:    constant(2),
			key_Playerspeed:        constant(2),
			key_RollingSpheres:     constant(2),
			key_Scoreboxfrequency:  constant(4),
			key_Shieldboxfrequency: constant(2),
			key_Weaponfrequency:    constant(0),
			key_Weapontype:         weapon(Weapon_None),
		},
	},
}

// generateUniqueRandomSubset generates an array of k unique random numbers
// between 0 (inclusive) and n (exclusive).
func generateUniqueRandomSubset(k, n int) []int {
	if n > k || n < 0 || k <= 0 {
		panic("Bad arguments to generateUniqueRandomSubset")
	}
	uniqueNumbersMap := make(map[int]struct{}, n)
	result := make([]int, 0, n)
	for len(result) < n {
		num := rand.Intn(k) // Generate a random number between 0 and k-1
		// Check if the number has already been generated
		if _, exists := uniqueNumbersMap[num]; !exists {
			uniqueNumbersMap[num] = struct{}{} // Add to the set
			result = append(result, num)       // Add to the result slice
		}
	}
	return result
}

type SandboxKV struct {
	Key   string
	Value int
}

type RoundDescription struct {
	Name       string
	SandboxKVs []SandboxKV
}

// Returns `n` RoundDescriptions that are all different from each others
func GetRoundDescriptions(n int) []RoundDescription {
	indexes := generateUniqueRandomSubset(len(templates), n)
	mapping := GetSandboxKeyMappings()
	rounds := make([]RoundDescription, 0, n)
	for _, index := range indexes {
		config := templates[index].keys
		kvs := make([]SandboxKV, 0, len(config))
		for k, v := range config {
			kvs = append(kvs, SandboxKV{Key: mapping[k], Value: v()})
		}
		rounds = append(rounds, RoundDescription{Name: templates[index].name, SandboxKVs: kvs})
	}
	return rounds
}
