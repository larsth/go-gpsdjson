package gpsdjson

//ATT is a vehicle-attitude report type.
//
//It is returned by digital-compass and gyroscope sensors; depending on device,
//it may include: heading, pitch, roll, yaw, gyroscope, and magnetic-field
//readings.
//
//Because such sensors are often bundled as part of
//marine-navigation systems, the ATT response may
//also include water depth.
//
//The "class" and "mode" fields will reliably be present.
//Others may be reported or not depending on the specific device type.
type ATT struct {
	//Fixed: "ATT"
	//Always? Yes. Type: string
	Class string `json:"class"`

	//Name of originating device
	//Always? Yes. Type: string
	Device string `json:"device"`

	//Seconds since the Unix epoch, UTC.
	//May have a fractional part of up to .001sec precision.
	//Always? Yes. Type: time.Duration
	Time Duration `json:"time"`

	//Heading, degrees from true north.
	//Always? No. Type: numeric
	Heading float64 `json:"heading,omitempty"`

	//Magnetometer status.
	//Always? No. Type: string
	MagSt string `json:"mag_st,omitempty"`

	//Pitch in degrees.
	//Always? No. Type: numeric
	Pitch float64 `json:"pitch,omitempty"`

	//Pitch sensor status.
	//Always? No. Type: string
	PitchSt string `json:"pitch_st,omitempty"`

	//Yaw in degrees
	//Always? No. Type: numeric
	Yaw float64 `json:"yaw,omitempty"`

	//Yaw sensor status.
	//Always? No. Type: string
	YawSt string `json:"yaw_st,omitempty"`

	//Roll in degrees.
	//Always? No. Type: numeric
	Roll float64 `json:"roll,omitempty"`

	//Roll sensor status.
	//Always? No. Type: numeric
	RollSt string `json:"roll_st,omitempty"`

	//Local magnetic inclination, degrees,
	//positive when the magnetic field points downward (into the Earth).
	//Always? No. Type: numeric
	Dip float64 `json:"dip,omitempty"`

	//Scalar magnetic field strength.
	//(Comment by Lars Tørnes Hansen: micro Teslas ?)
	//Always? No. Type: numeric
	MagLen float64 `json:"mag_len,omitempty"`

	//X component of magnetic field strength.
	//Always? No. Type: numeric
	MagX float64 `json:"mag_x,omitempty"`

	//Y component of magnetic field strength.
	//Always? No. Type: numeric
	MagY float64 `json:"mag_y,omitempty"`

	//Z component of magnetic field strength.
	//Always? No. Type: numeric
	MagZ float64 `json:"mag_z,omitempty"`

	//Scalar acceleration.
	//Always? No. Type: numeric
	AccLen float64 `json:"acc_len,omitempty"`

	//X component of acceleration.
	//Always? No. Type: numeric
	AccX float64 `json:"acc_x,omitempty"`

	//Y component of acceleration.
	//Always? No. Type: numeric
	AccY float64 `json:"acc_y,omitempty"`

	//Z component of acceleration.
	//Always? No. Type: numeric
	AccZ float64 `json:"acc_z,omitempty"`

	//X component of acceleration (gyro).
	//Always? No. Type: numeric
	GyroX float64 `json:"gyro_x,omitempty"`

	//Y component of acceleration (gyro).
	//Always? No. Type: numeric
	GyroY float64 `json:"gyro_y,omitempty"`

	//Water depth in meters.
	//Always? No. Type: numeric
	Depth float64 `json:"depth,omitempty"`

	//Temperature at sensor, degrees centigrade.
	//Always? No. Type: numeric
	Temperature float64 `json:"temperature,omitempty"`

	//Tag   string    `json:"tag"`
}
