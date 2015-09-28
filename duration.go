package gpsdjson

import (
	"bytes"
	"strconv"
	"strings"
	"time"
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
	f := d.Duration.Seconds()
	b := make([]byte, 0, 16)
	s := strconv.FormatFloat(f, byte('f'), -1, 64)
	b = append(b, s...)
	return b, nil
}

//UnmarshalJSON can unmarshal a JSON description of itself.
//
//UnmarshalJSON implements interface encoding/json.Unmarshaler
func (d *Duration) UnmarshalJSON(data []byte) error {
	var (
		str           string
		hasUnitSuffix bool
		err           error
		duration      time.Duration
		buf           bytes.Buffer
	)

	str = strings.TrimSpace(string(data))
	buf.WriteString(str)

	//"300ms", "-1.5h" or "2h45m".
	//Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	hasUnitSuffix = (strings.HasSuffix(str, "ns") ||
		strings.HasSuffix(str, "us") || strings.HasSuffix(str, "µs") ||
		strings.HasSuffix(str, "ms") || strings.HasSuffix(str, "s") ||
		strings.HasSuffix(str, "m") || strings.HasSuffix(str, "h"))

	if false == hasUnitSuffix {
		//assume seconds (written by the MarshalJSON method)
		buf.Write([]byte("s"))
	}

	if duration, err = time.ParseDuration(buf.String()); err != nil {
		return err
	}
	d.Duration = duration
	return nil
}
