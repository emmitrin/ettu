package ettu

import "fmt"

// VehicleInfo represents a vehicle (either tram or trolley) approaching the station.
type VehicleInfo struct {
	Route    string
	ETA      int // in minutes
	Distance int // in meters
}

func (v VehicleInfo) String() string {
	return fmt.Sprintf("%s: %d meters away, ETA: %d minutes", v.Route, v.Distance, v.ETA)
}
