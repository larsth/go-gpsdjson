package gpsdjson

import "time"

//SKY is a type used to report a sky view of the GPS satellite positions.
//
//If there is no GPS device available, or no skyview has been reported yet,
//only the "class" field will reliably be present.
//
//Many devices compute dilution of precision factors but do not include them in
//their reports. Many that do report DOPs report only HDOP, two-dimensional
//circular error.
//gpsd always passes through whatever the device actually reports, then attempts
//to fill in other DOPs by calculating the appropriate determinants in a
//covariance matrix based on the satellite view.
//DOPs may be missing if some of these determinants are singular.
//It can even happen that the device reports an error estimate in meters when the
//corresponding DOP is unavailable; some devices use more sophisticated error
//modeling than the covariance calculation.
type SKY struct {
	//Fixed: "SKY"
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Name of originating device.
	//Always? No. Type: string
	Device string `json:"device,omitempty"`

	//Time/date stamp in ISO8601 format, UTC.
	//May have a fractional part of up to .001sec precision.
	//Always? No. Type: time.Time
	Time time.Time `json:"time,omitempty"`

	//Longitudinal dilution of precision, a dimensionless factor which
	//should be multiplied by a base UERE to get an error estimate.
	//Always? No. Type: numeric
	Xdop float64 `json:"xdop,omitempty"`

	//Latitudinal dilution of precision, a dimensionless factor which
	//should be multiplied by a base UERE to get an error estimate.
	//Always? No. Type: numeric
	Ydop float64 `json:"ydop,omitempty"`

	//Altitude dilution of precision, a dimensionless factor which
	//should be multiplied by a base UERE to get an error estimate.
	//Always? No. Type: numeric
	Vdop float64 `json:"vdop,omitempty"`

	//Time dilution of precision, a dimensionless factor which
	//should be multiplied by a base UERE to get an error estimate.
	//Always? No. Type: numeric
	Tdop float64 `json:"tdop,omitempty"`

	//Horizontal dilution of precision, a dimensionless factor which
	//should be multiplied by a base UERE to get an error estimate.
	//Always? No. Type: numeric
	Hdop float64 `json:"hdop,omitempty"`

	//Spherical dilution of precision, a dimensionless factor which
	//should be multiplied by a base UERE to get an error estimate.
	//Always? No. Type: numeric
	Pdop float64 `json:"pdop,omitempty"`

	//Hyperspherical dilution of precision, a dimensionless factor which
	//should be multiplied by a base UERE to get an error estimate.
	//Always? No. Type: numeric
	Gdop float64 `json:"gdop,omitempty"`

	//List of satellite objects in skyview
	//Always? Yes. Type: list
	Satellites []Satellite `json:"satellites"`

	//Tag string `json:"tag"`
}
