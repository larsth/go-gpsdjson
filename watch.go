package gpsdjson

//WATCH is type used both for a "?WATCH" command
//request and the response from that command.
//
//A WATCH command (CommandWATCH) sets the watcher mode.
//In watcher mode, the gpsd daemon reports are dumped as TPV and SKY
//responses.
//
//A WATCH request changes the subscriber's policy.
//A WATCH request also sets or elicits a report of per-subscriber
//policy and the raw bit.
//
//A WATCH response describes the subscriber's policy, and will also
//include a DEVICES object.
//
// Undocumented struct field
//
//There is an additional boolean "timing" struct field which is
//undocumented because that portion of the interface is considered
//unstable and for (gpsd-)developers use only.
//
//Struct field "Timing" is not implemented in this Go library.
type WATCH struct {
	//Fixed: "WATCH"
	//
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Enable (true) or disable (false) the watcher mode.
	//Default is true.
	//
	//Always? No. Type: boolean
	Enable bool `json:"enable,omitempty"`

	//Enable (true) or disable (false) dumping of JSON reports.
	//Default is false.
	//
	//Always? No. Type: boolean
	JSON bool `json:"json,omitempty"`

	//Enable (true) or disable (false)
	//dumping of binary packets as pseudo-NMEA.
	//Default is false.
	//
	//Always? No. Type: boolean
	Nmea bool `json:"nmea,omitempty"`

	//Controls 'raw' mode. When this attribute is set to 1 for a channel,
	//gpsd reports the unprocessed NMEA or AIVDM data stream from whatever
	//device is attached.
	//Binary GPS packets are hex-dumped.
	//RTCM2 and RTCM3 packets are not dumped in raw mode.
	//When this attribute is set to 2 for a channel that processes binary
	//data, gpsd reports the received data verbatim without hex-dumping.
	//
	//Always? No. Type: numeric
	Raw float64 `json:"raw,omitempty"`

	//If true, apply scaling divisors to output before dumping;
	//default is false.
	//
	//Always? No. Type: boolean
	Scaled bool `json:"scaled,omitempty"`

	//If true, aggregate AIS type24 sentence parts.
	//If false, report each part as a separate JSON object,
	//leaving the client to match MMSIs and aggregate.
	//Default is false.
	//Applies only to AIS reports.
	//
	//Always? No. Type: boolean
	Split24 bool `json:"split24,omitempty"`

	//If true, emit the TOFF JSON message on each cycle,
	// and a PPS JSON message when the device issues 1PPS.
	//Default is false.
	//
	//Always? No. Type: boolean
	Pps bool `json:"pps,omitempty"`

	//If present, enable watching only of
	//the specified device rather than all devices.
	//Useful with raw and NMEA modes in which device responses aren't tagged.
	//Has no effect when used with enable:false.
	//
	//Always? No. Type: string
	Device string `json:"device,omitempty"`

	//URL of the remote daemon reporting the watch set.
	//If empty, this is a WATCH response from the local daemon.
	//
	//Always? No. Type: string
	Remote string `json:"remote,omitempty"`
}
