package gpsdjson

//TOFF is a type that contains the message that is emitted on each cycle
//and reports the offset between the host's clock time and the GPS time
//at top of second (actually, when the first data for the reporting
//cycle is received).
//
//Type TOFF is emitted once per second to watchers of a device and is
//intended to report the offset between the in-band report of the GPS, and
//seconds as reported by the system clock (which may be NTP-corrected) when
//the first valid timestamp of the reporting cycle is seen.
//
//Type TOFF contains second/microsecond pairs for two clocks; the realtime
//clock without NTP correction (the result of clock_gettime(CLOCK_REALTIME),
//but only to microsecond precision), and the ordinary system clock
//(which may be NTP corrected).
type TOFF struct {
	//Fixed: "TOFF"
	//
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Name of originating device
	//
	//Always? Yes. Type: string
	Device float64 `json:"device"`

	//Seconds from the GPS clock
	//
	//Always? Yes. Type: numeric
	RealSec float64 `json:"real_sec"`

	//Nanoseconds from the GPS clock
	//
	//Always? Yes. Type: numeric
	RealNsec float64 `json:"real_nsec"`

	//Seconds from the system clock
	//
	//Always? Yes. Type: numeric
	ClockSec float64 `json:"clock_sec"`

	//Nanoseconds from the system clock
	//
	//Always? Yes. Type: numeric
	ClockNsec float64 `json:"clock_nsec"`
}
