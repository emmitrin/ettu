package ettu

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const basePath = "https://online.ettu.ru"

func fetchLetters() ([]string, error) {
	resp, err := http.Get(basePath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []string

	nodes := doc.Find(".letter-link")
	nodes.Each(func(i int, selection *goquery.Selection) {
		href, ok := selection.Attr("href")

		if !ok {
			return
		}

		result = append(result, href)
	})

	if len(result) == 0 {
		return nil, errors.New("emmitrin/ettu: no stations available")
	}

	return result, nil
}
