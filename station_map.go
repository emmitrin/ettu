package ettu

import "strings"

// Selection is a common object which describes a selection of stations.
// It can be created with a search query, or can handle all available stations.
type Selection struct {
	Tram    []TramStation
	Trolley []TrolleyStation
}

// Search performs a search across all stations and returns a new Selection.
//
// The returned Selection is always non-nil, even if there are no search results.
// The second return value is false in case no stations were found.
func (s *Selection) Search(q string) (*Selection, bool) {
	if len(q) == 0 {
		return s, true
	}

	q = strings.ToLower(q)

	sel := &Selection{}
	foundOne := false

	for _, tram := range s.Tram {
		if containsWords(tram.Title, q) {
			foundOne = true
			sel.Tram = append(sel.Tram, tram)
		}
	}

	for _, trolley := range s.Trolley {
		if containsWords(trolley.Title, q) {
			foundOne = true
			sel.Tram = append(sel.Tram, trolley)
		}
	}

	return sel, foundOne
}

func containsWords(haystack, needle string) bool {
	haystack, needle = strings.ToLower(haystack), strings.ToLower(needle)
	words := strings.Split(needle, " ")

	for _, w := range words {
		if !strings.Contains(haystack, w) {
			return false
		}
	}

	return true
}
