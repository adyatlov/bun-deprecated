package bun

import (
	"strings"
	"testing"
)

const str = `“Would you tell me, please, which way I ought to go from here?”
“That depends a good deal on where you want to get to,” said the Cat.
“I don’t much care where-–” said Alice.
“Then it doesn’t matter which way you go,” said the Cat.
“-–so long as I get SOMEWHERE,” Alice added as an explanation.
“Oh, you’re sure to do that,” said the Cat, “if you only walk long enough.”`

func TestFindFirstLine(t *testing.T) {
	r := strings.NewReader(str)
	n, line, err := findLine(r, "SOMEWHERE")
	const expected = `“-–so long as I get SOMEWHERE,” Alice added as an explanation.`
	if line != expected {
		t.Errorf("Epected line = \"%v\", observed \"%v\"", expected, line)
	}
	if n != 5 {
		t.Errorf("Expected n = 5, observed n = %v", n)
	}
	if err != nil {
		t.Errorf("Expected err = nil, observed err = %v", err)
	}
}
