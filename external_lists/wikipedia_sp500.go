package external_lists

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type WikipediaSP500 struct {
	symbols []string
}

func NewWikipediaSP500() ExternalList {
	return &WikipediaSP500{}
}

func (ws *WikipediaSP500) GetSymbols() ([]string, error) {
	if len(ws.symbols) < 1 {
		resp, err := http.Get("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
		if err != nil {
			return nil, err
		}
		b := resp.Body
		defer b.Close()

		z := html.NewTokenizer(b)

		rightTable := false
		firstTdOnRow := true

		for {

			tt := z.Next()

			switch tt {
			case html.ErrorToken:
				if z.Err() == io.EOF {
					return ws.symbols, nil
				}
				return ws.symbols, z.Err()
			case html.EndTagToken:
				t := z.Token()
				if t.Data == "table" {
					rightTable = false
				}
				if t.Data == "tr" {
					firstTdOnRow = true
				}
			case html.StartTagToken:
				t := z.Token()
				if t.Data == "table" {
					for _, attr := range t.Attr {
						if attr.Key == "id" && attr.Val == "constituents" {
							rightTable = true
						}
					}
				}
				if firstTdOnRow && rightTable && t.Data == "td" {
					tt = z.Next()
					if tt == html.StartTagToken {
						nt := z.Token()
						if nt.Data == "a" {
							tt = z.Next()
							if tt == html.TextToken {
								nt = z.Token()
								ws.symbols = append(ws.symbols, strings.TrimSpace(strings.ReplaceAll(nt.Data, ".", "")))
							}
						}
					}
					firstTdOnRow = false
				}

			}
		}
	}
	return ws.symbols, nil
}
