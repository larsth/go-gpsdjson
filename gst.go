package gpsdjson

//GST is a type for pseudorange noise reports.
type GST struct {
	//Fixed: "GST"
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Name of originating device
	//Always? No. Type: numeric
	Device string `json:"device,omitempty"`

	//Seconds since the Unix epoch, UTC.
	//May have a fractional part of up to .001sec precision.
	//Always? No. Type: time.Duration
	Time Duration `json:"time,omitempty"`

	//Value of the standard deviation of the range inputs to the navigation
	//process (range inputs include pseudoranges and DGPS corrections).
	//Always? No. Type: numeric
	Rms float64 `json:"rms,omitempty"`

	//Standard deviation of semi-major axis of error ellipse, in meters.
	//Always? No. Type: numeric
	Major float64 `json:"major,omitempty"`

	//Standard deviation of semi-minor axis of error ellipse, in meters.
	//Always? No. Type: numeric
	Minor float64 `json:"minor,omitempty"`

	//Orientation of semi-major axis of error ellipse, in degrees from true north.
	//Always? No. Type: numeric
	Orient float64 `json:"orient,omitempty"`

	//Standard deviation of latitude error, in meters.
	//Always? No. Type: numeric
	Lat float64 `json:"lat,omitempty"`

	//Standard deviation of longitude error, in meters.
	//Always? No. Type: numeric
	Lon float64 `json:"lon,omitempty"`

	//Standard deviation of altitude error, in meters.
	//Always? No. Type: numeric
	Alt float64 `json:"alt,omitempty"`

	//Tag    string    `json:"tag"`
}
