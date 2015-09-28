package gpsdjson

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func tstCheckBool(name string, got bool, want bool) (hadFailed bool, s string) {
	if got != want {
		hadFailed = true
		s = fmt.Sprintf("Got %s : '%v'\n\nWant %s : '%v'\n",
			name, got, name, want)
	} else {
		hadFailed = false
		s = ""
	}
	return
}

func tstCheckTimeDuration(name string,
	got time.Duration, want time.Duration) (hadFailed bool, s string) {

	if want.Nanoseconds() != got.Nanoseconds() {
		hadFailed = true
		sGot := strconv.FormatInt(got.Nanoseconds(), 10)
		sWant := strconv.FormatInt(want.Nanoseconds(), 10)
		s = fmt.Sprintf("Got %s : '%s' nano seconds\n\nWant %s : '%s' nano seconds.\n",
			name, sGot, name, sWant)
	} else {
		hadFailed = false
		s = ""
	}
	return
}

func tstCheckErr(name string, got error, want error) (hadFailed bool, s string) {
	var (
		sWant   string
		sGot    string
		compare int
	)

	if want == nil {
		sWant = "<nil>"
	} else {
		sWant = want.Error()
	}

	if got == nil {
		sGot = "<nil>"
	} else {
		sGot = got.Error()
	}

	if compare = strings.Compare(sWant, sGot); compare == 0 {
		return false, ""
	} else {
		return true, fmt.Sprintf("Got %s : '%s'\n\nWant %s : '%s'\n",
			name, sGot, name, sWant)
	}
}

func tstCheckByteSlice(name string, got []byte, want []byte) (hadFailed bool, s string) {
	var (
		sWant1, sWant2 string
		sGot1, sGot2   string
		compare        int
	)

	if want == nil {
		sWant1 = "<nil>"
		sWant2 = "<nil"
	} else {
		sWant1 = fmt.Sprintf("%#v", want)
		sWant2 = string(want)
	}

	if got == nil {
		sGot1 = "<nil>"
		sGot2 = "<nil>"
	} else {
		sGot1 = fmt.Sprintf("%#v", got)
		sGot2 = string(got)
	}

	if compare = strings.Compare(sWant1, sGot1); compare == 0 {
		return false, ""
	} else {
		return true, fmt.Sprintf("Got %s : '%s'(%s)\n\nWant %s : '%s'(%s)\n",
			name, sGot1, sGot2, name, sWant1, sWant2)
	}
}
