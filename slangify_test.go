package slangify

import "testing"

func TestSlangify(t *testing.T) {
	cases := map[string]string{
		"Hello brother!":         "ayy bruv!",
		"The film was awesome":   "the film was dog's bollocks",
		"This boy is too insane": "dis lad is 2 bonkers",
		"You were kind of saucy": "u were kind of cheeky",
	}

	for from, to := range cases {
		cooked := Slangify(from)

		if cooked != to {
			t.Errorf("bullshit\nExpected: \"%s\"\nRecieved: \"%s\"\n",
				to, cooked)
		}
	}
}
