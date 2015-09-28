package gpsdjson

//POLL is a type that contains the data from the response from a
//a "?POLL" command.
//
//The POLL command requests data from the last-seen fixes on all active
//GPS devices.
//Devices must previously have been activated by ?WATCH to be pollable.
//
//Polling can lead to possibly surprising results when it is used on a
//device such as an NMEA GPS for which a complete fix has to be accumulated
//from several sentences. If you poll while those sentences are being emitted,
//the response will contain the last complete fix data and may be as much as
//one cycle time (typically 1 second) stale.
//
//The POLL response will contain a timestamped list of TPV objects describing
//cached data, and a timestamped list of SKY objects describing satellite
//configuration.
//If a device has not seen fixes, it will be reported with a mode field of zero.
//
// Note
//
//Client software should not assume the field inventory of the POLL response
//is fixed for all time.
//As gpsd collects and caches more data from more sensor types, those data
//are likely to find their way into this response.
type POLL struct {
	//Fixed: "POLL"
	//
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Timestamp in ISO 8601 format.
	//May have a fractional part of up to .001sec precision.
	//
	//Always? Yes. Type:
	Time Duration `json:"time"`

	//Count of active devices.
	//
	//Always? Yes. Type: numeric
	Active float64 `json:"active"`

	//Comma-separated list of TPV objects.
	//
	//Always? Yes. Type: JSON Array
	Fixes []TPV `json:"tpv"`

	//Comma-separated list of SKY objects.
	//
	//Always? Yes. Type: JSON Array
	SkyViews []SKY `json:"sky"`
}
