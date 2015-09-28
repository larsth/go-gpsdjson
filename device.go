package gpsdjson

//DEVICE is the request and response type from a "?DEVICE" command.
//
//If CommandDEVICE is followed by a ';' the state of a device, or sets
//(when followed by '=' and a DEVICE object) device-specific control bits,
// notably the device's speed and serial mode and the native-mode bit.
//The parameter-setting form (a request) will be rejected if more than
//one client is attached to the channel.
//
//Pay attention to the response, because it is possible for this command
//to fail if the GPS does not support a speed-switching command or only
//supports some combinations of serial modes. In case of failure, the
//daemon and GPS will continue to communicate at the old speed.
//
//Use the parameter-setting form with caution. On USB and Bluetooth GPSes
//it is also possible for serial mode setting to fail either because the
//serial adaptor chip does not support non-8N1 modes or because the device
//firmware does not properly synchronize the serial adaptor chip with the
//UART on the GPS chipset when the speed changes.
//These failures can hang your device, possibly requiring a GPS power cycle
//or (in extreme cases) physically disconnecting the NVRAM backup battery.
//
//The serial parameters will be omitted in a response describing a TCP/IP
//source such as an Ntrip, DGPSIP, or AIS feed.
//
//The gpsd daemon occasionally ships a bare DEVICE object to the client
//(that is, one not inside a DEVICES wrapper).
type DEVICE struct {
	//Fixed: "DEVICE"
	//
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Name the device for which the control bits are being reported,
	//or for which they are to be applied. This attribute may be omitted
	//only when there is exactly one subscribed channel.
	//
	//Always? No. Type: string
	Path string `json:"path,omitempty"`

	//Time the device was activated as an ISO8601 timestamp.
	//If the device is inactive this attribute is absent.
	//
	//Always? No. Type: string
	Activated string `json:"activated,omitempty"`

	//Bit vector of property flags. Currently defined flags are:
	//describe packet types seen so far (GPS, RTCM2, RTCM3, AIS).
	//Won't be reported if empty, e.g. before gpsd has seen
	//identifiable packets from the device.
	//
	//Always? No. Type: integer
	Flags DeviceFlags `json:"flags,omitempty"`

	//GPSD's name for the device driver type. Won't be reported before
	//gpsd has seen identifiable packets from the device.
	//
	//Always? No. Type: string
	Driver string `json:"driver,omitempty"`

	//Whatever version information the device returned.
	//
	// Always?
	//
	//When the daemon sees a delayed response to a probe
	//for subtype or firmware-version information.
	//
	// Type
	//
	//string
	Subtype string `json:"subtype,omitempty"`

	//Device speed in bits per second.
	//
	//Always? No. Type: integer
	Bps int `json:"bps,omitempty"`

	//N, O or E for no parity, odd, or even.
	//
	//Always? No. Type: string
	Parity string `json:"parity,omitempty"`

	//Stop bits (1 or 2).
	//
	//Always? No. Type: string
	Stopbits string `json:"stopbits,omitempty"`

	//0 means NMEA mode and 1 means alternate mode
	//(binary if it has one, for SiRF and Evermore
	//chipsets in particular).
	//
	//Attempting to set this mode on a non-GPS device
	//will yield an error.
	//
	//Always? No. Type: integer
	Native int `json:"native,omitempty"`

	//Device cycle time in seconds.
	//
	//Always? No. Type: real
	Cycle float64 `json:"cycle,omitempty"`

	//Device minimum cycle time in seconds.
	//Reported from ?CONFIGDEV when (and only when)
	//the rate is switchable.
	//
	//It is read-only and not settable.
	//
	//Always? No. Type: real
	Mincycle float64 `json:"mincycle,omitempty"`
}
