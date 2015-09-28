package gpsdjson

//ERROR is an error message from the gpsd daemon in response
// to a syntactically invalid command or an unknown command.
type ERROR struct {
	//Fixed: "ERROR"
	//
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Textual error message
	//
	//Always? Yes. Type: string
	Message string `json:"message"`
}
