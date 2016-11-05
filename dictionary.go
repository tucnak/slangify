package slangify

var British map[string]string

var Shortcuts map[string]string

func init() {
	British = map[string]string{
		"brother":    "bruv",
		"man":        "chap",
		"men":        "chaps",
		"guy":        "bloke",
		"boy":        "lad",
		"woman":      "lady",
		"women":      "ladies",
		"hello":      "ayy",
		"awesome":    "dog's bollocks",
		"brilliant":  "brill",
		"absolutely": "total",
		"crazy":      "mental",
		"insane":     "bonkers",
		"shit":       "bollocks",
		"tired":      "knackered",
		"amazed":     "gobsmacked",
		"gay":        "puff",
		"saucy":      "cheeky",
	}

	Shortcuts = map[string]string{
		"you":   "u",
		"your":  "ur",
		"we":    "v",
		"them":  "dem",
		"this":  "dis",
		"these": "dis",
		"that":  "dat",
		"eight": "8",
		"ate":   "8",
		"for":   "4",
		"to":    "2",
		"too":   "2",
	}
}
