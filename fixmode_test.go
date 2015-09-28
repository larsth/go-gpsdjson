package gpsdjson

import (
	"strconv"
	"strings"
	"testing"
)

type tdTFixMode struct {
	Input FixMode
	Want  string
}

var tdFixModeStringWant = []*tdTFixMode{
	&tdTFixMode{
		Input: FixNotSeen,
		Want:  "Not Seen",
	},
	&tdTFixMode{
		Input: FixNone,
		Want:  "None",
	},
	&tdTFixMode{
		Input: Fix2D,
		Want:  "2D",
	},
	&tdTFixMode{
		Input: Fix3D,
		Want:  "3D",
	},
	&tdTFixMode{
		Input: FixMode(5),
		Want:  "Unknown FixMode value",
	},
}

func TestFixModeString(t *testing.T) {
	var i int = 0
	for _, testItem := range tdFixModeStringWant {
		sGot := testItem.Input.String()
		if strings.Compare(testItem.Want, sGot) != 0 {
			t.Errorf("Got '%s'\nWant: '%s'\n", testItem.Want, sGot)
			t.Log("The test error ocurred at test item: tdFixModeStringWant[",
				strconv.Itoa(i), "].")
		}
		i++
	}
}

func BenchmarkFixModeString(b *testing.B) {
	var f = Fix3D
	for i := 0; i < b.N; i++ {
		_ = f.String()
	}
}
