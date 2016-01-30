package gpsdjson

import (
	"errors"
	"strconv"
	"testing"
)
import "time"

func TestDurationMarshalJSON(t *testing.T) {
	var (
		timeDuration time.Duration = 64 * time.Second
		input        *Duration     = &Duration{Duration: timeDuration}
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

var tdDurationUnmarshalJSONWant = []*tdTDurationUnmarshalJSON{
	//[0]
	&tdTDurationUnmarshalJSON{
		/*a nil input will create a "s" as the only input to the
		time.ParseDuration function */
		Input:        nil,
		WantErr:      errors.New("time: invalid duration s"),
		WantDuration: time.Second * 0,
	},
	//[1]
	&tdTDurationUnmarshalJSON{
		Input:        []byte("12ns"),
		WantErr:      nil,
		WantDuration: time.Nanosecond * 12,
	},
	//[2]
	&tdTDurationUnmarshalJSON{
		Input:        []byte("23us"),
		WantErr:      nil,
		WantDuration: time.Microsecond * 23,
	},
	//[3]
	&tdTDurationUnmarshalJSON{
		Input:        []byte("34µs"),
		WantErr:      nil,
		WantDuration: time.Microsecond * 34,
	},
	//[4]
	&tdTDurationUnmarshalJSON{
		Input:        []byte("46ms"),
		WantErr:      nil,
		WantDuration: time.Millisecond * 46,
	},
	//[5]
	&tdTDurationUnmarshalJSON{
		Input:        []byte("67s"),
		WantErr:      nil,
		WantDuration: time.Second * 67,
	},
	//[6]
	&tdTDurationUnmarshalJSON{
		Input:        []byte("78m"),
		WantErr:      nil,
		WantDuration: time.Minute * 78,
	},
	//[7]
	&tdTDurationUnmarshalJSON{
		Input:        []byte("90h"),
		WantErr:      nil,
		WantDuration: time.Hour * 90,
	},
	//[8]
	&tdTDurationUnmarshalJSON{
		Input:        []byte("30"),
		WantErr:      nil,
		WantDuration: time.Second * 30,
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
	for i, testItem = range tdDurationUnmarshalJSONWant {
		d = &Duration{}
		gotErr = d.UnmarshalJSON(testItem.Input)
		hadFailed, s = tstCheckErr("returned error", gotErr, testItem.WantErr)
		if hadFailed == true {
			t.Error(s)
			t.Log("The test error ocurred at test item: "+
				"tdDurationUnmarshalJSONWant[", strconv.Itoa(i), "].")
		}
		hadFailed, s = tstCheckTimeDuration("calculated time.Duration",
			d.Duration, testItem.WantDuration)
		if hadFailed == true {
			t.Error(s)
			t.Log("The test error ocurred at test item: "+
				"tdDurationUnmarshalJSONWant[", strconv.Itoa(i), "].")
		}
		i++
	}
}

func BenchmarkDurationMarshalJSON(b *testing.B) {
	var d = &Duration{}
	d.Duration = time.Second * 30
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
