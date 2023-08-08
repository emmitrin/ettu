package ettu

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrBadDistance = errors.New("ettu: bad vehicle distance label")
	ErrBadETA      = errors.New("ettu: bad eta label")
)

const StationPathPrefix = "/station/"

func FetchTimetableByID(stationID string) ([]VehicleInfo, error) {
	if !strings.HasPrefix(stationID, StationPathPrefix) {
		stationID = StationPathPrefix + stationID
	}

	resp, err := http.Get(basePath + stationID)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []VehicleInfo

	current := VehicleInfo{}

	var traverseError error

	divs := doc.Find("div")
	divs.EachWithBreak(func(i int, selection *goquery.Selection) bool {
		style, hasStyle := selection.Attr("style")
		if !hasStyle {
			return true
		}

		if len(selection.Nodes) == 0 {
			return true
		}

		switch {
		case strings.Contains(style, "width: 3em;"):
			// number (inside b-tag)

			current.Route = getText(selection.Nodes[0].FirstChild)

		case strings.Contains(style, "width: 4em;"):
			// ETA

			eta, err := parseETA(getText(selection.Nodes[0]))

			if err != nil {
				traverseError = err
				return false
			}

			current.ETA = eta

		case strings.Contains(style, "width: 5em;"):
			// distance

			dist, err := parseDistance(getText(selection.Nodes[0]))

			if err != nil {
				traverseError = err
				return false
			}

			current.Distance = dist

			result = append(result, current)

			current = VehicleInfo{}
		}

		return true
	})

	return result, traverseError
}

func (station *BaseStation) FetchTimetable() ([]VehicleInfo, error) {
	return FetchTimetableByID(station.Path)
}

func parseDistance(distance string) (int, error) {
	if !strings.HasSuffix(distance, " м") {
		return 0, ErrBadDistance
	}

	return strconv.Atoi(strings.TrimSuffix(distance, " м"))
}

func parseETA(distance string) (int, error) {
	if !strings.HasSuffix(distance, " мин") {
		return 0, ErrBadETA
	}

	return strconv.Atoi(strings.TrimSuffix(distance, " мин"))
}
