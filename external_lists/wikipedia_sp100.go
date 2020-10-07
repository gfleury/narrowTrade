package external_lists

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type WikipediaSP100 struct {
	symbols []string
}

func NewWikipediaSP100() ExternalList {
	return &WikipediaSP100{}
}

func (ws *WikipediaSP100) GetSymbols() ([]string, error) {
	if len(ws.symbols) < 1 {
		resp, err := http.Get("https://en.wikipedia.org/wiki/S%26P_100")
		if err != nil {
			return nil, err
		}
		b := resp.Body
		defer b.Close()

		z := html.NewTokenizer(b)

		rightTable := false

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
			case html.StartTagToken:
				t := z.Token()
				if t.Data == "table" {
					for _, attr := range t.Attr {
						if attr.Key == "id" && attr.Val == "constituents" {
							rightTable = true
						}
					}
				}
				if rightTable && t.Data == "td" {
					tt = z.Next()
					if tt == html.TextToken {
						nt := z.Token()
						ws.symbols = append(ws.symbols, strings.TrimSpace(strings.ReplaceAll(nt.Data, ".", "")))
					}
				}

			}
		}
	}
	return ws.symbols, nil
}
