package gpsdjson

import "github.com/larsth/go-gpsfix"

//TPV is type that  contains a time-position-velocity report.
//
//The "class" and "mode" fields will reliably be present.
//The "mode" field will be emitted before optional fields that may be
//absent when there is no fix.
//
//Error estimates will be emitted after the fix components they're
//associated with.
//
//Others may be reported or not depending on the fix quality.
//
//The documentation text for this type is from
// http://catb.org/gpsd/gpsd_json.html
type TPV struct {
	//Fixed: "TPV".
	//
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Name of originating device.
	//
	//Always? Yes. Type: string
	Device string `json:"device"`

	//NMEA mode: %d, 0=no mode value yet seen, 1=no fix, 2=2D, 3=3D.
	//
	////Always? Yes. Type: Mode
	Fix gpsfix.FixMode `json:"mode"`

	//Time/date stamp in ISO8601 format, UTC.
	//May have a fractional part of up to .001sec precision.
	//May be absent if mode is not 2 or 3.
	//
	//Always? No. Type: time.Time
	Time string `json:"time,omitempty"`

	//Estimated timestamp error (%f, seconds, 95% confidence).
	//Present if time is present.
	//
	//Always? No. Type: numeric
	Ept float64 `json:"ept,omitempty"`

	//Latitude in degrees: +/- signifies North/South.
	//Present when mode is 2 (2D fix) or 3 (3D fix).
	//
	//Always? No. Type: numeric
	Lat float64 `json:"lat,omitempty"`

	//Longitude in degrees: +/- signifies East/West.
	//Present when mode is 2 (2D fix) or 3 (3D fix).
	//
	//Always? No. Type: numeric
	Lon float64 `json:"lon,omitempty"`

	//Altitude in meters.
	//Present if mode is 3 (3D fix).
	//
	//Always? No. Type: numeric
	Alt float64 `json:"alt,omitempty"`

	//Longitude error estimate in meters, 95% confidence.
	//Present if mode is 2 (2D fix) or 3 (3D fix),
	//and DOPs can be calculated from the satellite view.
	//
	//Always? No. Type: numeric
	Epx float64 `json:"epx,omitempty"`

	//Latitude error estimate in meters, 95% confidence.
	//Present if mode is 2 (2D fix) or 3 (3D fix),
	//and DOPs can be calculated from the satellite view.
	//
	//Always? No. Type: numeric
	Epy float64 `json:"epy,omitempty"`

	//Estimated vertical error in meters, 95% confidence.
	//Present if mode is 2 (2D fix) or 3 (3D fix),
	//and DOPs can be calculated from the satellite view.
	//
	//Always? No. Type: numeric
	Epv float64 `json:"epv,omitempty"`

	//Course over ground, degrees from true north.
	//
	//Always? No. Type: numeric
	Track float64 `json:"track,omitempty"`

	//Speed over ground, meters/sec.
	//
	//Always? No. Type: numeric
	Speed float64 `json:"speed,omitempty"`

	//Climb (positive) or sink (negative) rate, meters/sec.
	//
	//Always? No. Type: numeric
	Climb float64 `json:"climb,omitempty"`

	//Direction error estimate in degrees, 95% confidence.
	//
	//Always? No. Type: numeric
	Epd float64 `json:"epd,omitempty"`

	//Speed error estinmate in meters/sec, 95% confidence.
	//
	//Always? No. Type: numeric
	Eps float64 `json:"eps,omitempty"`

	//Climb/sink error estimate in meters/sec, 95% confidence.
	//
	//Always? No. Type: numeric
	Epc float64 `json:"epc,omitempty"`

	//Tag string `json:"tag,omitempty"`
}
