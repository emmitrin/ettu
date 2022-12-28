package ettu

import (
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"net/http"
)

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			return c.Data
		}
	}

	return ""
}

func fetchStationForLetter(path string) (*Selection, error) {
	resp, err := http.Get(basePath + path)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	sel := &Selection{}

	doc.Find("h3").Each(func(i int, selection *goquery.Selection) {
		text := getText(selection.Nodes[0])

		switch text {
		case "Трамваи":
			for {
				selection = selection.Next()
				if len(selection.Nodes) == 0 {
					continue
				}

				node := selection.Nodes[0]

				if node.Data == "br" {
					continue
				} else if node.Data != "a" {
					break
				}

				station := TramStation{Type: TypeTram, Title: getText(node)}

				for _, attr := range node.Attr {
					if attr.Key == "href" {
						station.Path = attr.Val
					}
				}

				sel.Tram = append(sel.Tram, station)
			}

		case "Троллейбусы":
			for {
				selection = selection.Next()
				if len(selection.Nodes) == 0 {
					continue
				}

				node := selection.Nodes[0]

				if node.Data == "br" {
					continue
				} else if node.Data != "a" {
					break
				}

				station := TrolleyStation{Type: TypeTrolley, Title: getText(node)}

				for _, attr := range node.Attr {
					if attr.Key == "href" {
						station.Path = attr.Val
					}
				}

				sel.Trolley = append(sel.Trolley, station)
			}
		}
	})

	return sel, nil
}

// FetchAllStations fetches and returns all city's stations.
func FetchAllStations() (*Selection, error) {
	letters, err := fetchLetters()

	if err != nil {
		return nil, err
	}

	sel := &Selection{}

	for _, letter := range letters {
		letterSel, err := fetchStationForLetter(letter)
		if err != nil {
			return nil, err
		}

		sel.Tram = append(sel.Tram, letterSel.Tram...)
		sel.Trolley = append(sel.Trolley, letterSel.Trolley...)
	}

	return sel, err
}
