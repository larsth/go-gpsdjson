package gpsdjson

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"github.com/juju/errors"
)

//Duration is a type that embeds time.Duration, and has support for JSON.
//
//The duration is stored in seconds, if marshaled to JSON.
type Duration struct {
	time.Duration
}

//MarshalJSON can marshal itself into valid JSON.
//
//MarshalJSON implements interface encoding/json.Marshaler
func (d *Duration) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	ns := d.Duration.Nanoseconds()
	s := strconv.FormatInt(ns, 10)
	buf.WriteString(s)
	buf.WriteString("ns")

	return buf.Bytes(), nil
}

//UnmarshalJSON can unmarshal a JSON description of itself.
//
//UnmarshalJSON implements interface encoding/json.Unmarshaler
func (d *Duration) UnmarshalJSON(data []byte) error {
	var (
		str           string
		hasUnitSuffix bool
		err           error
		buf           bytes.Buffer
	)

	if data == nil {
		return errors.Trace(ErrNilByteSlice)
	}
	if len(data) == 0 {
		return errors.Trace(ErrEmptyByteSlice)
	}

	str = strings.TrimSpace(string(data))
	if len(str) == 0 {
		return errors.Trace(ErrEmptyString)
	}
	buf.WriteString(str)

	//"300ms", "-1.5h" or "2h45m".
	//Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	hasUnitSuffix = (strings.HasSuffix(str, "ns") ||
		strings.HasSuffix(str, "us") || strings.HasSuffix(str, "µs") ||
		strings.HasSuffix(str, "ms") || strings.HasSuffix(str, "s") ||
		strings.HasSuffix(str, "m") || strings.HasSuffix(str, "h"))

	if false == hasUnitSuffix {
		//assume seconds
		if _, err = strconv.ParseInt(str, 10, 64); err != nil {
			return errors.Annotate(err, "Not an integer")
		}
		buf.WriteString("s")
	}

	if d.Duration, err = time.ParseDuration(buf.String()); err != nil {
		return errors.Annotate(err, "time.ParseDuration(string) error")
	}
	return nil
}
