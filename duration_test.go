package gpsdjson

import (
	"fmt"
	"strconv"
	"testing"
)
import "time"

func TestDurationMarshalJSON(t *testing.T) {
	var (
		timeDuration time.Duration = 64 * time.Second
		input        *Duration     = &Duration{D: timeDuration}
		gotBytes     []byte
		gotErr       error
		//wantBytes: `64000000000ns`
		wantBytes = []byte{
			0x36, 0x34,
			0x30, 0x30, 0x30,
			0x30, 0x30, 0x30,
			0x30, 0x30, 0x30,
			0x6e, 0x73}
		wantErr error = nil

		hadFailed bool
		s         string
	)

	gotBytes, gotErr = input.MarshalJSON()
	hadFailed, s = tstCheckErr("returned error", gotErr, wantErr)
	if hadFailed == true {
		t.Error(s)
	}
	hadFailed, s = tstCheckByteSlice("returned byte slice", gotBytes, wantBytes)
	if hadFailed == true {
		t.Error(s)
	}
}

type tdTDurationUnmarshalJSON struct {
	Input        []byte
	WantErr      error
	WantDuration time.Duration
}

var tdDurationUnmarshalJSON = []*tdTDurationUnmarshalJSON{
	//[0]
	&tdTDurationUnmarshalJSON{
		/*a nil input will create a "s" as the only input to the
		time.ParseDuration function */
		Input:        nil,
		WantErr:      ErrNilByteSlice,
		WantDuration: time.Second * 0,
	},
	//[1]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`12ns`),
		WantErr:      nil,
		WantDuration: time.Nanosecond * 12,
	},
	//[2]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"12ns"`),
		WantErr:      nil,
		WantDuration: time.Nanosecond * 12,
	},
	//[3]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`23us`),
		WantErr:      nil,
		WantDuration: time.Microsecond * 23,
	},
	//[4]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"23us"`),
		WantErr:      nil,
		WantDuration: time.Microsecond * 23,
	},
	//[5]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`34µs`),
		WantErr:      nil,
		WantDuration: time.Microsecond * 34,
	},
	//[6]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"34µs"`),
		WantErr:      nil,
		WantDuration: time.Microsecond * 34,
	},
	//[7]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`46ms`),
		WantErr:      nil,
		WantDuration: time.Millisecond * 46,
	},
	//[8]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"46ms"`),
		WantErr:      nil,
		WantDuration: time.Millisecond * 46,
	},
	//[9]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`34µs`),
		WantErr:      nil,
		WantDuration: time.Microsecond * 34,
	},
	//[10]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"34µs"`),
		WantErr:      nil,
		WantDuration: time.Microsecond * 34,
	},
	//[11]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`67s`),
		WantErr:      nil,
		WantDuration: time.Second * 67,
	},
	//[12]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"67s"`),
		WantErr:      nil,
		WantDuration: time.Second * 67,
	},
	//[13]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`78m`),
		WantErr:      nil,
		WantDuration: time.Minute * 78,
	},
	//[14]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"78m"`),
		WantErr:      nil,
		WantDuration: time.Minute * 78,
	},
	//[15]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`90h`),
		WantErr:      nil,
		WantDuration: time.Hour * 90,
	},
	//[16]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"90h"`),
		WantErr:      nil,
		WantDuration: time.Hour * 90,
	},
	//[17]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`30`),
		WantErr:      nil,
		WantDuration: time.Second * 30,
	},
	//[18]
	&tdTDurationUnmarshalJSON{
		Input:        []byte(`"ticker-duration":"30"`),
		WantErr:      nil,
		WantDuration: time.Second * 30,
	},
	//[19]
	&tdTDurationUnmarshalJSON{
		/*is whitespace trimmed (removed)?, if successful, then yes.*/
		Input:        []byte("  "),
		WantErr:      ErrEmptyString,
		WantDuration: time.Second * 0,
	},
	//[20]
	&tdTDurationUnmarshalJSON{
		/*value 'a' is not valid input: Test the returned error. */
		Input: []byte("a"),
		WantErr: fmt.Errorf("%s",
			`strconv.ParseInt: parsing "a": invalid syntax`),
		WantDuration: time.Second * 0,
	},
	//[21]
	&tdTDurationUnmarshalJSON{
		/*value 's' is not valid input: Test the returned error. */
		Input:        []byte("s"),
		WantErr:      fmt.Errorf("%s", `time: invalid duration s`),
		WantDuration: time.Second * 0,
	},
	//[22]
	&tdTDurationUnmarshalJSON{
		/*The empty string '' is not valid input: Test the returned error. */
		Input:        []byte(``),
		WantErr:      fmt.Errorf("%s", `non-nil, but empty byte slice`),
		WantDuration: time.Second * 0,
	},
}

func TestDurationUnmarshalJSON(t *testing.T) {
	var (
		d      *Duration
		gotErr error

		hadFailed bool
		s         string
		testItem  *tdTDurationUnmarshalJSON
		i         int
	)
	for i, testItem = range tdDurationUnmarshalJSON {
		d = &Duration{}
		gotErr = d.UnmarshalJSON(testItem.Input)
		hadFailed, s = tstCheckErr("returned error", gotErr, testItem.WantErr)
		if hadFailed == true {
			t.Error(s)
			t.Log("The test error ocurred at test item: "+
				"tdDurationUnmarshalJSON[", strconv.Itoa(i), "].")
		}
		hadFailed, s = tstCheckTimeDuration("calculated time.Duration",
			d.D, testItem.WantDuration)
		if hadFailed == true {
			t.Error(s)
			t.Log("The test error ocurred at test item: "+
				"tdDurationUnmarshalJSONWant[", strconv.Itoa(i), "].")
		}
		i++
	}
}

func TestDurationUnmarshalJSONWithDebugging(t *testing.T) {
	DebugDurationUnmarshalJSON = true
	TestDurationUnmarshalJSON(t)
	DebugDurationUnmarshalJSON = false
}

func BenchmarkDurationMarshalJSON(b *testing.B) {
	var d = &Duration{}
	d.D = time.Second * 30
	for i := 0; i < b.N; i++ {
		_, _ = d.MarshalJSON()
	}
}

func BenchmarkDurationUnmarshalJSON(b *testing.B) {
	var (
		d     = &Duration{}
		input = []byte("30")
	)
	for i := 0; i < b.N; i++ {
		_ = d.UnmarshalJSON(input)
	}
}
