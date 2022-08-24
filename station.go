package ettu

import "strings"

// BaseStation describes a station. Since there is no difference between trolley and tram stations,
// the same object is used.
type BaseStation struct {
	Path  string
	Title string
}

type TramStation = BaseStation
type TrolleyStation = BaseStation

// StationByID searches and returns a station by its ID.
//
// The second parameter is set to false in case no station was found.
func (s *Selection) StationByID(id string) (BaseStation, bool) {
	if s == nil {
		return BaseStation{}, false
	}

	if !strings.HasPrefix(id, StationPathPrefix) {
		id = StationPathPrefix + id
	}

	for _, tram := range s.Tram {
		if tram.Path == id {
			return tram, true
		}
	}

	for _, trolley := range s.Trolley {
		if trolley.Path == id {
			return trolley, true
		}
	}

	return BaseStation{}, false
}
