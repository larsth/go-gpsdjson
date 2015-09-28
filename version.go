package gpsdjson

//VERSION is a type that contains the response from a "?VERION" command.
//
//The gpsd daemon ships a VERSION response to each client when the client
//first connects to it.
type VERSION struct {
	//Fixed: "VERSION"
	//
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Public release level
	//
	//Always? Yes. Type: string
	Release string `json:"release"`

	//Internal revision-control level.
	//
	//Always? Yes. Type: string
	Rev string `json:"rev"`

	//API major revision level.
	//
	//Always? Yes. Type: numeric
	ProtoMajor int `json:"proto_major"`

	//API minor revision level.
	//
	//Always? Yes. Type: numeric
	ProtoMinor int `json:"proto_minor"`

	//URL of the remote daemon reporting this version.
	//If empty, this is the version of the local daemon.
	//
	//Always? No. Type: string
	Remote string `json:"remote,omitempty"`
}
