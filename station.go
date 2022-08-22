package ettu

// BaseStation describes a station. Since there is no difference between trolley and tram stations,
// the same object is used.
type BaseStation struct {
	Path  string
	Title string
}

type TramStation = BaseStation
type TrolleyStation = BaseStation
