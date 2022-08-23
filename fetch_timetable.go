package ettu

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func FetchTimetableByID(stationID string) ([]CarInfo, error) {
	const stationPathPrefix = "/station/"

	if !strings.HasPrefix(stationID, stationPathPrefix) {
		stationID = stationPathPrefix + stationID
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

	var result []CarInfo

	current := CarInfo{}

	divs := doc.Find("div")
	divs.Each(func(i int, selection *goquery.Selection) {
		style, hasStyle := selection.Attr("style")
		if !hasStyle {
			return
		}

		if len(selection.Nodes) == 0 {
			return
		}

		switch {
		case strings.Contains(style, "width: 3em;"):
			// number (inside b-tag)

			current.Number = getText(selection.Nodes[0].FirstChild)

		case strings.Contains(style, "width: 4em;"):
			// ETA

			current.ETA = getText(selection.Nodes[0])

		case strings.Contains(style, "width: 5em;"):
			// distance

			current.Distance = getText(selection.Nodes[0])

			result = append(result, current)

			current = CarInfo{}
		}
	})

	return result, nil
}

func (station *BaseStation) FetchTimetable() ([]CarInfo, error) {
	return FetchTimetableByID(station.Path)
}
