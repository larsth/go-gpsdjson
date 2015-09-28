package gpsdjson

//PPS is type that contains a message which is emitted each
//time the gpsd server sees a PPS (Pulse Per Second) strobe
// from a device.
//
//Type PPS is emitted once per second to watchers of a device
//emitting PPS, and is intended to report the offset between the
//start of the GPS second (when the 1PPS arrives) and seconds as
//reported by the system clock (which may be NTP-corrected).
//
//Type PPS contains second/nanosecond pairs for two clocks;
//the realtime clock without NTP correction (the result of
//clock_gettime(CLOCK_REALTIME), but only to microsecond precision)
//and the ordinary system clock (which may be NTP corrected).
//
//There are various sources of error in the reported clock times.
//For PPS delivered via a real serial-line strobe, serial-interrupt
//latency plus processing time to the timer call should be bounded
//above by about 10 microseconds; USB-to-serial control-line emulation
//is known to add jitter of about 50 microseconds. (Both figures are
//for GPSD running in non-realtime mode on an x86 with a gigahertz
//clock and are estimates based on measured latency in other applications.)
type PPS struct {
	//Fixed: "PPS"
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
