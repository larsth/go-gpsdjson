package gpsdjson

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	DebugDurationUnmarshalJSON bool
)

//Duration is a type that embeds time.Duration, and has support for JSON.
//
//The duration is stored in seconds, if marshaled to JSON.
type Duration struct {
	D time.Duration
}

//MarshalJSON can marshal itself into valid JSON.
//
//MarshalJSON implements interface encoding/json.Marshaler
func (d *Duration) MarshalJSON() ([]byte, error) {
	var (
		ns  int64
		s   string
		buf bytes.Buffer
		p   []byte
	)

	ns = d.D.Nanoseconds()
	s = strconv.FormatInt(ns, 10)
	buf.WriteString(s)
	buf.WriteString("ns")

	p = make([]byte, 0, buf.Len())
	p = append(p, buf.Bytes()...)

	return p, nil
}

func (d *Duration) unmarshalJsonCut(str string) string {
	var (
		count      int
		i, j       int
		str2, str3 string
	)

	count = strings.Count(str, `"`)
	if count > 0 {
		j = strings.LastIndex(str, `"`)
		str2 = str[:j]
		i = 1 + strings.LastIndex(str2, `"`)
		str3 = str2[i:]
	} else {
		str3 = str
	}

	if DebugDurationUnmarshalJSON {
		fmt.Println("gpsdjson.*Duration.unmarshalJsonCut:")
		fmt.Printf("\t j=%d\n", j)
		fmt.Println("\t str=", str)
		fmt.Println("\t str2=", str2)
		fmt.Printf("\t i=%d\n", i)
		fmt.Println("\t str3=", str3)
	}

	return str3
}

//UnmarshalJSON can unmarshal a JSON description of itself.
//
//UnmarshalJSON implements interface encoding/json.Unmarshaler
func (d *Duration) UnmarshalJSON(data []byte) error {
	const NA = "N/A"
	var (
		str1, str2 string
		hasSuffix  bool
		err        error
		buf        bytes.Buffer
	)

	if data == nil {
		return ErrNilByteSlice
	}
	if len(data) == 0 {
		return ErrEmptyByteSlice
	}

	str1 = strings.TrimSpace(string(data))
	if len(str1) == 0 {
		return ErrEmptyString
	}
	str2 = d.unmarshalJsonCut(str1)
	buf.WriteString(str2)

	//"300ms", "-1.5h" or "2h45m".
	//Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	hasSuffix = (strings.HasSuffix(str2, "ns") ||
		strings.HasSuffix(str2, "us") || strings.HasSuffix(str2, "µs") ||
		strings.HasSuffix(str2, "ms") || strings.HasSuffix(str2, "s") ||
		strings.HasSuffix(str2, "m") || strings.HasSuffix(str2, "h"))

	if false == hasSuffix {
		if _, err = strconv.ParseInt(str2, 10, 64); err != nil {
			return err
		}
		//assume seconds
		buf.WriteString("s")
	}

	if d.D, err = time.ParseDuration(buf.String()); err != nil {
		return err
	}

	return nil
}
