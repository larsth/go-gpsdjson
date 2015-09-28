package gpsdjson

//Satellite is a type that does not have a "class" field,
//as they are never shipped outside of a SKYReport.
type Satellite struct {
	//PRN ID of the satellite.
	//1-63 are GNSS satellites,
	//64-96 are GLONASS satellites,
	//100-164 are SBAS satellites
	//Always? Yes. Type: numeric
	PRN float64 `json:"PRN"`

	//Azimuth, degrees from true north.
	//Always? Yes. Type: numeric
	Az float64 `json:"az"`

	//Elevation in degrees.
	//Always? Yes. Type: numeric
	El float64 `json:"el"`

	//Signal strength in dB.
	//Always? Yes. Type: numeric
	Ss float64 `json:"ss"`

	//Used in current solution?
	//(SBAS/WAAS/EGNOS satellites may be flagged used if the solution has
	// corrections from them, but not all drivers make this information
	//available.
	//Always? Yes. Type: bool
	Used bool `json:"used"`
}
