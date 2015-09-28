package gpsdjson

//DEVICES is the response type from a "?DEVICES" command.
//
//The gpsd daemon occasionally ships a bare DEVICE object to the client
//(that is, one not inside a DEVICES wrapper)
type DEVICES struct {
	//Fixed: "DEVICES"
	//
	//Always? Yes. Type: string
	Class string `json:"class"`

	//List of device descriptions
	//
	//Always? Yes. Type: list
	Devices []DEVICE `json:"devices"`

	//URL of the remote daemon reporting the device set.
	//If empty, this is a DEVICES response from the local daemon.
	//
	//Always? No. Type: string
	Remote string `json:"remote,omitempty"`
}
