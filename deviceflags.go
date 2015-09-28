package gpsdjson

//DeviceFlags contains flags about wether or not
//GPS, RTCM2, RTCM3, AIS, or nothing had been seen.
//
//DeviceFlags can SeenNothing or one to many of the other flags.
type DeviceFlags int32

const (
	//SeenNothing means that nothing had been seen.
	SeenNothing DeviceFlags = iota
	//SeenGPS means that GPS data had been seen.
	SeenGPS DeviceFlags = 1
	//SeenRTCM2 means that RTCM version 2 data had been seen.
	SeenRTCM2 DeviceFlags = 2
	//SeenRTCM3 means that RTCM version 3 data had been seen.
	SeenRTCM3 DeviceFlags = 4
	//SeenAIS means that AIS data had been seen.
	SeenAIS DeviceFlags = 8
)
